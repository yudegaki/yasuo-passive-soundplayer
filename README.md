# YasuoPassiveSoundPlayer

League of Legends のチャンピオン「Yasuo」のパッシブスキルが発動した際に任意の音を再生するアプリケーションです。

## Build for Windows

### build
```bash
GOOS=windows GOARCH=amd64 go build -o YasuoPassiveSoundPlayer.exe main.go
```

## How to use

アプリケーションをビルド後、実行ファイルと同じディレクトリに `sound.mp3`という名前で、再生したい音声ファイル(mp3) を配置してください。

League of Legends の試合中にアプリケーションを実行すると、パッシブスキルが発動した際に音声ファイルが再生されます。

### exec

```commandprompt
YasuoPassiveSoundPlayer.exe
```
