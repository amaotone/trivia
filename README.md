# Trivia

Trivia is a tool makes your life richer.

## Install

```bash
$ go get -u github.com/amaotone/trivia
```

## Usage

Get a word randomly from wikipedia with `trivia` command.

```bash
$ trivia
2016–17 Elche CF season
Elche Club de Fútbol, S.A.D. (Valencian: Elx Club de Futbol, S.A.D.), is a Spanish football team based in Elche, Province of Alicante, in the Valencian Community.During the 2016-17 campaign they will be competing in the following competitions: Segunda League, Copa del Rey.
```

## Option

Specify the language:

```bash
$ trivia -l ja
マーティン・ガードナー
マーティン・ガードナー (英語: Martin Gardner、1914年10月21日 - 2010年5月22日) は、アメリカ合衆国の数学者、著述家、アマチュア手品師。科学的懐疑論者であり、疑似科学・超常現象批判でも知られている。生涯に70冊以上もの著作を遺した。
```

Set your default language:

```bash
$ trivia set -l ja
$ trivia
カウ高校とパハラ小学校
カウ高校とパハラ小学校（カウこうこうとパハラしょうがっこう、英語: Kau High and Pahala Elementary School）はアメリカ合衆国ハワイ州ハワイ島のカウ地区のパハラにある公立高校（9年～12年）、中学、小学校（幼稚園を含む）である。
```

Available languages are listed in https://meta.wikimedia.org/wiki/List_of_Wikipedias
