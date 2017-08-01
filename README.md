## barista

Keeps your Mac awake while it's on the clock.

Give it a config file that maps days of the week to time ranges to stay awake:

```json
{
  "monday": "09:00 - 19:00",
  "wednesday": "09:00 - 19:00",
  "thursday": "09:00 - 12:00",
  "friday": "09:00 - 19:00"
}
```

And barista will keep your Mac from idle sleeping during the hours specified.
Display sleep isn't prevented.

Todo:

- LaunchAgent
- Option to prevent display sleep

