# NS-RPC

The definitive way to display your Nintendo Switch games in Discord. 🎮

## Introduction

NS-RPC (Nintendo Switch Rich Presence) is a Wails app for Windows and macOS.
It makes it easy for anyone to share what they are playing on the Switch to Discord in a fancy Rich Presence, like a PC game.

This app was built using [Wails](https://wails.io) (🏴󠁧󠁢󠁷󠁬󠁳󠁿 pride) and [SolidJS](https://solidjs.com).

![NS-RPC's design](https://i.imgur.com/FRbQwzC.png)

### With NS-RPC you can..

- Display that you are using your Switch across all of Discord.
- Select from an extensive list of games to show off.
- Set a custom status message to let everyone know exactly what you're doing.
- Pin your favourite games into a quick list.
- Experience my _questionable_ user interface.

## Prerequisites

All you need to get going is some common sense and the [Discord App](https://discordapp.com) installed to the same machine.

Users running Windows 10 or earlier _may_ encounter issues running NS-RPC due to Wails' use of **Microsoft WebView2** on Windows. **If** you do encounter problems, ensure this is installed.

## Installing

If you're looking for convenience, you'll find already built copies of NS-RPC for
both Windows and macOS [here](https://github.com/Da532/NS-RPC/releases).

## Rewrite

Long time users may realise this is a brand new app! 
NS-RPC's original codebase was not something I wanted to maintain.
It was the first project I wrote in JavaScript and I utilised Electron for this.

The new version uses Wails rather than Electron which I much prefer working in.
The frontend uses SolidJS. I much prefer using this to React for its sheer speed and removal of jank, while still using JSX.

## Anything else?

Not as of yet. If you have feature suggestions or need support, head over to this handy [Discord server](https://discord.gg/StDcdMu) and talk to us.

Have a good one!
