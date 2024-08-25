###### *<div align = right><sub>Design By Achno</sub></div>*
<div align = center><img src="assets/file.png"><br><br>

&ensp;[<kbd> <br> Overview <br> </kbd>](#overview-)&ensp;
&ensp;[<kbd> <br> Theme <br> </kbd>](#themes-)&ensp;
&ensp;[<kbd> <br> Usage <br> </kbd>](#usage-)&ensp;
&ensp;[<kbd> <br> Installation <br> </kbd>](#installation-)&ensp;
&ensp;[<kbd> <br> Contributions <br> </kbd>](#contributions-)&ensp;
<br><br><br><br></div>


```


██╗ ██╗ ██╗ ██╗       ██████╗  ██████╗  ██████╗██╗  ██╗███████╗ █████╗ ████████╗
╚██╗╚██╗╚██╗╚██╗     ██╔════╝ ██╔═══██╗██╔════╝██║  ██║██╔════╝██╔══██╗╚══██╔══╝
 ╚██╗╚██╗╚██╗╚██╗    ██║  ███╗██║   ██║██║     ███████║█████╗  ███████║   ██║   
 ██╔╝██╔╝██╔╝██╔╝    ██║   ██║██║   ██║██║     ██╔══██║██╔══╝  ██╔══██║   ██║   
██╔╝██╔╝██╔╝██╔╝     ╚██████╔╝╚██████╔╝╚██████╗██║  ██║███████╗██║  ██║   ██║   
╚═╝ ╚═╝ ╚═╝ ╚═╝       ╚═════╝  ╚═════╝  ╚═════╝╚═╝  ╚═╝╚══════╝╚═╝  ╚═╝   ╚═╝   
                                                                                
                                                                              
```

# Overview 🖼️

Gocheat is a TUI app which allows you  to create beautiful custom cheatsheets for your keybindings,hotkeys or shell aliases 

> I have hundreds of keybinds and tens of aliases that i cant remember, so i needed something that can look them up in seconds ;)

<br>

https://github.com/user-attachments/assets/cd7039de-203d-47a2-889c-a14639a5e94e

<br>

## Features

- `Fuzzy filtering` mispell words and still find them. You also have 2 modes, filter by `keybind/description` or filter by `tag`, check `Usage` section for more info.

- Easily and quickly `add` and `remove` entries of keybindings or aliases via the TUI itself.

- `Customize theme` change the colors of gocheat to the ones you prefer.

- `Lightweight`

- `Vim keys`

## Planned features

`Tabs` : Tabs which allow you to create multiple cheatsheets, for example one for `keybinds` and one for `shell aliases`

`Pin items`: Pinned items will be shown first before all items so you dont even have to filter frequently forgoten items 


# Themes 🎨

Colors for the Theme can be configured in `~/.config/gocheat/config.json`  in `"styles":{}`

```json
{
  "items": [
    {
      "title": "Minimize Window : meta + m",
      "tag": "Kwin"
    },
    {
      "title": "Maximize window : meta + up",
      "tag": "Kwin"
    },
  ],
  "styles": {
    "subtext": "#6c7086",
    "accent": "#b4befe"
  }
}

```
Important❗ The `background` color for gocheat is derived the background color of your terminal  

In the coming updates the color for the `filter`,`cursor` and the `arrow icon` for the forms will be configurable

# Usage ⚙️

Once you have launched the TUI you can hit `ctrl+h` to show the Help screen which shows you the keybinds for every screen. The tldr version of them are : 


| Keybinds      | Description   |Screen |
| ------------- |:-------------:| -----:|
| `ctrl+j`      | Add an entry  | List |
| `Enter`      | Confirm an entry | Form |
| `ctrl+x`      | Delete an entry| List |
| `/`           | Start filtering| List |
| `ctrl+f`      | Toggle Filter by Tag  | List |
| `esc`      | Go back to the List screen or exit filtering  | - |
| `ctrl+c`      | Exit the app  | - |

<br>

Notes 🗒️: You can modify the `~/.config/gocheat/config.json` directly to add,remove,edit entries

# Instalatlion 📦

### Arch linux - AUR 

```
yay -S gocheat
```

### Build from source

🔨 Clone the repo, build the project and move it inside your `$PATH`

```
git clone https://github.com/Achno/gocheat
cd gocheat
go build
sudo cp gowall /usr/local/bin/
gocheat
```

Notes 🗒️ : if you have `$GOPATH` setup correctly
Eg. you have the following in your .zshrc / .bashrc
```bash
export GOPATH=$(go env GOPATH)
export GOBIN=$GOPATH/bin
export PATH=$PATH:$GOBIN
```
You can simply use `go install`
```bash
go install github.com/Achno/gocheat@latest
```

# Contributions :handshake:

Feel free to suggest any cool features that would improve gocheat even further by opening an `issue`  

If you want to contribute a feature or fix a bug please open a `Pull request` 


