# waybar-hyprland-window-name

A Waybar module that displays the name of the focused window in Hyprland.

## Build

```bash
make
```

## Usage

```jsonc
// ~/.config/waybar/config.jsonc
{
  // ...
  "modules-left": [/* ... */, "custom/window_name", /* ... */],
  // ...
  "custom/window_name": {
    "exec": "~/.config/waybar/waybar-hyprland-window-name"
  },
  // ...
}
```
