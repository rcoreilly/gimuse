// Copyright (c) 2020, The GiMuse Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/emer/etable/etable"
	"github.com/gomidi/midi/writer"
)

func GenBeatsTable(notes *etable.Table) {
	// todo: could write out to a table, look at it, generate a table, etc
}

func WriteSMF(fname string, notes *etable.Table) {
	dir := os.TempDir()
	f := filepath.Join(dir, "smf-test.mid")

	defer os.Remove(f)

	err := writer.WriteSMF(f, 2, func(wr *writer.SMF) error {

		wr.SetChannel(11) // sets the channel for the next messages
		writer.NoteOn(wr, 120, 50)
		wr.SetDelta(120)
		writer.NoteOff(wr, 120)

		wr.SetDelta(240)
		writer.NoteOn(wr, 125, 50)
		wr.SetDelta(20)
		writer.NoteOff(wr, 125)
		writer.EndOfTrack(wr)

		wr.SetChannel(2)
		writer.NoteOn(wr, 120, 50)
		wr.SetDelta(60)
		writer.NoteOff(wr, 120)
		writer.EndOfTrack(wr)
		return nil
	})

	if err != nil {
		fmt.Printf("could not write SMF file %v\n", f)
		return
	}
}
