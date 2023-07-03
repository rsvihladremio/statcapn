//	Copyright 2023 Dremio Corporation
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//	http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// package pkg is where we build all of the code that is shared
package pkg

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path"
	"strings"
	"time"

	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/disk"
	"github.com/shirou/gopsutil/v3/mem"
)

type Args struct {
	IntervalSeconds int
	DurationSeconds int
	OutFile         string
}
type SystemMetricsRow struct {
	CollectionTimeStamp    time.Time `json:"collectionTimestamp"`
	rawUserCPUPercent      float64
	rawSystemCPUPercent    float64
	rawIdleCPUPercent      float64
	rawIOWaitCPUPercent    float64
	rawNiceCPUPercent      float64
	rawIRQCPUPercent       float64
	rawSoftIRQCPUPercent   float64
	rawStealCPUPercent     float64
	rawGuestCPUPercent     float64
	rawGuestNiceCPUPercent float64
	UserCPUPercent         string `json:"userCPUPercent"`
	SystemCPUPercent       string `json:"systmeCPUPercent"`
	IdleCPUPercent         string `json:"idleCPUPercent"`
	NiceCPUPercent         string `json:"niceCPUPercent"`
	IOWaitCPUPercent       string `json:"ioWaitCPUPercent"`
	IRQCPUPercent          string `json:"irqCPUPercent"`
	SoftIRQCPUPercent      string `json:"softIRQCPUPercent"`
	StealCPUPercent        string `json:"stealCPUPercent"`
	GuestCPUPercent        string `json:"guestCPUPercent"`
	GuestNiceCPUPercent    string `json:"guestCPUNicePercent"`
	QueueDepth             string `json:"queueDepth"`
	DiskLatency            string `json:"diskLatency"`
	ReadBytes              int64  `json:"readBytes"`
	WriteBytes             int64  `json:"writeBytes"`
	FreeRAMMB              int64  `json:"freeRAMMB"`
	CachedRAMMB            int64  `json:"cachedRAMMB"`
}

type CollectionParams struct {
	IntervalSeconds int
	DurationSeconds int
	RowWriter       func(SystemMetricsRow) error
}

