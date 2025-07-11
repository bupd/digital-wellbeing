# digital-wellbeing
Linux app more than Digital Wellbeing found in Androids.


![ChatGPT Image Apr 18, 2025, 03_14_14 PM](https://github.com/user-attachments/assets/c8c61ce8-708f-43a0-ae84-a78ae7321c6f)


## Installation
```shell
go install github.com/bupd/digital-wellbeing/cmd/dw@latest
```

## Run
```shell
dw
```

## Config
```toml
# Default configuration for Digital Wellbeing
# Add your configuration here
PORT = "8888"
# you can set to data for a single database
# daily or weekly to use new database
DBNAME = "daily" # will use a new DB for everyday
```

## Add DB Migrations
I am still looking for a way on how to make the migrations on newly created DBs
for now copy all files `sql/schema` and place it in `~/.digital-wellbeing/migrations/`

## Output

![WhatsApp Image 2025-02-19 at 5 59 31 AM](https://github.com/user-attachments/assets/8e20fe00-94a5-449d-9a1f-8232cc548d10)


## Tools I am gonna use
xev && evtest for logging different devices input which would look something like the below




```bash
so the below are the keys got from

```
❯ xev | sed -ne '/^KeyPress/,/^$/p'
```

KeyPress event, serial 34, synthetic NO, window 0x5a00001,
    root 0x1d6, subw 0x0, time 12136878, (476,588), root:(1447,604),
    state 0x0, keycode 174 (keysym 0x1008ff15, XF86AudioStop), same_screen YES,
    XLookupString gives 0 bytes:
    XmbLookupString gives 0 bytes:
    XFilterEvent returns: False

KeyPress event, serial 34, synthetic NO, window 0x5a00001,
    root 0x1d6, subw 0x0, time 12138880, (476,588), root:(1447,604),
    state 0x0, keycode 173 (keysym 0x1008ff16, XF86AudioPrev), same_screen YES,
    XLookupString gives 0 bytes:
    XmbLookupString gives 0 bytes:
    XFilterEvent returns: False

KeyPress event, serial 34, synthetic NO, window 0x5a00001,
    root 0x1d6, subw 0x0, time 12141301, (476,588), root:(1447,604),
    state 0x0, keycode 172 (keysym 0x1008ff14, XF86AudioPlay), same_screen YES,
    XLookupString gives 0 bytes:
    XmbLookupString gives 0 bytes:
    XFilterEvent returns: False

KeyPress event, serial 34, synthetic NO, window 0x5a00001,
    root 0x1d6, subw 0x0, time 12144303, (476,588), root:(1447,604),
    state 0x0, keycode 171 (keysym 0x1008ff17, XF86AudioNext), same_screen YES,
    XLookupString gives 0 bytes:
    XmbLookupString gives 0 bytes:
    XFilterEvent returns: False

KeyPress event, serial 34, synthetic NO, window 0x5a00001,
    root 0x1d6, subw 0x0, time 12148230, (476,588), root:(1447,604),
    state 0x0, keycode 64 (keysym 0xffe9, Alt_L), same_screen YES,
    XLookupString gives 0 bytes:
    XmbLookupString gives 0 bytes:
    XFilterEvent returns: False

```

## Digital wellbeing should have the following features:
- taking all keystrokes & mouse clicks & mouse drags.
- keystroke and mouse movement.
	- this metric shows mouse vs keyboard graph. (% of keyboard used over mouse. etc)
- should have open windows and tabs.
- measure ram and disk usage & list of devices connected.
- total screen time time in each app
- read browser information and which apps use more of your memory.
- idle time. (auto power off after 10 mins of inactivity.)
- note volume % with time of connected devices
- note power usage (if possible)
- list how many times sounds are heard (i.e monitoring headphones with times)
- sync with calendar and update the schedule and see if its work and you are not working add it as a missed task.
	- try adding it as a thing not productive (% of distracted from the work.)


## commands that are useful
- xev - monitoring the input and output devices such as mouse
  - evtest - this is also an extension of xev with pretty logging events only specific events.

- wmctrl -l -- to list all open windows
- xdotool getwindowfocus getwindowname -- get current focused window (every 5 secs).
- check -> xdotool getwindowfocus getwindowclassname -- get the current class if it is kitty execute the above one to get the window name to determine youtube or any other thing
- also if shown exec sessionizer or t -- means coding as per our thing. this should be configurable setting the rules that is teaching the app how to track it.
-

# IDEAS
- heatmap showing keys pressed over time
- most active hours / peak typing hours
- peak key ranking based on hour/day etc
