# shinyhunter

A cli tool for automatically grinding shiny starters in Pokémon FireRed.
Place mouse on normally colored starter inside the menu where you check if it's a shiny.
This pixel location is saved after the 5-second countdown at the beginning.
After every softreset and navigation back to that point, it checks if the color of this pixel changed.
After that, you simply have to keep the emulator open.
Tested successfully on a Mac with mgba emulator with text speed 3 and no fast-forwarding.