func CollectSystemMetrics(params CollectionParams) error {
	if params.DurationSeconds < 1 {
		return fmt.Errorf("duration must be at least 1 second %v", params.DurationSeconds)
	}
	if params.IntervalSeconds < 1 {
		return fmt.Errorf("interval must be at least 1 second %v", params.IntervalSeconds)
	}
	interval := time.Second * time.Duration(params.IntervalSeconds)
	iterations := params.DurationSeconds / params.IntervalSeconds
	if iterations < 1 {
		return fmt.Errorf("interval of %v cannot be greater than the duration of %v", params.IntervalSeconds, params.DurationSeconds)
	}

	prevDiskIO, err := disk.IOCounters()
	if err != nil {
		return err
	}
	prevCPUTimes, err := cpu.Times(false)
	if err != nil {
		return err
	}
	for i := 0; i < iterations; i++ {
		// Sleep
		time.Sleep(interval)

		// CPU Times
		cpuTimes, err := cpu.Times(false)
		if err != nil {
			return err
		}

		// Memory
		memoryInfo, err := mem.VirtualMemory()
		if err != nil {
			return err
		}

		// Disk I/O
		diskIO, err := disk.IOCounters()
		if err != nil {
			return err
		}

		var weightedIOTime, totalIOs uint64
		var readBytes, writeBytes float64
		for i, io := range diskIO {
			p := prevDiskIO[i]
			weightedIOTime += io.WeightedIO - p.WeightedIO
			totalIOs += io.IoTime - p.IoTime

			if prev, ok := prevDiskIO[io.Name]; ok {
				readBytes += float64(io.ReadBytes-prev.ReadBytes) / 1024
				writeBytes += float64(io.WriteBytes-prev.WriteBytes) / 1024
			}
		}
		prevDiskIO = diskIO
		total := getTotalTime(cpuTimes[0], prevCPUTimes[0])
		var queueDepth float64
		var diskLatency float64
		if weightedIOTime > 0 {
			queueDepth = float64(weightedIOTime) / 1000
			diskLatency = float64(weightedIOTime) / float64(totalIOs)
		}

		row := SystemMetricsRow{}
		row.CollectionTimeStamp = time.Now()
		user := cpuTimes[0].User - prevCPUTimes[0].User
		if user > 0 {
			row.rawUserCPUPercent = (user / total) * 100
		}
		row.UserCPUPercent = fmt.Sprintf("%.2f", row.rawUserCPUPercent)
		system := cpuTimes[0].System - prevCPUTimes[0].System
		if system > 0 {
			row.rawSystemCPUPercent = (system / total) * 100
		}
		row.SystemCPUPercent = fmt.Sprintf("%.2f", row.rawSystemCPUPercent)
		idle := cpuTimes[0].Idle - prevCPUTimes[0].Idle
		if idle > 0 {
			row.rawIdleCPUPercent = (idle / total) * 100
		}
		row.IdleCPUPercent = fmt.Sprintf("%.2f", row.rawIdleCPUPercent)
		nice := cpuTimes[0].Nice - prevCPUTimes[0].Nice
		if nice > 0 {
			row.rawNiceCPUPercent = (nice / total) * 100
		}
		row.NiceCPUPercent = fmt.Sprintf("%.2f", row.rawNiceCPUPercent)
		iowait := cpuTimes[0].Iowait - prevCPUTimes[0].Iowait
		if iowait > 0 {
			row.rawIOWaitCPUPercent = (iowait / total) * 100
		}
		row.IOWaitCPUPercent = fmt.Sprintf("%.2f", row.rawIOWaitCPUPercent)

		irq := cpuTimes[0].Irq - prevCPUTimes[0].Irq
		if irq > 0 {
			row.rawIRQCPUPercent = (irq / total) * 100
		}
		row.IRQCPUPercent = fmt.Sprintf("%.2f", row.rawIRQCPUPercent)

		softIRQ := cpuTimes[0].Softirq - prevCPUTimes[0].Softirq
		if softIRQ > 0 {
			row.rawSoftIRQCPUPercent = (softIRQ / total) * 100
		}
		row.SoftIRQCPUPercent = fmt.Sprintf("%.2f", row.rawSoftIRQCPUPercent)
		steal := cpuTimes[0].Steal - prevCPUTimes[0].Steal
		if steal > 0 {
			row.rawStealCPUPercent = (steal / total) * 100
		}
		row.StealCPUPercent = fmt.Sprintf("%.2f", row.rawStealCPUPercent)

		guestCPU := cpuTimes[0].Guest - prevCPUTimes[0].Guest
		if guestCPU > 0 {
			row.rawGuestCPUPercent = (guestCPU / total) * 100
		}
		row.GuestCPUPercent = fmt.Sprintf("%.2f", row.rawGuestCPUPercent)
		guestCPUNice := cpuTimes[0].GuestNice - prevCPUTimes[0].GuestNice
		if guestCPUNice > 0 {
			row.rawGuestNiceCPUPercent = (guestCPUNice / total) * 100
		}
		row.GuestNiceCPUPercent = fmt.Sprintf("%.2f", row.rawGuestNiceCPUPercent)

		prevCPUTimes = cpuTimes
		row.DiskLatency = fmt.Sprintf("%.2f", diskLatency)
		row.QueueDepth = fmt.Sprintf("%.2f", queueDepth)

		var memoryFreeMB float64
		if memoryInfo.Free > 0 {
			memoryFreeMB = float64(memoryInfo.Free) / (1024 * 1024)
		}
		var memoryCachedMB float64
		if memoryCachedMB > 0 {
			memoryCachedMB = float64(memoryInfo.Cached) / (1024 * 1024)
		}
		row.FreeRAMMB = int64(memoryFreeMB)
		row.CachedRAMMB = int64(memoryCachedMB)
		if err := params.RowWriter(row); err != nil {
			return err
		}
	}
	return nil
}

