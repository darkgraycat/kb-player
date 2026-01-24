## MATH!
For each sample at time t:
```go
sample = amplitude * sin(2 * π * frequency * t)
```

## STATES LC
┌─────────────┐
│   Normal    │ (print help, menus, etc.)
│   Terminal  │
└──────┬──────┘
       │ enter raw mode
       ▼
┌─────────────┐      ┌──────────────┐
│  Raw Mode   │◄────►│ Audio Loop   │ (concurrent, playing notes)
│  Input Loop │      │ (goroutine)  │
└──────┬──────┘      └──────────────┘
       │ exit raw mode
       ▼
┌─────────────┐
│   Normal    │ (back to regular terminal)
│   Terminal  │
└─────────────┘

## TODO: Recording into text file
Need a way to be able to record entered keys.
Then the main window should "watch" the file (I think I need duplex stream here)
Ability to play "file" on backgroud
