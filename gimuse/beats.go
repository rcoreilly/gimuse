// Copyright (c) 2020, The GiMuse Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"log"
	"sync"
	"time"

	"gitlab.com/gomidi/midi/writer"
	"gitlab.com/gomidi/rtmididrv"
)

// LiveBeats sends live beats out over midi device
func LiveBeats() error {
	drv, err := rtmididrv.New()
	if err != nil {
		log.Println(err)
		return err
	}

	// make sure to close all open ports at the end
	defer drv.Close()

	out, err := drv.OpenVirtualOut("gimuse")
	if err != nil {
		log.Println(err)
		return err
	}
	defer out.Close()

	// the writer we are writing to
	wr := writer.New(out)

	wr.SetChannel(10) // percussion

	wait := &sync.WaitGroup{}

	go Repeats(wr, wait, 1000, 50, 64, 350, 50) // 64 = E4,  50 = high tom
	// time.Sleep(75 * time.Millisecond)
	// go Repeats(wr, 1000, 51, 80, 199, 1) // 51 = ride cymbal

	// time.Sleep(100 * time.Millisecond)
	go Repeats(wr, wait, 1000, 48, 64, 450, 50) // 48 = high-mid tom

	time.Sleep(100 * time.Millisecond)
	wait.Wait()

	return nil
}

// Repeats fires a repeated key at given velocity for given amount on and off
func Repeats(wr writer.ChannelWriter, wait *sync.WaitGroup, reps, key, vel int, onMSec, offMSec time.Duration) {
	wait.Add(1)
	for i := 0; i < reps; i++ {
		writer.NoteOn(wr, uint8(key), uint8(vel))
		time.Sleep(onMSec * time.Millisecond)
		writer.NoteOff(wr, uint8(key))
		time.Sleep(offMSec * time.Millisecond)
	}
	wait.Done()
}