func SystemMetrics(args Args) error {
	var w io.Writer
	var rowWriter func(SystemMetricsRow) error
	var cleanup func() error
	outputFile := args.OutFile

	if strings.HasSuffix(outputFile, ".json") {
		f, err := os.Create(path.Clean(outputFile))
		if err != nil {
			return fmt.Errorf("unable to create file %v due to error '%w'", outputFile, err)
		}
		w = f
		// we manually close this so we do not care that we are not handling the error
		defer f.Close()

		bufWriter := bufio.NewWriter(w)
		cleanup = func() error {
			if err := bufWriter.Flush(); err != nil {
				return fmt.Errorf("unable to flush metrics file %v due to error %w", outputFile, err)
			}
			if err := f.Close(); err != nil {
				return fmt.Errorf("unable to close metrics file %v due to error %w", outputFile, err)
			}
			return nil
		}
		//write json file
		rowWriter = func(row SystemMetricsRow) error {
			str, err := json.Marshal(&row)
			if err != nil {
				return fmt.Errorf("unable to marshal row %#v due to error %w", row, err)
			}
			_, err = bufWriter.Write(str)
			if err != nil {
				return fmt.Errorf("unable to write to json file due to error %w", err)
			}
			return nil
		}
		if err != nil {
			return fmt.Errorf("unable to write metrics file %v due to error %w", outputFile, err)
		}
	} else {
		if outputFile == "" {
			cleanup = func() error { return nil }
			w = os.Stdout
		} else {
			f, err := os.Create(path.Clean(outputFile))
			if err != nil {
				return fmt.Errorf("unable to create file %v due to error '%w'", outputFile, err)
			}
			cleanup = func() error {
				if err := f.Close(); err != nil {
					return fmt.Errorf("unable to close metrics file %v due to error %w", outputFile, err)
				}
				return nil
			}
			w = f
			// we don't care as this is just an emergency cleanup we manually call "cleanup" which closes the file anyway
			defer f.Close()
		}

		//write metrics.txt file
		template := "%25s\t%10s\t%10s\t%10s\t%10s\t%10s\t%10s\t%10s\t%10s\t%10s\t%10s"
		floatTemplate := "%.2f"
		percentTemplate := "%s%%"
		txtHeader := fmt.Sprintf(template, "Timestamp", "usr %%", "sys %%", "iowait %%", "other %%", "idl %%", "Queue", "Latency (ms)", "Read (MB/s)", "Write (MB/s)", "Free Mem (GB)")
		if _, err := fmt.Fprintln(w, txtHeader); err != nil {
			return fmt.Errorf("unable to write metrics file %v due to error %w", outputFile, err)
		}
		rowWriter = func(row SystemMetricsRow) error {
			otherCPU := row.rawNiceCPUPercent + row.rawIRQCPUPercent + row.rawSoftIRQCPUPercent + row.rawStealCPUPercent + row.rawGuestCPUPercent + row.rawGuestNiceCPUPercent
			var readBytesMB, writeBytesMB, freeRAMGB float64
			if row.ReadBytes > 0 {
				readBytesMB = float64(row.ReadBytes) / (1024 * 1024)
			}
			if row.WriteBytes > 0 {
				writeBytesMB = float64(row.WriteBytes) / (1024 * 1024)
			}
			if row.FreeRAMMB > 0 {
				freeRAMGB = float64(row.FreeRAMMB) / 1024.0
			}
			rowString := fmt.Sprintf(template,
				row.CollectionTimeStamp.Format(time.RFC3339),
				fmt.Sprintf(percentTemplate, row.UserCPUPercent),
				fmt.Sprintf(percentTemplate, row.SystemCPUPercent),
				fmt.Sprintf(percentTemplate, row.IOWaitCPUPercent),
				fmt.Sprintf(percentTemplate, fmt.Sprintf(floatTemplate, otherCPU)),
				fmt.Sprintf(percentTemplate, row.IdleCPUPercent),
				row.QueueDepth,
				row.DiskLatency,
				fmt.Sprintf(floatTemplate, readBytesMB),
				fmt.Sprintf(floatTemplate, writeBytesMB),
				fmt.Sprintf(floatTemplate, freeRAMGB))
			if _, err := fmt.Fprintln(w, rowString); err != nil {
				return fmt.Errorf("unable to write metrics file %v due to error %w", outputFile, err)
			}
			return nil
		}
	}
	params := CollectionParams{
		DurationSeconds: args.DurationSeconds,
		IntervalSeconds: args.IntervalSeconds,
		RowWriter:       rowWriter,
	}

	if err := CollectSystemMetrics(params); err != nil {
		return fmt.Errorf("unable to collect system metrics with error %v", err)
	}

	return cleanup()
}

func getTotalTime(c cpu.TimesStat, p cpu.TimesStat) float64 {
	current := c.User + c.System + c.Idle + c.Nice + c.Iowait + c.Irq +
		c.Softirq + c.Steal + c.Guest + c.GuestNice
	prev := p.User + p.System + p.Idle + p.Nice + p.Iowait + p.Irq +
		p.Softirq + p.Steal + p.Guest + p.GuestNice
	return current - prev
}