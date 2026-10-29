# Music Roulette 🎵

Lab 6 · Favorites Roulette — iOS Development, KBTU
Author: Inara Assylzhanova

Single-screen iOS app (Swift, UIKit, Storyboard, no Auto Layout) that shows a random card from my 10 favorite music artists. Each tap on the button shows a new artist (never the same one twice in a row) with a photo, name and short description.

## Category
Music — 10 favorite artists: Валентин Стрыкало, Darkhan Juzz, Монеточка, Dos Mukasan, Noize MC, Face, Кайрат Нуртас, Нурлан Онербаев, Кино, Батырхан Шукенов.

## Screenshots

| Launch screen | Main screen | Home Screen |
|---|---|---|
| ![Launch](screenshots/launch.png) | ![Main](screenshots/main.png) | ![Home](screenshots/home.png) |

## How it works
- Data is stored as an array of `Favorite` structs (`imageName`, `title`, `description`).
- Images are stored in `Assets.xcassets`.
- Layout is built in `Main.storyboard` with `UIImageView`, `UILabel` and `UIButton`, without constraints.
- The button picks a random index from all indices except the last shown one, so the same artist never appears twice in a row.
- No force unwraps except `@IBOutlet` declarations.
