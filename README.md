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
    "exec": "~/.config/waybar/waybar-hyprland-window-name",
    "return-type": "json"
  },
  // ...
}
```

```css
/* ~/.config/waybar/style.css */
/* ... */
.waybar-hyprland-window-name {
  font-weight: 700;
  padding-top: 2px;
  padding-left: 4px;
  padding-right: 4px;
}
/* ... */
```
