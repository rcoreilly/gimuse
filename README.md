# gimuse

Experiments with music in Go / GoGi.

Initially, using Go to generate rhythmic MIDI sequences of various sorts, which can then be played / recorded on GarageBand.

Using this awesome midi library which just works on mac: https://gitlab.com/gomidi/midi

Sending midi events live using time.Sleep is not very reliable.. better to write the SMF file.

# Tempo

* https://en.wikipedia.org/wiki/Tempo
* https://guitargearfinder.com/guides/convert-ms-milliseconds-bpm-beats-per-minute-vice-versa/

* 75 BPM = 800 msec quarter note (beat) = just slower than *andante*

* 100 BPM = 600 msec quarter note = *moderato*

* 120 BPM = 500 msec quarter note = *allegro*

* 150 BPM = 400 msec quarter note = *vivace*

# Midi keys

| Clef | Note    | Key # | Frequency Hz |
| -----|---------| ------|-----------|
| Bass | C3      | 48   | 131     |
|      | C♯3/D♭3  | 49  |  139     |
|      | D3      | 50   |  147     |    
|      | D♯3/E♭3  | 51   |  156     |    
|      | E3      | 52    | 165     |    
|      | F3      | 53    | 175     |    
|      | F♯3/G♭3  | 54   |  185     |    
|      | G3      |  55   |  196     |    
|      | G♯3/A♭3  | 56   |  208     |    
|      | A3      | 57   |  220     |    
|      | A♯3/B♭3  | 58  |   233     |    
|      | B3      | 59   |  247     |    
| Both | C4 (middle C)  |  60  |  262   | 
| Treble | C♯4/D♭4 | 61  |  277     |    
|      | D4      | 62   |  294     |    
|      | D♯4/E♭4  | 63  |   311     |    
|      | E4      | 64   |  330     |    
|      | F4      | 65   |  349     |    
|      | F♯4/G♭4  | 66  |   370     |    
|      | G4      | 67   |  392     |    
|      | G♯4/A♭4 |  68  |   415     |    
|      | A4      | 69   |  440     |    
|      | A♯4/B♭4  | 70  |   466     |    
|      | B4      | 71   |  494     |    
|      | C5      | 72   |  523     |    

## Percussion

| Key | Instrument         |
|-----|--------------------|
| 35  | Acoustic Bass Drum |
| 36  | Bass Drum 1        |
| 37  | Side Stick/Rimshot |
| 38  | Acoustic Snare     |
| 39  | Hand Clap          |
| 40  | Electric Snare     |
| 41  | Low Floor Tom      |
| 42  | Closed Hi-hat      |
| 43  | High Floor Tom     |
| 44  | Pedal Hi-hat       |
| 45  | Low Tom            |
| 46  | Open Hi-hat        |
| 47  | Low-Mid Tom        | 
| 48  | Hi-Mid Tom         | 
| 49  | Crash Cymbal 1     | 
| 50  | High Tom           | 
| 51  | Ride Cymbal 1      |
| 52  | Chinese Cymbal     |
| 53  | Ride Bell          |
| 54  | Tambourine         |
| 55  | Splash Cymbal      |
| 56  | Cowbell            |
| 57  | Crash Cymbal 2     |
| 58  | Vibra Slap         |
| 59  | Ride Cymbal 2      |
| 60  | High Bongo         |
| 61  | Low Bongo          |
| 62  | Mute High Conga    |
| 63  | Open High Conga    |
| 64  | Low Conga          |
| 65  | High Timbale       |
| 66  | Low Timbale        |
| 67  | High Agogô         |
| 68  | Low Agogô          |
| 69  | Cabasa             |
| 70  | Maracas            |
| 71  | Short Whistle      |
| 72  | Long Whistle       |
| 73  | Short Güiro        |
| 74  | Long Güiro         |
| 75  | Claves             |
| 76  | High Wood Block    |
| 77  | Low Wood Block     |
| 78  | Mute Cuíca         |
| 79  | Open Cuíca         |
| 80  | Mute Triangle      |
| 81  | Open Triangle      |

