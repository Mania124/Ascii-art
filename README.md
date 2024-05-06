## Objectives
What is meant by a graphic representation using ASCII, is to write the string received using ASCII characters, as you can see in the example below:
```
    @@@@@@BB@@@@``^^``^^``@@BB$$@@BB$$
    @@%%$$$$^^^^WW&&8888&&^^""BBBB@@@@
    @@@@@@""WW8888&&WW888888WW``@@@@$$
    BB$$``&&&&WWWW8888&&&&8888&&``@@@@
    $$``&&WW88&&88&&&&8888&&88WW88``$$
    @@""&&&&&&&&88888888&&&&&&88&&``$$
    ``````^^``^^^^^^````""^^``^^``^^``
    ""WW^^@@@@^^``````^^BB@@^^``^^&&``
    ^^&&^^@@````^^``&&``@@````^^^^&&``
    ``WW&&^^""``^^WW&&&&""``^^^^&&88``
    ^^8888&&&&&&WW88&&88WW&&&&88&&WW``
    @@``&&88888888WW&&WW88&&88WW88^^$$
    @@""88&&&&&&&&888888&&``^^&&88``$$
    @@@@^^&&&&&&""``^^^^^^8888&&^^@@@@
    @@@@@@^^888888&&88&&&&MM88^^BB$$$$
    @@@@@@BB````&&&&&&&&88""``BB@@BB$$
    $$@@$$$$$$$$``````````@@$$@@$$$$$$
```

- This project should handle an input with numbers, letters, spaces, special characters and \n.
- Take a look at the ASCII manual.

# ASCII-ART
* This program is designed in GO
* Only standard Go libraries were utilised
* Its main purpose is to print out ascii-art.

## Overview
* This program is designed to take input from the command line interface and produce a coreesponding graphical ascii pattern.
* The patterns in use are predetermined by a file that is provided with the program
* All paterns for numbers, letters, spaces, special characters and new line("\n") are handled 

## Installation
Clone the repository and cd into it. The program accepts string input to output its ascii graphics and an optional flag, that determines the type of ascii graphic to use
#### procedure 
```bash
git clone https://learn.zone01kisumu.ke/git/nichotieno/ascii-art.git
cd ascii-art
go run . "Text to output its graphis" [-s or -t or shadow or thinkertoy]
```

## Dimension
* The output covers 8 lines for every character-art with an allowance of a single line for spacing at the top of the art

## Requirements
* The program requires the user to input atleast 2(two) command line arguments for it to output the art
* Failuire to adhere to this will lead to the program informing the user of the unavailability and the user can relaunch the program again twith necessary modification.

* This work has been done by @nichotieno (captain) and @hshikuku under the guidance and stewardship of @zone01kisumu

#### usage
```bash
student$ go run . "" | cat -e
student$ go run . "\n" | cat -e
$
student$ go run . "Hello\n" | cat -e
 _    _          _   _          $
| |  | |        | | | |         $
| |__| |   ___  | | | |   ___   $
|  __  |  / _ \ | | | |  / _ \  $
| |  | | |  __/ | | | | | (_) | $
|_|  |_|  \___| |_| |_|  \___/  $
                                $
                                $
$