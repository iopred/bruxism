package triviaplugin

import "strings"

func init() {
	mirc := `80s Films: "In the Southeast, they say if you want to go to heaven, you have to change planes in Atlanta."*Accidental Tourist
80s Films: A ---------- to a Kill*View
80s Films: A ---------- to India*Passage
80s Films: Andy McCarthy does Robby Lowe's Mom, and she's Jackie Bissette!*Class
80s Films: Baby ----------*Boom
80s Films: Big ---------- in Little China*Trouble
80s Films: Bill Murray, post-SNL, pre-Groundhog. Think spaghetti*Meatballs
80s Films: Black ----------*Rain
80s Films: Dead ---------- Don't Wear Plaid*Men
80s Films: Dirty ----------*Dancing
80s Films: Driving Miss ----------*Daisy
80s Films: First ----------*Blood
80s Films: For Your ---------- Only*Eyes
80s Films: Guy really gets into playing videogames*Tron
80s Films: Hmm. Let's make Chevy Chase a spy!*Fletch
80s Films: Less than ----------*Zero
80s Films: Mr. ----------*Mom
80s Films: My Left ----------*Foot
80s Films: No Way ----------*Out
80s Films: Proof that there's life in the afterlife. And Geena Davis's lips*Beetlejuice
80s Films: Ricki Lake. Before her television show*Hairspray
80s Films: Say ----------*Anything
80s Films: Sex, Lies and ----------*Videotape
80s Films: Some Kind of ----------*Wonderful
80s Films: The Big ----------*Chill
80s Films: The bitch is back, and it's not Sigourney*Aliens
80s Films: The Evil ----------*Dead
80s Films: The Last ---------- of Christ*Temptation
80s Films: The paean to late nights and teens in the 1950s*Diner
80s Films: The Right ----------*Stuff
80s Films: The ---------- Brothers*Blues
80s Films: The ---------- Crystal*Dark
80s Films: The ---------- Guy*Lonely
80s Films: This is ---------- Tap*Spinal
80s Films: Three Men and a ----------*Baby
80s Films: Urban ----------*Cowboy
80s Films: War----------*Games
80s Films: ---------- Boys*Bad
80s Films: ---------- By Me*Stand
80s Films: ---------- Eye*Cat's
80s Films: ---------- of Dreams*Field
80s Films: ---------- of Fire*Chariots
80s Films: ---------- of the Universe*Masters
80s Films: ---------- of War*Casualties
80s Films: ---------- Taste*Bad
80s Films: ---------- That Girl?*Who's
80s Films: ---------- the 13th*Friday
A "double sheet bend" is a type of what*knot
A block of compressed coal dust used as fuel*briquette
A boat or raft with two parallel hulls*catamaran
A bone specialist is a ----------*osteopath
A capital D is the Roman numeral for which number*five hundred
A character named 'Spearchucker Jones' was deleted from this famous American television show's cast after only five episodes?*MASH
A charge of dwai is for what*driving while ability impaired
A complex alcohol constituent of all animal fats and oils*cholesterol
A dolphin can remember a specific ---------- better than a human*tone
A famous RAH novel, as well as a number believed to be cursed*the number of the beast
A Group of Cattle is called a?*herd
A group of ducks is called*brace
A Group of Lion is called a?*pride
A herb or drug described as 'haemostatic' performs which effect*stops bleeding
A hoop worn under skirts is called a what*farthingale
A kipper is what type of smoked fish*herring
A large French country house*chateau
A light canvas shoe with a plaited sole*espadrille
A magnum of champagne is how many litres*1.5
A male singer whose sexual organs have been modified is known as a what*castrato
A person refusing to join a strike*blackleg
A sadhu is a holy man in which country*india
A sailor who has not yet crossed the equator is referred to by what name*Pollywog
A salad containing diced apple, celery, walnuts and mayonnaise is known as what*waldorf salad
A ships officer in charge of equipment and crew*boatswain
A short legged hunting dog*basset
A short women's jacket without fastenings*bolero
A small crown*coronet
A small pickled cucumber*gherkin
A terrapin is a type of ----------*Turtle
A very tall center and a real ladies man*wilt the stilt
A wild ox*bison
A word like 'NASA' formed from the initials of other words is a(n) ----------*Acronym
Aarchie Moore, was world champion in what sport from 1952 1962*boxing
About which Prime Minister was it said "He could never see a belt without hitting below it"*lloyd george
According to John Aubrey's Brief Lives , what card game did the English poet, Sir John Suckling, invent in 1630*cribbage
According to superstition, what do you do when you stub the toes on your right foot*make a wish
According to the title of a famous novel, there are how many 'Years of Solitude*one hundred
According to tradition, which animals desert a sinking ship*rats
According to U.S. law, what may not be granted on a useless invention, on a method of doing business, on mere printed matter, or on a device or machine that will not operate*patent
Acronym for quasi-stellar radio source, any of the blue, starlike objects that are strong radio emitters and the spectra of what exhibit a strong red shift*quasar
Acronym Soup: JVC*japan victor company
Acronym Soup: ACK*acknowledgement
Acronym Soup: ANSI*american national standards institute
Acronym Soup: AOL*america on line
Acronym Soup: AWGTHTGTTA*are we going to have to go through this/that again
Acronym Soup: BAC*by any chance
Acronym Soup: BAK*back at keyboard
Acronym Soup: BBL*be back later
Acronym Soup: BOS*boyfriend over shoulder
Acronym Soup: BRT*be right there
Acronym Soup: BST*british summer time
Acronym Soup: BTOBD*be there or be dead
Acronym Soup: BWL*bursting with laughter
Acronym Soup: COD*cash on delivery
Acronym Soup: DBN*doing business
Acronym Soup: DFM*don't flame me
Acronym Soup: DIIK*damned if i know
Acronym Soup: DOB*date of birth
Acronym Soup: DWIM*do what i mean
Acronym Soup: E2EG*ear to ear grin
Acronym Soup: EMSG*email message
Acronym Soup: FIRST*forum of incident response and security teams
Acronym Soup: FOAF*friend of a friend
Acronym Soup: FOD*finger of death
Acronym Soup: FURTB*full up ready to burst
Acronym Soup: GAFIA*get away from it all
Acronym Soup: GIWIST*gee i wish i'd said that
Acronym Soup: GMT*greenwich mean time
Acronym Soup: HTTP*hyper text transfer protocol
Acronym Soup: IAAD*i am a doctor
Acronym Soup: IAE*in any event
Acronym Soup: IBTD*i beg to differ
Acronym Soup: IMBO*in my biased opinion
Acronym Soup: IWALY*i will always love you
Acronym Soup: JIC*just in case
Acronym Soup: JTLYK*just to let you know
Acronym Soup: JTUSK*just thought you should know
Acronym Soup: JTYMLTK*just thought you might like to know
Acronym Soup: L8TRZ*laters
Acronym Soup: LHU*lord help us
Acronym Soup: LLAP*live long and prosper
Acronym Soup: LMC*lost my connection
Acronym Soup: LoTR*lord of the rings
Acronym Soup: LSFIAB*like shooting fish in a barrel
Acronym Soup: LYLAB*love you like a brother
Acronym Soup: MD*mailed
Acronym Soup: MHM*members helping members
Acronym Soup: MHOTY*my hat's off to you
Acronym Soup: MIPS*meaningless information per second
Acronym Soup: MORF*male or female
Acronym Soup: NAGI*not a good idea
Acronym Soup: NDM*no disrespect meant
Acronym Soup: NIDWTC*no i don't want to chat
Acronym Soup: NIH*not invented here
Acronym Soup: NOOTO*nothing out of the ordinary
Acronym Soup: NYM*new york minute
Acronym Soup: OTT*over the top
Acronym Soup: PAW*parents are watching
Acronym Soup: PDQ*pretty damn quick
Acronym Soup: PIMP*pee in my pants
Acronym Soup: PLMKO*please let me know ok
Acronym Soup: PMBI*pardon my butting in
Acronym Soup: PTO*please turn over
Acronym Soup: Q*queue
Acronym Soup: QPQ*quid pro quo
Acronym Soup: RA*red alert
Acronym Soup: ROFLMHO*rolling on floor laughing my head off
Acronym Soup: ROTFLAS*rolling on the floor laughing and snorting
Acronym Soup: RTS*read the screen; real time strategy
Acronym Soup: SASS*short attention span society/syndrome
Acronym Soup: SMOFF*serious mode off
Acronym Soup: SMOP*small matter of programming
Acronym Soup: SNR*signal to noise ratio
Acronym Soup: SOHB*sense of humour bypass
Acronym Soup: ST-DS9*star trek deep space 9
Acronym Soup: SUFID*screwing up face in disgust
Acronym Soup: TBE*to be expected
Acronym Soup: TIATLG*truly i am the living god
Acronym Soup: TMIKTLIU*the more i know the less i understand
Acronym Soup: TN*telnet
Acronym Soup: TRDMC*tears running down my cheeks
Acronym Soup: UUCP*unix-tounix copy
Acronym Soup: VR*virtual reality
Acronym Soup: W/*with
Acronym Soup: WAIS*wide area information server
Acronym Soup: WC*way cool
Acronym Soup: WYM*what you mean?
Acronym Soup: YCLIU*you can look it up
Acronym Soup: YGTBK*you've got to be kidding
Acronym Soup: YHBW*you have been warned
Acronym Soup: YKYATP*you know you're a tired parent
Actor Arnold Schwarzenegger bought the first Hummer manufactured for civilian use when*1992
Actor ---------- Borgnine*Ernest
Actually caused by layers of hot air refracting sunlight*mirage
Advertising film which is informative and purportedly objective*infomercial
After whom is the month of July named*julius caesar
Air is 21% oxygen, 78% ----------, and 1% other gases*nitrogen
Alberta's shield on the coat of arms, bears the cross of*saint george
Algebra: Define the value of X: -10x - 19 = 19 - 8x*minus nineteen
Algebra: Define the value of X: 99 + 5x = 2000 - 5x - 911*ninety nine
All Hebrew orignating names that end with the letters "el" have something to do with what*god
Alphabetically speaking, which is the last of the 26 Irish counties. Most people say Wexford, but they're wrong*.wicklow
Although his career was snuffed out in the same plane crash that killed Buddy Holly, which east L.A kid had a memorable top ten hit about his girlfriend Donna*Richie Valens
Always ----------*coca cola
America's country's first commercial oil well was located in what state*pennsylvania
American indians used beads as currency. What was it called*wampum
American inventor and teacher of the deaf, most famous for his invention of the telephone*alexander graham bell
An addition to a will is called a*codicil
An America reindeer*caribou
An animal is a fish if it has ----------*gills
An area seperating potential belligerents*buffer zone
An underground layer of water filled rock is called an*aquifer
Animal Trivia: ---------- and short-tailed shrews get by on only two hours of sleep a day*elephants
Animal Trivia: ---------- are the only animals born with horns. Both males and females are born with bony knobs on the forehead*giraffes
Animal Trivia: ---------- can swim for a 1/2 mile without resting, and they can tread water for 3 days straight*rats
Animal Trivia: ---------- feel safest when they are crowded together, hundreds in a group*flamingoes
Animal Trivia: ---------- have the best eyesight of any breed of dog*greyhounds
Animal Trivia: ---------- herds post their own sentries. When danger threatens, the sentry raises its trunk and though it may be as far as a half-mile away, the rest of the herd is instantly alerted. how this communication takes place is not understood*elephant
Animal Trivia: ---------- may travel great distances on their migrations. The Arctic tern travels from the top of the world, the Arctic - to the bottom, the Antarctic. Round trip in a single year: 25,000 miles in all*birds
Animal Trivia: A camel can shut its nostrils during a ----------*desert sandstorm
Animal trivia: a camel with one hump is a dromedary, while a camel with two humps is a ----------*bactrian
Animal Trivia: A donkey is an "ass", but an ass is not always a donkey. The word "ass" refers to several hoofed mammals of the genus Equus, including the ----------*onager
Animal Trivia: A garter snake can give birth to ----------*85 babies
Animal Trivia: A giant Pacific ---------- can fit its entire body through an opening no bigger than the size of its beak*octopus
Animal Trivia: A male ---------- that has been neutered is known as a "wether."*goat
Animal Trivia: A mole can dig a tunnel ---------- feet long in one night*three hundred
Animal Trivia: A young pigeon that has not yet flown is a ----------*squab
Animal Trivia: All porcupines float in ----------*water
Animal Trivia: Baby beavers are called kits or ----------*kittens
Animal Trivia: Because its tongue is too short for its beak, the ---------- must juggle its food before swallowing it*toucan
Animal Trivia: Because the natural habitat of ---------- is of little use to man - the alkaline African lake waters support few fish and cannot be used for human consumption or irrigation - and also because their resting areas are typically inaccessible, the birds are rarely disturbed, unlike other African wild birds*flamingos
Animal Trivia: Boxers were named after their habit of playing. At the beginning of play with another dog, a Boxer will stand on his hind legs and bat at his opponent, appearing to "box" with his ----------*front paws
Animal Trivia: Bull giraffes forage higher in trees than cow giraffes which reduces food competition between the sexes. Long-legged giraffes walk with the limbs on one side of the body lifted at the same time. This gait is called a pace and allows a longer stride which saves ----------*steps and energy
Animal Trivia: Cats are the only domestic animals that walk directly on their ----------, not on their paws. This method of walking is called "digitigrade". When cats scratch furniture, it isn't an act of malice. They are actually tearing off the ragged edges of the sheaths of their talons to expose the new sharp ones beneath*claws
Animal Trivia: Elephants, lions, and camels roamed ---------- 12,000 years ago*alaska
Animal Trivia: Every ----------, there is a peak in Canada wildlife population, especially among the muskrats, red fox, skunks, mink, lynx, and rabbits. The population of grasshoppers of the world tends to rise and fall rhythmically in 9.2-year cycles*9.6 years
Animal Trivia: From crocodile farms, Australia exports about 5,000 crocodile skins a year. Most go to Paris, where a crocodile purse can sell for more than ----------*$10,000
Animal Trivia: If they are well treated, camels in captivity can live to the age of ----------*fifty
Animal Trivia: It seems to biologists that, unlike their humpback whale relatives whose underwater song evolves from year to year, killer whales retain individual ---------- unchanged over long periods, possibly even for life*dialects
Animal Trivia: It takes an average of 345 squirts to yield a gallon of milk from a cow's ----------*udder
Animal Trivia: Javelinas are free-ranging, yet territorial animals that travel in small herds. One of the reasons they travel in numbers is so they can huddle to stay warm - they don't handle cold well and can ----------*freeze to death quickly
Animal Trivia: Milk snakes lay about 13 eggs - in piles of animal ----------*manure
Animal Trivia: Monkeys will not eat red meat or ----------*butter
Animal Trivia: Pink elephants? In regions of India where the soil is red, elephants take on a permanent pink tinge because they regularly spray dust over their bodies to protect themselves against ----------*insects
Animal Trivia: The ---------- snake found in the state of Arizona is not poisonous, but when frightened, it may hiss loudly and vibrate its tail like a rattlesnake*gopher
Animal Trivia: The ---------- whale is the mammal with the heaviest brain - about six times heavier than a human's*sperm
Animal Trivia: The Alaskan ---------- is the largest deer of the New World. It attains a height at the withers in excess of 7 feet and, when fully grown, weighs up to 1,800 pounds*moose
Animal Trivia: The armor of the ---------- is not as tough as it appears. It is very pliable, much like a human fingernail*armadillo
Animal Trivia: The average adult ---------- weighs 21 pounds*raccoon
Animal Trivia: The average giraffe's ---------- is two or three times that of a healthy man*blood pressure
Animal Trivia: The average porcupine has more than 30,000 quills. Porcupines are excellent swimmers because their quills are hollow and serve as pontoons to keep them ----------*afloat
Animal Trivia: The bat is the only mammal that can ----------*fly
Animal Trivia: The Dalmatian dog is named for the Dalmatian Coast of ----------, where it is believed to have been originally bred*croatia
Animal Trivia: The electric eel lives in the Amazon River and its tributaries in South America. The rivers churn up a lot of mud and the eels cannot see well in them. Two less powerful electric fish are the electric catfish and ray. Electric rays live in warm ocean water, and they can give off a charge of sufficient force to stun a human. The biggest electric ray, the Atlantic torpedo ray, can weigh ---------- pounds*two hundred
Animal Trivia: The fastest animal on four legs is the ----------, which races at speeds up to 70 miles per hour in short distances. it can accelerate to 45 miles per hour in two seconds*cheetah
Animal Trivia: The female condor lays a single egg once every ----------*two years
Animal Trivia: The female king crab incubates as many as 400,000 young for 11 months in a brood pouch under her ----------*abdomen
Animal Trivia: The fur of the vicuna, a small member of the camel family which live in the Andes mountains of Peru, is so fine that each hair is less than two-thousandths of an inch. The animal was considered sacred by the Incas, and only royalty could wear its ----------*fleece
Animal Trivia: The hummingbird is the only bird that can ----------*fly backwards
Animal Trivia: The largest species of seahorse measures ----------*eight inches
Animal Trivia: The leech has 32 brains - 31 more than a ----------*human
Animal Trivia: The leech will gorge itself up to ---------- its body weight and then just fall off its victim*five times
Animal Trivia: The life expectancy of the average mockingbird is ----------*ten years
Animal Trivia: The maximum life span of ---------- has been documented to be over 200 years in exceptional cases. The average life span of the large colorful fish, however, is 25 to 35 years*koi
Animal Trivia: The Rufous is the only species of hummingbird to nest in Alaska. They migrate 2,000 miles to Mexico each winter, and then back to Alaska in the ----------*spring
Animal Trivia: The sea cucumber, a purplish-brown creature covered with ----------, has a unique defense strategy. When attacked, it throws out sticky threads from its mouth, which entangles its enemy. The sea cucumber can then quickly escape*warts
Animal Trivia: The smell of a ---------- can be detected by a human a mile away*skunk
Animal Trivia: The tarantula spends most of its life within its burrow, which is an 18-inch vertical hole with an inch-wide opening. When male tarantulas are between the ages of 5 to 7 years, they leave the burrow in search of a female, usually in the early fall. This migration actually signals the end of their life cycle. The males mate with as many females as they can, and then they die around mid-----------*november
Animal Trivia: The three-toed ---------- of tropical America can swim easily, but it can only drag itself across bare ground*sloth
Animal Trivia: The two ---------- of a dolphin's brain work independently. For 8 hours, the entire brain is awake. The left side then sleeps for 8 hours. When it wakes up, the right side sleeps for 8 hours. Thus, the dolphin gets 8 hours of sleep without ever having to stop physically*hemispheres
Animal Trivia: There are about 40 different ---------- in a birds wing*muscles
Animal Trivia: There is no mention of cats or rats in the ----------*bible
Animal Trivia: Though human noses have an impressive 5 million olfactory cells with which to smell, sheepdogs have 220 million, enabling them to smell 44 times better than ----------*man
Animal Trivia: Unlike dolphins, porpoises are not very ----------*sociable
Animal Trivia: Wandering ---------- spread their wings, clack bills, and shake heads in a ritual dance. Bonds between courting birds may last the whole of a 50-year lifetime*albatrosses
Animal Trivia: When young abalones feed on red seaweed, their shells turn ----------*red
Animal Trivia: While many people believe that a camel's humps are used for water storage, they are actually made up of ----------. The hump of a well-rested, well-fed camel can weigh up to eighty pounds*fat
Animals that once existed and exist no more, are called ----------*Extinct
Annapolis & Minneapolis contain the suffix "polis", which in Greek means ----------*city
Anothe name for an artists workshop or studio*ateller
Another name for guardian angels is*watchers
Another name for phencyclidine hydrochloride*angel dust
Another name for wood alcohol is*methanol
Anthocyanins are compounds which produce what*colors
Anti tank rocket launcher*bazooka
Any of a large group of chemicals almost exclusively organic in nature, used for the coloring of textiles, inks, food products, & other substances*dyes
Any of various scientific recording devices designed to register a person's bodily responses to being questioned?*polygraph
Apart from Gottfried Leibniz, which famous scientist developed the many techniques of calculus in mathematics*isaac newton
Approx 800 people died at a firework display in Paris in what year*1770
Araucaria or Chile Pine has a more common name, what is it*monkey puzzle tree
Architectural style developed in the Eastern Empire*byzantine
As clear as a ----------*Bell
As pretty as a ----------*picture
As what did Kotex first manufactured in WWI*bandages
As what is Beethoven's piano sonata in C-sharp minor more commonly known*the moonlight sonata
As what is Hungary also known*magyar
As what is the Devonian period also known*age of fish
As what was John F. Kennedy airport formerly known*idlewild
As what was Louis XIV also known*sun king
As what was sony's video recorder known*beta-max
Ashord/V. Simpson)*whitney houston
At what theme park are the Looney Toons associated with*Six Flags
At which university did the poet Philip Larkin work as a librarian*hull
Athropod with worm like body and many legs*centipede
Atlanta 1996 Olympics: This country's medal tally was: 0 Gold, 1 Silver, 0 Bronze, 1 in Total*bahamas
Atlanta 1996 Olympics: This country's medal tally was: 0 Gold, 1 Silver, 0 Bronze, 1 in Total*latvia
Atlanta 1996 Olympics: This country's medal tally was: 1 Gold, 0 Silver, 0 Bronze, 1 in Total*burundi
Atlanta 1996 Olympics: This country's medal tally was: 1 Gold, 0 Silver, 0 Bronze, 1 in Total*syria
Atlanta 1996 Olympics: This country's medal tally was: 1 Gold, 0 Silver, 1 Bronze, 2 in Total*thailand
Atlanta 1996 Olympics: This country's medal tally was: 1 Gold, 1 Silver, 0 Bronze, 2 in Total*croatia
Atlanta 1996 Olympics: This country's medal tally was: 2 Gold, 4 Silver, 2 Bronze, 8 in Total*sweden
Atlanta 1996 Olympics: This country's medal tally was: 3 Gold, 0 Silver, 1 Bronze, 4 in Total*ireland
Atlanta 1996 Olympics: This country's medal tally was: 4 Gold, 1 Silver, 1 Bronze, 6 in Total*turkey
Atlanta 1996 Olympics: This country's medal tally was: 4 Gold, 3 Silver, 2 Bronze, 9 in Total*czech republic
Atlanta 1996 Olympics: This country's medal tally was: 4 Gold, 4 Silver, 6 Bronze, 14 in Total*spain
Aussie Slang: Ankle biter*young child
Aussie Slang: Battler*someone who works hard
Aussie Slang: Blue ass fly*someone doing something very fast
Aussie Slang: Bull dust*a lie
Aussie Slang: Dacks*trousers or shorts
Aussie Slang: Do the lolly*to get very angry
Aussie Slang: Dunny*toilet
Aussie Slang: Earbashing*someone talking to you for a long time
Aussie Slang: Eat a horse, and chase the jockey*you are very very hungry
Aussie Slang: Fair dinkum*honest, genuine
Aussie Slang: Footy*ozzie rules football
Aussie Slang: Get nicked*to get caught
Aussie Slang: Grot*person who is very dirty or untidy
Aussie Slang: Holy-dooly*an expression of surprise
Aussie Slang: In your dreams*telling someone that it isn't going to happen
Aussie Slang: Kiwi*someone who lives in new zealand
Aussie Slang: Knocker*someone who makes derogatory remarks
Aussie Slang: Oldies*parents
Aussie Slang: Paddock*grazing field for sheep and cattle
Aussie Slang: Rack off*told to go away angrily
Aussie Slang: Sanga*sandwich
Aussie Slang: Shonky*poor quality
Aussie Slang: Stone the crows*exclamation of amazement
Aussie Slang: Tinnie*can of beer or a aluminum boat
Aussie Slang: Wowser*a killjoy
Aussie Slang: Wrapped*excited about something good that has happened
Authority charged with the disposition of legal actions involving children*juvenile court
Authors/poets: Who wrote Illiad*alexander pope
Authors: The Silence of the Lambs*thomas harris
Authors: Who wrote A Bridge to Far*cornelius ryan
Authors: Who wrote A Christmas Carol*charles dickens
Authors: Who wrote A Clockwork Orange*anthony burgess
Authors: Who wrote As I Lay Dying*william faulkner
Authors: Who wrote Cardinal of the Kremlin*tom clancy
Authors: Who wrote Catch 22*joseph heller
Authors: Who wrote Chitty Chitty Bang Bang*ian fleming
Authors: Who wrote Cry the Beloved Country*alan paton
Authors: Who wrote Deptford Trilogy*robertson davies
Authors: Who wrote El Gringo*carlos fuentes
Authors: Who wrote Fahrenheit 451*ray bradbury
Authors: Who wrote Ice Palace*edna furber
Authors: Who wrote Interview with a Vampire*anne rice
Authors: Who wrote Ironweed*William Kennedy
Authors: Who wrote Islands in the Stream*ernest hemingway
Authors: Who wrote Jane Eyre*charlotte bronte
Authors: Who wrote Lightning*Dean Koontz
Authors: Who wrote Little Country*charles de lint
Authors: Who wrote Lolita*vladimir nabokov
Authors: Who wrote Movable Feast*ernest hemingway
Authors: Who wrote Murders of Rue Morgue*edgar allan poe
Authors: Who wrote Pilgrims Progress*john bunyan
Authors: Who wrote Portnoy's Complaint*philip roth
Authors: Who wrote Puppet Masters*robert heinlein
Authors: Who wrote Silas Marner*george elliot
Authors: Who wrote Single and Single*john le carre
Authors: Who wrote Someplace to Be Flying*charles de lint
Authors: Who wrote Sound and the Fury*william faulkner
Authors: Who wrote Sum of All Fears*tom clancy
Authors: Who wrote The Runaway Jury*john grisham
Authors: Who wrote Travels*marco polo
Authors: Who wrote Vet in Harness*james herriot
Authors: Who wrote Wizard of Oz*l frank baum
Authors: Zorba the Greek*nikos kazantzakis
Baby Names Beginning With "A": Meaning: Good*aggie
Baby Names Beginning With "A": Meaning: Healer*althea
Baby Names Beginning With "A": Meaning: High, Exalted*aram
Baby Names Beginning With "A": Meaning: Island*avalon
Baby Names Beginning With "A": Meaning: Last Daughter*audi
Baby Names Beginning With "A": Meaning: Love*amorina
Baby Names Beginning With "A": Meaning: Lovely*ah
Baby Names Beginning With "A": Meaning: Loving, Kind-Hearted*aloha
Baby Names Beginning With "A": Meaning: Noble, Kind*ada
Baby Names Beginning With "A": Meaning: Pleasant*africa
Baby Names Beginning With "A": Meaning: Prayer*atira
Baby Names Beginning With "A": Meaning: Protector of Men*alejandro
Baby Names Beginning With "A": Meaning: Pure*azra
Baby Names Beginning With "A": Meaning: Strong as an Eagle*arnold
Baby Names Beginning With "A": Meaning: To Tie*avongara
Baby Names Beginning With "A": Meaning: Unity, One*ace
Baby Names Beginning With "A": Meaning: Vigilant, Watchful*argus
Baby Names Beginning With "A": Meaning: Wisdom*akili
Baby Names Beginning With "A": Meaning: Youthful*atalo
Baby Names Beginning With "B": Meaning: A Broad Lea, Meadow*bradley
Baby Names Beginning With "B": Meaning: A Flame*blaze
Baby Names Beginning With "B": Meaning: Like a King*basil
Baby Names Beginning With "B": Meaning: Messenger, Friend*bud
Baby Names Beginning With "B": Meaning: Of the City of Bubastis*bast
Baby Names Beginning With "B": Meaning: Son of Evan*bevan
Baby Names Beginning With "B": Meaning: Son of My Right Hand*benjamin
Baby Names Beginning With "B": Meaning: Strong, Virtuous, Honorable*breanna
Baby Names Beginning With "B": Meaning: Victorious*berenice
Baby Names Beginning With "B": Meaning: Warrior*boris
Baby Names Beginning With "B": Meaning: Yellow Haired*boyd
Baby Names Beginning With "C": Meaning: A Marsh or an herb*curry
Baby Names Beginning With "C": Meaning: Gracious and Womanly*carrieann
Baby Names Beginning With "C": Meaning: Harbor*chelsa
Baby Names Beginning With "C": Meaning: Hound of the plain*conway
Baby Names Beginning With "C": Meaning: Ladylike*cirila
Baby Names Beginning With "C": Meaning: Little Maiden*coralie
Baby Names Beginning With "C": Meaning: Maiden*corinna
Baby Names Beginning With "C": Meaning: Manly, Full Grown*charles
Baby Names Beginning With "C": Meaning: Near the creek*creighton
Baby Names Beginning With "C": Meaning: Rocky Land or Singer*chantel
Baby Names Beginning With "C": Meaning: Song*carmen
Baby Names Beginning With "C": Meaning: Soothsayer*cybil
Baby Names Beginning With "C": Meaning: Strong Man*cairbre
Baby Names Beginning With "C": Meaning: To Life*chaim
Baby Names Beginning With "C": Meaning: To the Castle*castel
Baby Names Beginning With "C": Meaning: Very Beautiful*calixte
Baby Names Beginning With "C": Meaning: Victorious People*caelan
Baby Names Beginning With "C": Meaning: Warrior*calhoun
Baby Names Beginning With "D": Meaning: A Gift*donato
Baby Names Beginning With "D": Meaning: Hidden Nook*darnell
Baby Names Beginning With "D": Meaning: Leader of the People*didrika
Baby Names Beginning With "D": Meaning: Longed For*desana
Baby Names Beginning With "D": Meaning: Mouth of a River*delta
Baby Names Beginning With "D": Meaning: Of God*dominique
Baby Names Beginning With "D": Meaning: Of the Devil*desdemona
Baby Names Beginning With "D": Meaning: Powerful, Rich Ruler*dick
Baby Names Beginning With "D": Meaning: Queenly*daria
Baby Names Beginning With "D": Meaning: Roaming*dessa
Baby Names Beginning With "D": Meaning: The Beginning*davu
Baby Names Beginning With "E": Meaning: Gracious Protector*esme
Baby Names Beginning With "E": Meaning: Greek Mythological Figure*eurydice
Baby Names Beginning With "E": Meaning: Happy*edith
Baby Names Beginning With "E": Meaning: Hazelnut*evelia
Baby Names Beginning With "E": Meaning: Honored, Distinguished*efrat
Baby Names Beginning With "E": Meaning: Justice*euridice
Baby Names Beginning With "E": Meaning: Life*eshe
Baby Names Beginning With "E": Meaning: Light*elaine
Baby Names Beginning With "E": Meaning: Man*enos
Baby Names Beginning With "E": Meaning: Noble in Counsel*ethelda
Baby Names Beginning With "E": Meaning: Noble*elgin
Baby Names Beginning With "E": Meaning: Peace, Enlightened*eron
Baby Names Beginning With "E": Meaning: Son of Ed*edison
Baby Names Beginning With "E": Meaning: Very Fruitful*ephraim
Baby Names Beginning With "F": Meaning: A Field*fell
Baby Names Beginning With "F": Meaning: Leafy Branch*fronde
Baby Names Beginning With "F": Meaning: Of the Woods, Forest*forrest
Baby Names Beginning With "F": Meaning: The Hollow*floyd
Baby Names Beginning With "F": Meaning: White Shoulder*fennella
Baby Names Beginning With "G": Meaning: A Dove*giona
Baby Names Beginning With "G": Meaning: Grace of God*grace
Baby Names Beginning With "G": Meaning: Loving*guillermina
Baby Names Beginning With "G": Meaning: Merry, Happy*gaye
Baby Names Beginning With "G": Meaning: My Joy, Rejoice*gili
Baby Names Beginning With "G": Meaning: Son of the Grey-Haired One*grayson
Baby Names Beginning With "H": Meaning: Hollow in the Valley*holden
Baby Names Beginning With "H": Meaning: Lamb*hamal
Baby Names Beginning With "H": Meaning: Meadow on the Cliff*hanley
Baby Names Beginning With "H": Meaning: Rabbit*hazeka
Baby Names Beginning With "H": Meaning: Shining Brightly*haruki
Baby Names Beginning With "H": Meaning: Spacious Meadow*harley
Baby Names Beginning With "H": Meaning: Star*hoshiko
Baby Names Beginning With "H": Meaning: The Flower*hyacinth
Baby Names Beginning With "I": Meaning: Innocent*ince
Baby Names Beginning With "I": Meaning: Mother*ina
Baby Names Beginning With "I": Meaning: Mythological Creature*iphigenia
Baby Names Beginning With "I": Meaning: Only Son*iggi
Baby Names Beginning With "I": Meaning: Snow*istas
Baby Names Beginning With "J": Meaning: God is Willing*joel
Baby Names Beginning With "J": Meaning: Happy*jabulani
Baby Names Beginning With "J": Meaning: Happy*jovita
Baby Names Beginning With "J": Meaning: Holy Man*jeroen
Baby Names Beginning With "J": Meaning: Jade*jadzia
Baby Names Beginning With "J": Meaning: Prominent*jael
Baby Names Beginning With "J": Meaning: Supplanter*jocelin
Baby Names Beginning With "J": Meaning: Youthful*julius
Baby Names Beginning With "K": Meaning: A Church*kirby
Baby Names Beginning With "K": Meaning: Golden*kin
Baby Names Beginning With "K": Meaning: Handsome, Beautiful*kevin
Baby Names Beginning With "K": Meaning: Harp*koto
Baby Names Beginning With "K": Meaning: Hunter*kacela
Baby Names Beginning With "K": Meaning: Maiden*koren
Baby Names Beginning With "K": Meaning: Perfect*kamil
Baby Names Beginning With "K": Meaning: Poet*kavi
Baby Names Beginning With "K": Meaning: Pure*katrina
Baby Names Beginning With "K": Meaning: Rejoicer, Waterfall Pool*kaelin
Baby Names Beginning With "K": Meaning: Rosy Reflection in the Sky*kawena
Baby Names Beginning With "K": Meaning: Royal*kennedy
Baby Names Beginning With "K": Meaning: Song, Songstress*karmina
Baby Names Beginning With "K": Meaning: The Loved One*kendi
Baby Names Beginning With "K": Meaning: Treasure House*kura
Baby Names Beginning With "L": Meaning: Gracious, Poetic*leighanna
Baby Names Beginning With "L": Meaning: Harness Maker*lorimer
Baby Names Beginning With "L": Meaning: Joy, Gladness*letitia
Baby Names Beginning With "L": Meaning: King*loe
Baby Names Beginning With "L": Meaning: Land*lancelot
Baby Names Beginning With "L": Meaning: Lilac*lilia
Baby Names Beginning With "L": Meaning: Linden Trees Near the Water*lindsey
Baby Names Beginning With "L": Meaning: Lioness*liona
Baby Names Beginning With "L": Meaning: Lord*laval
Baby Names Beginning With "L": Meaning: Moon*lucine
Baby Names Beginning With "L": Meaning: Pretty One*linda
Baby Names Beginning With "L": Meaning: Shining*lara
Baby Names Beginning With "L": Meaning: Small Cove*logan
Baby Names Beginning With "L": Meaning: Suave*lizina
Baby Names Beginning With "L": Meaning: Tall One*lang
Baby Names Beginning With "L": Meaning: The Mountain*lamont
Baby Names Beginning With "L": Meaning: Town near the Brook*lynton
Baby Names Beginning With "L": Meaning: Young Girl, Maiden*lassie
Baby Names Beginning With "M": Meaning: Goodness*meged
Baby Names Beginning With "M": Meaning: Great Spring*maxwell
Baby Names Beginning With "M": Meaning: Great*more
Baby Names Beginning With "M": Meaning: Hawaiian form of MARY*mele
Baby Names Beginning With "M": Meaning: Ill-fated Luck*mallory
Baby Names Beginning With "M": Meaning: Lucky*maimun
Baby Names Beginning With "M": Meaning: May*mei
Baby Names Beginning With "M": Meaning: Meadow*maitland
Baby Names Beginning With "M": Meaning: Mine*mio
Baby Names Beginning With "M": Meaning: Miracle Worker*maxima
Baby Names Beginning With "M": Meaning: Moor, Dark Skinned*maurice
Baby Names Beginning With "M": Meaning: Peaceful*mykel
Baby Names Beginning With "M": Meaning: Pearl*margaux
Baby Names Beginning With "M": Meaning: Pearl*margo
Baby Names Beginning With "M": Meaning: Protector from the Sea*meredith
Baby Names Beginning With "M": Meaning: Righteous Way*michi
Baby Names Beginning With "M": Meaning: Sailor*murray
Baby Names Beginning With "M": Meaning: Scholarly Accomplishments*mendel
Baby Names Beginning With "M": Meaning: Single, Alone*monita
Baby Names Beginning With "M": Meaning: Son of Hugh*magee
Baby Names Beginning With "M": Meaning: Son of Ken or Kenna*mckenna
Baby Names Beginning With "M": Meaning: Son of Man*manning
Baby Names Beginning With "M": Meaning: Song-like*melody
Baby Names Beginning With "M": Meaning: The Thrush*mavis
Baby Names Beginning With "M": Meaning: The Tree/Victory*myrtle
Baby Names Beginning With "M": Meaning: Who is like God?*mikko
Baby Names Beginning With "M": Meaning: Wise Little Raccoon*mika
Baby Names Beginning With "M": Meaning: Wished-For Child*mariam
Baby Names Beginning With "N": Meaning: Her Life*nakeisha
Baby Names Beginning With "N": Meaning: New Moon*neona
Baby Names Beginning With "N": Meaning: Pleasant*naomi
Baby Names Beginning With "N": Meaning: Powerful*nero
Baby Names Beginning With "N": Meaning: Saint Worshipper*nevina
Baby Names Beginning With "N": Meaning: Second Wife*nyeki
Baby Names Beginning With "N": Meaning: Son of NEAL*nelson
Baby Names Beginning With "N": Meaning: Standing Tall*nibaw
Baby Names Beginning With "N": Meaning: The Flowers*napua
Baby Names Beginning With "O": Meaning: Golden Woman*orla
Baby Names Beginning With "O": Meaning: Lives beside the Oaks*ogden
Baby Names Beginning With "O": Meaning: Ode, Melodic*odette
Baby Names Beginning With "O": Meaning: Of the Sea*ormanda
Baby Names Beginning With "O": Meaning: Strong, Courageous*ondrea
Baby Names Beginning With "O": Meaning: The Orient, East*orien
Baby Names Beginning With "O": Meaning: To Cross*okal
Baby Names Beginning With "O": Meaning: Venerable*oistin
Baby Names Beginning With "P": Meaning: Lover of Flowers*philantha
Baby Names Beginning With "P": Meaning: Silent Worker (From Penelope)*penny
Baby Names Beginning With "P": Meaning: Small*paola
Baby Names Beginning With "P": Meaning: Stone*petronella
Baby Names Beginning With "Q": Meaning: How much*quant
Baby Names Beginning With "Q": Meaning: Queen or Female Companion*queenie
Baby Names Beginning With "R": Meaning: Grand*ronda
Baby Names Beginning With "R": Meaning: Island*rylan
Baby Names Beginning With "R": Meaning: Link Together*rivka
Baby Names Beginning With "R": Meaning: Over the Red River*redford
Baby Names Beginning With "R": Meaning: Peaceful, Queen*reina
Baby Names Beginning With "R": Meaning: Queen*ranee
Baby Names Beginning With "R": Meaning: Roses*rhoda
Baby Names Beginning With "R": Meaning: Satisfied*reda
Baby Names Beginning With "R": Meaning: Steward*reeves
Baby Names Beginning With "R": Meaning: Wise Protection*raymond
Baby Names Beginning With "S": Meaning: Jane, God is Gracious*sheena
Baby Names Beginning With "S": Meaning: Long, Heavy Nail*spike
Baby Names Beginning With "S": Meaning: Ostrich from Water*sadira
Baby Names Beginning With "S": Meaning: Red*scarlett
Baby Names Beginning With "S": Meaning: Rocky Meadow*stanley
Baby Names Beginning With "S": Meaning: Shipmaster*skipper
Baby Names Beginning With "S": Meaning: Stem of Bamboo*shino
Baby Names Beginning With "S": Meaning: To Stain*sully
Baby Names Beginning With "S": Meaning: Treasure, Prize*sima
Baby Names Beginning With "S": Meaning: Valuable*sterling
Baby Names Beginning With "S": Meaning: Warrior*sloan
Baby Names Beginning With "S": Meaning: Wealthy*sumana
Baby Names Beginning With "S": Meaning: Well-going*sivney
Baby Names Beginning With "S": Meaning: Wisdom*sonia
Baby Names Beginning With "T": Meaning: A Fairy Queen*tania
Baby Names Beginning With "T": Meaning: A measure of land*tate
Baby Names Beginning With "T": Meaning: Land of Owen, Young Soldier*tyrone
Baby Names Beginning With "T": Meaning: Mirror Image*toshi
Baby Names Beginning With "T": Meaning: Mole, Gopher*topo
Baby Names Beginning With "T": Meaning: Nation, Tribe*taifa
Baby Names Beginning With "T": Meaning: Perfect*tamma
Baby Names Beginning With "T": Meaning: River*tyne
Baby Names Beginning With "T": Meaning: Rocky Hill*taran
Baby Names Beginning With "T": Meaning: Sad*trista
Baby Names Beginning With "T": Meaning: The Pure One*thrine
Baby Names Beginning With "U": Meaning: To Arrive*uday
Baby Names Beginning With "U": Meaning: Wolf*ulf
Baby Names Beginning With "V": Meaning: Great*velika
Baby Names Beginning With "V": Meaning: Health or Love*valentina
Baby Names Beginning With "V": Meaning: Life*vito
Baby Names Beginning With "V": Meaning: Longings are Waterfalls*visola
Baby Names Beginning With "V": Meaning: Small*vaughn
Baby Names Beginning With "V": Meaning: Truth*verda
Baby Names Beginning With "V": Meaning: Youthful*verne
Baby Names Beginning With "V": Meaning: Youthful*vernon
Baby Names Beginning With "W": Meaning: Path of a Wolf*wolfgang
Baby Names Beginning With "W": Meaning: Shield*walta
Baby Names Beginning With "W": Meaning: The Path thru the Woods*woodrow
Baby Names Beginning With "W": Meaning: To Wade in Water*wade
Baby Names Beginning With "X": Meaning: WANG  Desire*xi
Baby Names Beginning With "Y": Meaning: Quiet*yoshi
Baby Names Beginning With "Y": Meaning: Son*yaro
Baby Names Beginning With "Y": Meaning: Strength of God*yael
Baby Names Beginning With "Y": Meaning: Violet Flower*yoland
Baby Names Beginning With "Z": Meaning: Growing*zaida
Baby Names Beginning With "Z": Meaning: Present*zavad
Baby Names Beginning With "Z": Meaning: The Lord is my Rock*zuriel
Baby Names Beginning With "Z": Meaning: White Wave*zenevieva
Baby Names Beginning With "Z": Meaning: Zeal*zelia
Back in the day at Walt Disney studios, Walt's brother Ray (yes, Ray) reportedly peddled what to employees on the lot*insurance
Baklava is a form of*dessert
Barbie's nautical-sounding sister*skipper
BaseBall - The Chicago ----------*cubs
Baseball the st louis ----------*cardinals
Baseball: the Baltimore ----------*orioles
Basketball the Boston ----------*celtics
Bat Masterson was a sportswriter for what paper*morning telegraph
Beasley how much does park place cost in monopoly*four hundred fifty dollars
Because metal was scarce during world war ii, of what were the oscars made*wood
Because Moses felt he was "slow of speech and slow of tongue", who often acted as his spokesperson*aaron
Ben Affleck played CT on what 80's science education television show?*Voyage of the Mimi
Benazir Bhutto regained power in 1993 after being ousted how many years before*three
Biko was involved in what protest movement?*Apartheid
Bill justis was a studio musician when he recorded this 'sloppy' instrumental in october 1957*raunchy
Bill Watterson, cartoonist for Calvin & Hobbes, is the first cartoonist to use what word in his cartoon*booger
Biological community of interacting organisms and their physical environment*ecosystem
Bismarck is the capital of ----------*north dakota
Bissau is the capital of ----------*guinea-bissau
Bond: What was the first James Bond film?*Dr. No
Book of the Old Testament, third of the five biblical books called the Pentateuch?*leviticus
Born May 7, 1901, He Starred In This Movie: Dallas - 1950*gary cooper
British rock-music group that rivaled the popularity of the group's early contemporaries, The Beatles*the rolling stones
Briton's say 'tarmac', Americans say ----------*runway
Brothers A terrapin is a type of ----------*Turtle
Bruce Willis, Arnold Schwarzenegger and Sylvester Stallone own which London restaurant*planet hollywood
Buenos Aires is the capital of ----------*argentina
Burn to stop the flow of blood*cauterize
By what Alias does Ferris Bueller get into Chez Luis?*Abe Frohman
By what name is Maurice Micklewhite better known*michael caine
By what name is the bird Troglodytes Troglodytes better known*WREN
By what other name do we know table tennis*ping pong
By who was gerald ford almost assassinated*squeaky fromme
Can a bat stand up*no
Can a platypus see under water*no
Can you swim in the sea of showers*no
Canada is seperated on an imaginary line along the ----------*49th parallel
Capital cities: Finland*helsinki
Carmenta is the roman goddess of ----------*childbirth
Carolyn Weston's novel Poor, Poor Ophelia was the basis for what show*streets of san francisco
Carter what do goldfish lose if kept in dimly lit or running water*colour
Cat stevens 'want's to try to love again but ----------'*the first cut is the
1980s GrabBag : He lost the 1977 NYC mayoral bid before taking successful aim at Albany*mario cuomo
1980s GrabBag: Country that saw 32 of its citizens convicted of spying from 1981-88*united states
1980s GrabBag: First shortstop since Carew to lead all-star voting 2 years in a row*ozzie smith
1980s GrabBag: The last team Tom Seaver tried to pitch for*new york mets
1980s GrabBag: This country's President Zia ul-Haq was killed in a 1988 plane crash*pakistan
1993 The Year: billion of financial aid went to this nation*russia
1993 The Year: This man became South Korea's first civilian leader*kim young sam
19th Cent Art: The two prominent subjects in Sir Edwin Landseer's "Man Proposes, God Disposes"*polar bear
60s: First nation ever to resign from the united nations*indonesia
70s Authors: Born on the Fourth of July*ron kovic
70s Authors: The Friends of Eddie Coyle*george v. higgins
70s Authors: The Summer before the Dark*doris lessing
70s Authors: The Uses of Enchantment*bruno bettleheim
80s: G.D. Searle & Co put this brand sweetener on the market in 1983*nutrasweet
Ad Jingles: They make the very best chocolate*nestle
Ad Slogans: "Just for the taste of it"*diet coke
Ad Slogans: "You will"*att
Ad Slogans: "You'll love the way we fly"*delta airlines
Ads: This product is named for its chief component, muriate of berberine*murine
Advertising: Jhirmack hair products were advertised by this beauty*victoria principal
Advertising: You are advised never to leave home without this*american express
Alcohol: Monastic order that established the California wine industry*franciscan
American Beers - State: Icehouse:*wisconsin
Anatomy: Your ---------- holds your head to your shoulders*neck
Anime: What is the name of the male lead in ----------Vision of Escaflowne----------*van fanel
Anime: What is the name of the most recent Tenchi Muyo series*shin tenchi muyo
Artists Hometowns: Steve Miller*madison
Arts: In what field of study would you find "flying buttresses"*architecture
Asimov Anthony: In the Xanth series, what is our world called*mundania
Astronomy: Does Uranus have an aurora*yes
Astronomy: This cluster of stars is also known as the Seven Sisters*pleiades
Astronomy: What is the name for a group of stars*constellation
Author: How To Win Friends and Influence People*dale carnegie
Authors: The Prince of Tides*conroy
Barbie Dolls: Barbie's "MOD'ern" cousin*francie
Barbie: Barbie variety with suntanned skin*malibu
Barbie: Complete the Barbie outfit name : Here Comes The -----*bride
Barbie: Ken's buddy*allan
Barbie: What Barbie wears when she visits Japan*kimono
Beer: The process of extracting sugar from malt by soaking in water*mashing
Bestsellers: Lake Wobegon is located in this state*minnesota
Bestsellers: Noble House can be found in this British colony*hong kong
Bestsellers: The first horror novel to reach the top, it became a Linda Blair movie*the exorcist
Bestsellers: The shortest bestselling title, by either Hepburn or King*me
Biology: Every human has one of these on their tummies*navel
Books for the Bored: Who maintained law and order in Noddy's Toyland*mr. plod
Books for the Hip Reader: Illusions: The Adventures of a Reluctant Messiah?*richard bach
Books for the Hip Reader: Last of the Really Great Whangdoodles*andrews
Books for the Hip Reader: Men to Match My Mountains?*irving stone
Books: Author of One Police Plaza, 1984 best-seller*caunitz
Books: Little girls are made of sugar, spice, and ---------- ----*everything nice
Books: Men Against the Sea was part two of this trilogy*bounty
Books: The author of "Heidi"*johanna spyri
Books: This event prompted Mailer to write The Naked and the Dead*pearl harbor
Books: This novel inspired the TV series "The Six Million Dollar Man"*cyborg
Books: Wrote The Teachings of Don Juan: A Yaqui Way of Knowledge*castaneda
Booze Grabbag: Spice that a bartender would dust your Brandy Flip with*nutmeg
Booze Names: 1 oz. gin and 1 oz. orange juice*orange blossom
Booze Names: 1 oz. gin, 1/2 oz. dry vermouth, 1/2 oz. sweet vermouth, 1/2 oz. orange juice*bronx cocktail
Booze Names: Vodka, consomme, lemon, tabasco sauce, salt, pepper, celery salt*bullshot
Cars: Combine a Van & a Car & you get this word*caravan
Cars: Sister car of the Nissan Quest*mercury villager
Cars: The original name for this Pontiac car was the Banshee, for its mythicality*firebird
Cars: Volvo's chairman resigned in 1993 in protest of a merger with this automaker*renault
Cartoon Trivia: In what year did both Peanuts and Beetle Bailey first appear*1950
Cartoon Trivia: Name Alley Oop's girl friend*oola
Cartoon Trivia: On what T.V. show could Tom Terrific be found*captain kangaroo
Cartoon Trivia: Tess Trueheart married which plainclothes detective*dick tracy
Cartoon Trivia: What type of plant does Broom Hilda sell*venus flytrap
Cartoon Trivia: What was the name of Speed Racer's car*the mach five
Cartoon Trivia: Who is Snoopy's arch enemy*baron
Cartoon Trivia: Who says, "Th-th-th-that's all folks!"*porky pig
Celebrity Lovers: What is the FIRST NAME of U.S. Vice President Al Gore's wife*tipper
Celebrity: The director of Jaws, Raiders of the Lost Ark*stephen spielberg
Character Creators: Henry Esmond*william thackery
Character Creators: Jane Eyre*charlotte bronte
Chemistry: What is the main component of air*nitrogen
Children's Literature: Children's series that is really a religious allegory. (C.S. Lewis)*chronicles of narnia
Children's Literature: Judy Blume's first children's bestseller*the tales of a fourth grade nothing
Children's Literature: Leonard Wibberly on the Duchy of Grand Fenwick?*the mouse that roared
Chips: Dielectric thickness can be calculated from this electrical measurement*capacitance
Chips: Early MOSFET pioneer, Stanford grad & author of the EE's device bible:*sze
Cigarettes: Cigarette brand, or Robert Guillaume and shrubs*benson and hedges
Civil War: USA Colonel who led the famed bayonet charge down Little Round Top*joshua lawrence chamberlain
Classic Board Games: What large game company's logo contains a spiral*parker brothers
Clive Barker: To the Seerkind, normal people are known as ---------- (Weaveworld)*cuckoos
Clothes: This company's logo is a sailboat*nautica
Clothes: Three-letter clothing outlet, or a space or void*gap
Contemporary Authors: Washington D.C., Myra Breckinridge*gore vidal
Couples: Miss Piggy and ----------*kermit
Couples: Romeo and ----------*juliet
Couples: Sluggo and ----------*nancy
Crime Stories: He played Caspar Gutman in the 1941 John Huston film*sydney greenstreet
Crime Stories: Raymond Chandler's gumshoe*phillip marlowe
Crime Writers: The Hot Rock*donald e. westlake
Cyberpunk: Cyberpunk term for "logging in" to a cyberspace system:*jacking in
Cyberpunk: In this predecessor of Cyberpunk novels, the hero was Guy Montag:*fahrenheit 451
Dandy Candy: "There's no wrong way of eating" this*reeses peanut butter cups
DC Secret Identities: Kay Challis*crazy jane
DC Secret Identities: Maggie Sawyer*maggie sawyer
Definitions: --isms: The theory that man cannot prove the existence of a god*agnosticism
Definitions: -isms: Excessive emphasis on financial gain*commercialism
Definitions: A receptacle for holy water is a(n) ----------*font
Definitions: Any object worn as a charm may be called a(n) ----------*amulet
Definitions: Doraphobia is the fear of ----------*fur
Definitions: Eleutherophobia is a fear of ----------*freedom
Definitions: Legal Terms: A crime more serious than a misdemeanor*felony
Definitions: Legal Terms: The people chosen to render a verdict in a court*jury
Definitions: The study of man and culture is known as ----------*anthropology
Definitions: The study of natural phenomena: motion, forces, light, sound, etc. is called ----------*physics
Definitions: The study of religion is ----------*theology
Definitions: This word is used as the international radio distress call*mayday
Definitions: What is a dried plum called*prune
Definitions: What is the common name for a Japanese dwarf tree*bonsai
Devils Dictionary: A ship big enough to carry two in fair weather, but only one in foul*friendship
Devils Dictionary: A woman with a fine prospect of happiness behind her*bride
Dr Seuss: The Big-hearted Moose*thidwick
Dr Seuss: This elephant hatched and egg and heard a who*horton
Emoticons: (^o^)*happiness
Emoticons: :D*big smile
Famous Canadians: Known for her series of books about Anne Shirley*lucy maud montgomery
Famous Canadians: Portly comedian, Second City alum, and CFL franchise co-owner*john candy
Famous Gills: Montreal University that more ore less fits the category*mcgill university
Famous People: Which actress married Richard Burton twice*elizabeth taylor
Fashion: Model that married David Bowie*iman
Fast Food: Chain with a hat as a logo, makes roast beef burgers among other things*arbys
Fast Food: Lots of people like this brown liquid with fries*gravy
Fast Food: This "Fresh is the taste" chain is -everywhere-*subway
Fictional Detectives: Creator of Perry Mason*erle stanley gardner
Food and Drink : Mustard, ketchup and onions on a hotdog are all ----------*condiments
Food: Mexican dish with minced and seasoned meat packed in cornmeal and corn husks*tamale
French Food AKA: Delicate egg whites baked at a high temperature, literally means "a breath"*souffle
Fun: Cocktails: Cognac (brandy) and white creme de menthe make a(n) ----------*stinger
Games: How many balls are used in a game of snooker in addition to the cue ball*twenty-one
Games: Name the only woman suspect in the game of "Cluedo" who isn't married*miss scarlett
General: At one time, 6 white beads of this Indian currency were worth one penny*wampum
General: Baseballer Joe Schlabotnik's greatest fan*charlie brown
General: Originally made in Nimes, France, this fabric was called serge denimes*denim
General: The 2 months added when the Roman calendar was expanded from 10 to 12 months*january & february
General: The Russian sled drawn by 3 horses abreast*troika
General Knowledge : What is this sign called "&"*ampersand
Generation X Toys: Atari competitor that featured better graphics*intellevision
Generation X Toys: Large plastic animals gobbled marbles in this game*hungry hungry hippos
Geographic Trivia : What is the most populous city in North America*mexico
Geography: In what city is the Smithsonian Institute*washington
Geography: In which city is Red Square*moscow
Geography: In which city is the Coliseum located*rome
Geography: Into what body of water does the Danube River flow*black sea
Geography: Linz, Austria is a leading port on which river*danube
Geography: Madrid and Lisbon are both located near this river*tagus
Geography: Name the continent that consists of a single country*australia
Geography: Name the longest river in Asia*yangtze
Geography: On what island is the Blue Grotto*capri
Geography: On what peninsula are Spain and Portugal located*the iberia n peninsula
Geography: On which river is the Aswan High Dam*nile
Geography: The sun sets in the ----------*west
Geography: What canal connects Lake Ontario and Lake Erie*welland
Geography: What city is the Christian Science Monitor based in*boston
Geography: What country formed the union of Tanganyika and Zanzibar*tanzania
Geography: What English city does the Prime Meridian pass through*greenwich
Geography: What is the capital of Andorra*andorra la vella
Geography: What is the capital of Angola*luanda
Geography: What is the capital of Australia*canberra
Geography: What is the capital of Bhutan*thimphu
Geography: What is the capital of Colorado*denver
Geography: What is the capital of Equatorial Guinea*malabo
Geography: What is the capital of Finland*helsinki
Geography: What is the capital of Iraq*baghdad
Geography: What is the capital of Ireland*dublin
Geography: What is the capital of Kuwait*kuwait city
Geography: What is the capital of Monaco*monaco
Geography: What is the capital of Niger*niamey
Geography: What is the capital of Papua New Guinea*port moresby
Geography: What is the capital of Solomon Islands*honiara
Geography: What is the capital of Switzerland*bern
Geography: What is the capital of Tanzania*dar es salaam
Geography: What is the capital of Tonga*nuku'alofa
Geography: What is the capital of Turkmenistan*ashkhabad
Geography: What is the capital of Tuvalu*fanafuti
Geography: What is the highest mountain in Canada*mt. logan
Geography: What is the largest of the countries in Central America*nicaragua
Geography: What is the official language of Egypt*arabic
Geography: What mountain range separates Europe from Asia*urals
Geography: What U.S. state is known as The Land of 10,000 Lakes*minnesota
Geography: Where is Beacon Street*boston
Geography: Which Central American country extends furthest north*belize
Geography: Which city is known as Motown*detroit
Geography: Which element makes up 2.83% of the Earth's crust*sodium
Geography: Which element makes up 5% of the Earth's crust*iron
Geography: Which Irish city is famous for its crystal*waterford
Geography: Which is the Earth's fourth largest continent*south america
Geography: Which state is the Garden State*new jersey
Geography: Which U.S. city is known as the Biggest Little City in the World*reno
Geography: Which U.S. state borders a Canadian territory*alaska
Geogrpahy: Name the capital city of Utah*salt lake city
Geology: The violet variety of quartz is called ----------*amethyst
Happy Days takes place in this city*milwaukee
History: General Sherman burned this city in 1864*atlanta
History: He ruled Rome when Christ was born*caesar augustus
History: In what year of WW II did Russia declare war on Japan*1945
History: In which country did the Boxer Rebellion take place*china
History: Israel occupied the Golan Heights. Whose territory was it*syria
History: Name the incident in which tea was dumped into the harbour*boston tea party
History: This military attack took place on Dec. 7, 1941*pearl harbour
History: U.S. President, Herbert C. ----------*hoover
History: What was the instrument of execution during the "Reign of Terror"*guillotine
History: Which military battle took place in 1815*waterloo
History: Which president was responsible for the Louisiana Purchase*jefferson
History: Who succeeded Churchill when he resigned in 1955*sir anthony eden
Hitchhiker Guide: Author of the ----------Hitchhiker's Guide to the Galaxy---------- series*douglas adams
Hitchhiker Guide: If you stick it in your ear,it acts as a translator by feeding on brain energy*babel fish
Hitchhiker Guide: Name of the Paranoid Android*marvin
Holidays: Mithraism's (Sun God worship) big day falls on the same day as this holiday*christmas
Hollywood: Charles Laughton played Quasimodo in this epic film*hunchback of notre dame
Hollywood: Film Title: Fahrenheit ---------- (a number)*451
Hollywood: Forrest ---------- liked shrimp*gump
Hollywood: He directed the movie E.T*stephen spielberg
Hollywood: In "Gone With the Wind", Scarlett regains her wealth by investing in what type of business*sawmill
Hollywood: This was the first cartoon talking picture*steamboat willie
Hollywood: What is the name of the Volkswagen in the film, "The Love Bug"*herbie
Hollywood: What was Dorothy's last name in "The Wizard of Oz"*gale
Hollywood: What was the first film directed by Robert Redford*ordinary people
Hollywood: Which character in "Forrest Gump" loved shrimp*bubba
Hollywood: Which planet was the "Planet of the Apes"*earth
Hollywood: Who played Brad Pitt's cop partner in the movie "Seven"*morgan freeman
In what languages except english did Einstuerzende Neubauten record 'blume'*french and japanese
Intl Beers: Prestige Stout:*haiti
Junk Food: Puffy white soft pillows of sugar; good raw or roasted over a campfire*marshmallows
Junk Food: Soft drink with the slogan: "For a new generation"*pepsi
Junk Food: The southern (U.S.) word for these are "goobers"*peanuts
Kids in the Hall: A Mark McKinney character says "I'm ---------- your head!"*crushing
Language: Many Meanings: Fuel, vapor, flattulence, helium. What is it*gas
Languages: What ONE word fits ----------stream; ----------hill; ----------pour*down
Literature: From which Shakespeare play is this line taken: To be or not to*hamlet
Literature: This Shakespearean king was the actual king of Scotland for 17 years*macbeth
Lord of the Rings: From whom did Bilbo obtain The Ring*gollum
Made In Canada: James Labrie's band previous to Dream Theater:*winter rose
Magazines: This magazine used to boast a circulation of 7,777,777*better homes and gardens
Mathematics: The angles inside a square total ---------- degrees*360
Medicine: A bone specialist is a(n) ----------*osteopath
Medicine: A loss of memory is known as ----------*amnesia
Medicine: Hepatitis affects the ----------*liver
Medicine: How many pints of blood does the average human have in his/her body*twelve
Medicine: Name the largest artery in the human body*aorta
Misc Games: How much is the luxury tax (in dollars) in Monopoly*seventy five
Misc Games: Word derived from "shah mat", from the arabic for "the king is dead"*checkmate
Music: The Who's rock musical stars Elton John. It's called ----------*tommy
Mythology: The sea gods had a three-pronged spear called a(n) ----------*trident
Mythology: What's heaven to fallen Norse warriors*valhalla
Name That Celebrity: Author of The Hobbit and The Lord of the Rings*j.r.r. tolkein
Name The Poet: The Branch Will Not Break*james wright
Name Their Job: Harold Stassen*mayor
National Anthems: ...At your feet, two oceans roar for your noble mission*panama
Nature: A one-humped camel is called a ----------*dromedary
Nature: A terrapin is a type of ----------*turtle
Nature: An animal is a fish if it has ----------*gill
Nature: Dogs bark. What do donkeys do*bray
Nature: The fins of which fish are made into a soup*shark
Nature: This animal is the symbol of the U.S. Democratic Party*donkey
Nature: This animal's shell is used to make attractive jewelry*abalone
Nature: This ugly creature has patches of red on his rear-end*mandrill
Nature: What does a camel store in its hump*fat
Nature: What is a male swine called*boar
Nature: What large herbivore sleeps only one hour a night*antelope
Nature: What word is used for a male duck*drake
NetHack: This tool will instantly get your pets around you*magic whistle
Novelty Songs: Type of car outpacing the Cadillac in "Beep Beep"*nash rambler
Original Titles: The Whale by Herman Melville*moby dick
Peanuts Comics: The catcher on the gang's baseball team*schroeder
Peanuts Comics: This is what Marci calls Peppermint Patti*sir
People: Director of the FBI who lived from 1895-1972*j. edgar hoover
People: First female cabinet member*oveta hobby
Phonetic Radio Call Signs: DRKSKY*delta romeo kilo sierra kilo yankee
Pinball: Rudy the dummy says "It's only pinball!" in this 1991 machine*funhouse
Pinball: This mammoth 1979 Atari game featured a pool-ball sized pinball*hercules
Potpourri: First state to secede from the Union in 1861*south carolina
Potpourri: The worlds first drive-in church was in this state*florida
Potpourri: This Greek admiral of Darius I sailed upto the the Indus river in the 5thc*scylax
Quite a Year for Plums----------*bailey white
Quotes: "A cynic is someone who knows the price of everything and the value of nothing*wilde
Quotes: How do I love thee?*elizabeth browning
Religion: Followers of the Unification Church are called ----------*moonies
Religion: Name the holiest day in the Jewish calendar*yom kippur
Say Cheese: Norwegian origin; caramel flavor; sandwich, snack*gjetost
Say Cheese: Probably French origin; tangy, sharp; appetizer, salad, dessert*blue
Say Cheese: Swiss origin; nutty, sharper than Swiss; cooking, dessert*gruyere
Scents: Charlie and Jontue manufacturer*revlon
Sci Fi Authors: Creator of the Meg & Timothy series*madeleine lengle
Sci Fi Authors: Nancy ----------r----------ss*nancy kress
Sci Fi Authors: Ursula Le----------*ursula leguin
Sci Fi: L. Ron Hubbard began writing this series but died before finishing*mission earth
Sci Fi: Publisher who used to write sci-fi short stories, Lester ---------- ----------*del rey
Sci Fi: The number of Rama spacecraft to reach the solar system*three
Science: Ethylene glycol is frequently used in automobiles.. How*anti-freeze
Science: The filament of a regular light bulb is usually made of ----------*tungsten
Science: The name for the Russian equivalent of Skylab is ----------*salyut
Science: The vernal equinox is the beginning of ----------*spring
Second City : Rabat-Sale*morocco
Sherlock Holmes: 'The five orange pips' saw Holmes oppose this racist organization*ku klux klan
Sherlock Holmes: This famous thriller writer was Dr Watson in the '32 film The Sign of Four*ian fleming
Sherlock Holmes: To where does Holmes finally retire?*sussex downs
Smallish lunchtime pizzas from Pizza Hut are called?*personal pan
Snow Crash: Name the book's main character*hiro protagonist
Snow Crash: Name the hacker who is infected with Snow Crash, owner of The Black Sun:*da5id
Snow Crash: What is the language spoken by taxi drivers?*taxilinga
Snow Crash: What is your image or icon called in the metaverse?*avatar
Snow Crash: What security company arrests Y.T. at the beginning of the book?*metacops
Sport: Baseball: The Toronto ----------*bluejays
Sport: Basketball: The Denver ----------*nuggets
Sport: Football: The Baltimore ----------*colts
Sport: Hockey: The Los Angeles ----------*kings
Sport: How many games must you win to win a normal set in tennis*six
Sport: How many players are there on a water polo team*seven
Sport: In which sport is the term "wishbone" used*football
Sport: In which sport is the term, "Hang ten" used*surfing
Sport: What do the letters ERA mean in baseball*earned run average
Sport: What sport has a hooker in a scrum*rugby
Sport: Which is the only position in soccer allowed to handle the ball*goalie
Sports: Where were the 1920 Olympics held*antwerp, belgium
Sports Actors: Besides Field of Dreams, what other baseball movie starred Kevin Costner*bull durham
Stephen King: Maine city which King calls his home*bangor
Stephen King: Short story featuring a home-made, magic computer?*word processor of the gods
Stephen King: What trail of Harold's does Stu Redmen follow cross country in the Stand*snickers bar wrappers
The Bible: The Nebuchadnezzar king was of this nation*babylon
The Bible: This part of King Saul's belongings was displayed in the temple of Dagon*his head
The Royal Family: Prince Charles's title*prince of wales
Toys Games: Eva Gabor and Johnny Carson popularized this game by climbing over each other*twister
Toys Games: In this game players take turns placing disks on an 8x8 board*othello
Anzac troops come from which 2 countries*australia and new zealand
At the time of Julius Caesar, who was the ruler of Egypt*cleopatra
Chemically pure gold contains how many karats*twenty four
From what is velvet made*silk
How is Samuel Clemens better known*mark twain
How is the mathematically related structure of beads strung on parallel wires in a rectangular frame better known*abacus
How many bits are in a nibble*four
How many faces has an icosahedron*twenty
How many gold medals did Jesse Owens win in the 1936 Berlin Olympics*four
How many seconds are in a day*86400
In "Star Trek", what colour was Mr Spock's blood*green
In ancient Egypt which animal was considered sacred*cat
In the Gregorian calendar after 10,000 years by how many days will the calendar be wrong by*three
In the period 978-1016 England was ruled by which "Unready" king*ethelred
In which US city is the Sears tower*chicago
In which year was the Rosetta stone written*196 bc
Jonquil is a shade of which colour*yellow
Most people wear a watch on their ---------- wrist*left
Similes: As cute as a(n) ----------*button
Similes: As graceful as a(n) ----------*swan
Similes: As pale as a(n) ----------*ghost
The book "Wamyouruijoshou" was the first to use what word*kite
The term Septiquinquennial represents how many years*seventy five
The term Sesquincentennial represents how many years*one hundred and fifty
What are noctilucent, cirrus, and cirrostratus categories of*clouds
What city was founded in 753 BC*rome
What did Temujin change his name to*genghis khan
What does the ancient Greek word "electron" mean*amber
What is Harry Houdini famous for being*escapologist
What is the fastest growing species of grass*bamboo
What is the square root of -1*i
What is the state capital of Florida*tallahassee
What name is given to a settlement which is clustered around a central point*nucleated
What name is given to the point where a river starts*source
What religion was founded by Lao-tzu*taoism
What sport is governed by the rules drafted by the Marquis of Queensbury*boxing
What was Citizen Kane's first name*charles
What was the first transatlantic radio message sent*s
When was the Rosetta stone found*1799
Where are the great Walls of Babylon located in the modern day world*iraq
Which American state is known as the Lone Star State*texas
Which artificial fiber was invented in 1938*nylon
Which famous piece of artwork depcits the Battle of Hastings*bayeux tapestry
Which group of people elect the pope*cardinals
Which group was formed in 1972 by Don Henley and Glen Fry*the eagles
Which living bird has the longest wingspan*the albatross
Which part of a cat's eye reflects light*tapetum
Which people's republic stands between the Bay of Bengal and the foothills of the Himalayas, bordered by India and Burma*bangladesh
Which U.S City is the home of the Mowton Record Company*detroit
Which word is used to mean, malicious enjoyment at the misfortunes of others*schadenfreude
Who appears on the 10,000 dollar (US) note*salmon chase
Who does the statue, the Colossus of Rhodes, depict*helios
Who invented Tetris*alexi pazhitnov
Who invented the toothbrush*william addis
Who painted the ceiling of the Sistine Chapel*michelangelo
Who said 'The greater our knowledge increases, the more our ignorance unfolds'*john f. kennedy
Who was Ancient Egyptian fertility god*min
Who was the first man to reach the North Pole*robert edwin peary
Who was the first woman to fly the Atlantic alone*amelia earhart
Who wrote the Belgariad*leigh and david eddings
Words containing 'ten': A choice cut of meat*tenderloin
UK 50s: The 1951 Festival of Britain was centred on which city?*london
Vampires: Muppet vampire enjoyed doing this*counting
Chekhov Quotations: "Doctors can bury their mistakes, Architects can only advise their clients to plant vines."*Frank Lloyd Wright
Chemical compounds or mixtures that undergo rapid burning or decomposition with the generation of large amounts of gas and heat and the consequent production of sudden pressure effects*Explosives
Cheyenne, Navahoe and Arapaho are all what*native american tribes
Chief monetary unit of germany*deutschmark
Chub, gudgeon and perch are all types of what*freshwater fish
Closely related to pascal, niklaus wirth also played a part in what computer language's creation*modula
Clothes designer Alexander McQueen works for which fashion house*givenchy
Cocktails: whiskey, angostura bitters, and sugar make an*old fashion
Common name for a large sea turtle, named for the color of its fat, although the animal is brownish overall*green turtle
common name for a large sea turtle, named for the color of its fat, although the animal is brownish overall?*green turtle
Common ore of iron, and one of the most commonly occurring minerals in nature*Goethite
Company what is the abbreviation for lake minnetonka*lake tonka
Complete this saying 'All ship shape and'*bristol fashion
Condition in a circuit in which the combined impedances of the capacity and induction to alternating currents cancel each other out or reinforce each other?*Resonance
Confuscious Say: He who ------ in church, sits in own pew*farts
Confuscious Say: Man who run behind car get ----------*exhausted
Confuscious Say: Nail on board is not good as -------- on bench*screw
Conifer with dark foliage*cypress
Conventionally middle class materialist*bourgeois
Coolidge what heisman trophy winner returned his first nfl kickoff for a touchdown*tim brown
Countries of the world:northern part of central American Isthmus, major cities include Quezaltenango & Escuintla*guatemala
Countries of the world:western Asia, the capital is Tehran*iran
Countries of the world:western coast of South American, major cities include Arequipa & Trujillo*peru
Creation of programs, databases etc.for computer applications*authoring
Credit card on which magnetically encoded information is stored to be read by an electronic device*swipe card
Creek what position has been held by 266 men, 33 of whom have died violently*pope
Crusher Which word is derived from "user of hashish"*assassin
Dakar is the capital of ----------*senegal
Days of the week - whats the only day named for a planet*saturday
Death of body tissue usually caused by bad circulation*gangrene
Decorations what shadow team driver was killed testing, prior to the 1974 south african grand prix*peter revson
Delage guy delage claimed to be the first person to swim across which ocean*atlantic ocean
Denver is the capital of ----------*colorado
Did you know that if ---------- had not been shot, & not convicted for killing JFK, he would have been convicted for killing Officer Tippet*lee harvey oswald
Dirk, poniard, and stiletto are all types of what*daggers
Disease of animals, especially birds, monkeys, & humans, caused by infection by protozoans of the genus plasmodium & characterized by chills & intermittent fever*malaria
Disney: "Hair Facts" from the Broadway version of "Beauty and The Beast": Four characters' wigs are made of -------------: Mrs. Potts, Lumiere, Madame de la Grande Bouche, and the Sugar Bowl from "Be Our Guest."*yak hair
Disney: "Hair Facts" from the Broadway version of "Beauty and The Beast": The 30-inch length human hair needed to build Belle's wig was specially imported from ------------------*india
Disney: "Hair Facts" from the Broadway version of "Beauty and The Beast": The Beast's tail is made up of seven yards of -----------------*human hair
Do arteries carry blood towards or away from the heart*away
Does a cat groom itself more in cold weather or in warm weather*warm
Does Barry Manilow know you raid his wardrobe?*Breakfast Club
Does elizabeth ii face to the left or right on a british coin*right
Doraphobia is the fear of ----------*fur
Driving: what country is identified by the letter b*belgium
During supersonic flight, what temperature does the skin on the nose of a concorde reach*two hundred and sixty degrees fahrenheit
During which conflict did the battles of Alma and Inkermann take place*the crimean war
During World War II, which London theatre boasted, "We never closed"*the windmill
Dutch-born Swiss scientist, who discovered basic principles of fluid behavior*daniel bernoulli
Earth's outer layer of surface soil or crust is called the ----------*lithosphere
Edgar allan poe introduced mystery fiction's first fictional detective, auguste c. dupin, in what 1841 story*the murders in the rue morgue
Edgar Allan Poe introduced mystery fictions first fictional detective, Auguste C. Dupin, in what 1841 story*the murders in the rue morgue
Effect that occurs when two or more waves overlap or intersect*interference
Eglantine Jebb founded which charitable organisation*save the children fund
Either of the contibuters to the 4-note Hindol "Taril ha juj Girlja Shankur"*marathe
El cid was the name of what college's mascot goat*annapolis naval academy
Electrical circuit made by depositing conductive material on the surface of an insulating base*printed circuit board
Eve of All Saints Day*halloween
Every human first spent about half an hour as a single what*cell
Excessive discharge of blood from blood vessels, caused by pathological condition of the vessels or by traumatic rupture of one or more vessels*haemmorage
Excluding man, what is the longest lived land mammal*elephant
Famous Last Words: Let it down -------*slowly
Famous Last Words: You wouldn't hit a guy with -------- on, would you*glasses
Fandible, lateral line, & dorsal fin are parts of a(n) ----------*fish
Fear of dryness is called*xerophobia
Fife the only member of the band zz top without a beard has what last name*beard
Fill in: blind as a ----------*bat
Film: who played "sister agnes" in agnes of god*meg tilly
Film: who played an indian in tell them willie boy is here*katharine ross
Fired unglazed pottery*bisque
Fishy short story also known as Creation Took Eight Days*goldfish bowl
Floating wreckage at sea*flotsam
Fluid produced in the lacrimal glands above the outside corner of each eye*tears
Football the chicago ----------*bears
Football the new orleans ----------*saints
Football the San Diego ----------*chargers
Football the seattle ----------*seahawks
For how long is the note sustained at the end of the beatles' song 'a day in the life'*forty seconds
For how much did peter minuit buy manhattan island*24 dollars
For the development of a vaccine against which disease is Jonas Edward Salk best remembered*poliomyelitis (polio)
For what did the knights of the round table search*the holy grail
For what genre of book is isaac asimov famous*science fiction
For what is the Italian town of Carrara world famous*marble
For which ad campaign was the line 'i can't believe i ate the whole thing' used*alka seltzer
For which decoration do the letters C.G.M. stand*conspicuous gallantry medal
Form of visible electric discharge between rain clouds or between a rain cloud and the earth (Electricity)?*lightning
Former baseball star chuck connors hits a bull's-eye with adult-western*rifleman
Founded in 1896, what was IBM formerly called*tabulating machine company
Four European countries keep Greenwich Mean Time. The UK and Ireland are two, name either of the others*iceland
Fox what is the capital of idaho*boise
Francophobia is a fear of ----------*anything french
From what material is the ring made in Sumo Wrestling*clay
From what words is dublin derived*dubh linn
From which American state do the Bighorn Mountains arch northwest into southern Montana*wyoming
From which Marx Brothers film comes the line 'Either he's dead, or my watch has stopped*a day at the races
From which plant family do vanilla pods come*orchid
From which Shakespeare play does the line 'All the world's a stage' come*as you like it
Generals Gowon, Abasanjo and Abacha have all been leaders of which African State*nigeria
Generation X Toys: Once scarce, pudgy dolls that came with their own birth certificates*cabbage patch kids
Genus of annual and perennial herbs (Buttercup) containing about 20 species, grown for their showy flowers*Adonis
Geography: ----------- is smaller than the state of Montana (116,304 square miles and 147,138 square miles, respectively)*italy
Geography: --------------- got its start as a major tourist destination during the early days of World War II. German U-boats threats off the eastern United States compelled the wealthy to find new places to vacation. At one time one had to be a millionaire to enjoy ----------, but that hasn't been the case for years*acapulco
Geography: Coral reefs, sometimes called the "rain forests of the sea," cover more than 6,500 square miles in the -------------, the Gulf of Mexico, off Florida, and the Pacific. They are home to an estimated 550 species of fish, and are major tourist attractions*caribbean
Geography: For centuries, Spain's ----------------- has been and still is one of the world's largest*fishing fleet
Geography: In ----------------, the Presidential highway links the towns of Gore and Clinton*new zealand
Geography: One prominent feature of the Greek climate is its ample sunshine. The sun shines in Greece about 3,000 hours per year. Greece's heavily indented shores give the country extraordinary beauty, quite unique in the --------------. The length of the Greek coastline is estimated at 15,000 kilometers.Mediterranean
Geography: Ruby Falls, America's highest underground waterfall open to the public, is located on historic Lookout Mountain in Chattanooga, ---------------*tennessee
Geography: The -------- river has frozen over at least twice, in 829 and 1010 A.D*nile
Geography: The ------------- got its name from the occasionally extensive blooms of algae that, upon dying, turn the sea's normally intense blue-green waters to red*red sea
Geography: The --------------- comprises an area as large as Europe. Its total land mass is some 3,565,565 square miles*sahara desert
Geography: The city of Los Angeles is more than one-third the size of the entire state of --------------*rhode island
Geography: The Federated States of -----------, located at the Eastern Caroline Islands in the northwest Pacific Ocean, has more than 600 islands and 40 volcanos*micronesia
Geography: The highest mountain in the British Isles, Ben Nevis in western ----------, is just 4,406 feet high. In many other countries, a "mountain" of this size would be considered something less than a large hill*scotland
Geography: The longest main street in America, 33 miles in length, can be found in Island Park, ----------*idaho
Geography: Twenty-three states in the U.S. border an ------------*ocean
Geography: Under a treaty dating back to 1918, if the Grimaldis of --------- should ever be without a male heir, --------- would cease to exist as a sovereign state and would become a self-governing French protectorate*monaco
Geography: Yuma, ------------ has the most sun of any locale in the U.S. - it averages sunny skies 332 days a year*arizona
George Stephenson was born in what year*1781
Guiyaquil is the largest city in which country*ecuador
Haggard as what is merle haggard also known as*okie from muskogee
Half years who made her show business debut at the age of 2 1/2 as part of her family's vaudeville act on the 'new grand theater stage'*judy garland
Harrison What do the San Joaquin kit fox, Hawaiian hawk and Ocelot have in common*endangered species
Haven What was the name of Flash Gordon's girlfriend*dale arden
he said 'i have nothing to offer but blood, tears, toil and sweat'?*winston churchill
Heavier-than-air craft that derives its lift not from fixed wings like those of conventional airplanes, but from a power-driven rotor or rotors, revolving on a vertical axis above the fuselage?*helicopter
Hippophobia is a fear of ----------*horses
Hockey the vancouver ----------*canucks
Hoffman who wrote about a british agent named george smiley*john le carr
Hours how many times do your ribs move every year during breathing*five million
Household items such as television sets and audio equipment are know as*brown goods
How did Mark Chapman shock the world*shot john lennon
how did Rose Nilin's husband Charlie die on The Golden Girls?*He died of a heart Attack while making love to Rose.
How did the crew of Red Dwarf get brought back to life?*By Nanobots
how does the mermaid buy the gift for tom hanks in "splash"?*her necklace
How far does the cruise liner 'queen elizabeth ii' move for each gallon of diesel it burns*six inches
How is abba calling for help*sos
How long does it take a fully loaded supertanker to stop from travelling at normal speed*twenty minutes
How long was Jonah in the whale's stomach*3 days
How many blades are there on a kayak paddle*two
How many bonus points in Scrabble if all seven tiles played at once*fifty
How many cards are there in each suit of a standard deck*thirteen
How many children did adam and eve have*three
How many children did president william henry harrison have*ten
How many cigars did Sir Winston Churchill ration himself to a day*fifteen
How many consecutive years was the ed sullivan show on tv*twenty three
How many days were the american hostages held in Iran*four hundred & forty four
How many feet are there in one fathom*six
How many folds does a monopoly board have*one
How many letters are used for roman numerals*seven
How many member states are there in the United Arab Emirates*seven
How many miles are there in a league*three
How many people did andrew cunanan kill before killing gianni versace*four
How many players are there in a men's lacrosse team*ten
How many sheets of paper are there in a ream*five hundred
How many sides does a dodecagon have*twelve
How many stitches are on a regulation baseball*one hundred and eight
How many teeth does a walrus have*eighteen
How many VCs were awarded in the Falklands War*two
How many years elapsed between the creation of the Republic of Vietnam and Saigon falling to the communists*thirty
How old are oak trees before they produce acorns*fifty
How old was leann rhimes when she became a country music star*fourteen
Hudson how many points are awarded to the winning driver of a formula 1 grand prix race*ten
Hukusai and Hiroshige were famous Japanese what*artists
Hydrophobophobia is the fear of*rabies
I1948 Olivia ---------- -John (in Cambridge, England), singer, born*newton
If a chemical is 'anhydrous' what does it not contain*water
If a dish is served A la Chantilly, what would be its main ingredient*whipped cream
If you "peg out" what game are you playing?*cribbage
If you were born on 01 October what star sign (Zodiac) would you be*libra
If you were born on 02 September what star sign (Zodiac) would you be*virgo
If you were born on 04 December what star sign (Zodiac) would you be*sagittarius
If you were born on 04 July what star sign (Zodiac) would you be*cancer
If you were born on 05 August what star sign (Zodiac) would you be*leo
If you were born on 05 June what star sign (Zodiac) would you be*gemini
If you were born on 07 April what star sign (Zodiac) would you be*aries
If you were born on 07 May what star sign (Zodiac) would you be*taurus
If you were born on 08 December what star sign (Zodiac) would you be*sagittarius
If you were born on 09 April what star sign (Zodiac) would you be*aries
If you were born on 10 October what star sign (Zodiac) would you be*libra
If you were born on 12 September what star sign (Zodiac) would you be*virgo
If you were born on 13 October what star sign (Zodiac) would you be*libra
If you were born on 14 May what star sign (Zodiac) would you be*taurus
If you were born on 14 November what star sign (Zodiac) would you be*scorpio
If you were born on 15 January what star sign (Zodiac) would you be*capricorn
If you were born on 17 July what star sign (Zodiac) would you be*cancer
If you were born on 18 August what star sign (Zodiac) would you be*leo
If you were born on 18 October what star sign (Zodiac) would you be*libra
If you were born on 21 November what star sign (Zodiac) would you be*scorpio
If you were born on 22 February what star sign (Zodiac) would you be*pisces
If you were born on 23 June what star sign (Zodiac) would you be*cancer
If you were born on 24 August what star sign (Zodiac) would you be*virgo
If you were born on 24 February what star sign (Zodiac) would you be*pisces
If you were born on 24 March what star sign (Zodiac) would you be*aries
If you were born on 25 October what star sign (Zodiac) would you be*scorpio
If you were born on 26 December what star sign (Zodiac) would you be*capricorn
If you were born on 26 March what star sign (Zodiac) would you be*aries
If you were born on 26 November what star sign (Zodiac) would you be*sagittarius
If you were born on 26 October what star sign (Zodiac) would you be*scorpio
If you were born on 27 May what star sign (Zodiac) would you be*gemini
If you were born on 28 February what star sign (Zodiac) would you be*pisces
If you were born on 28 June what star sign (Zodiac) would you be*cancer
If you were born on 29 April what star sign (Zodiac) would you be*taurus
If you were born on 30 June what star sign (Zodiac) would you be*cancer
If you were born on 31 May what star sign (Zodiac) would you be*gemini
Ilex is the botanical name of which shrub*holly
Implant*breast implants
In "peanuts", what is the surname of lucy and linus*van pelt
In '64, whom did J Edgar Hoover call America's "most notorious liar"*martin luther king jr
In 'dawson's creek', who does michelle williams play*jennifer lindley
In 'star wars', who was darth vader's face*sebastian shaw
In 'startrek', who did william shatner play*captain james t kirk
In 'the wizard of oz', what was dorothy's dog's name*toto
In -322 BC Aristotle dies of ----------*indigestion
In 1000 Leif ---------- discovers "Vinland" (possibly America)*ericson
In 1066 Battle of Hastings, in which William the ---------- wins England*conqueror
In 1066 William the ---------- crowned William I of England*conqueror
In 1087 William I The Conqueror, King of England and Duke of---------- , dies*normandy
In 1189 England's King ---------- (the Lion-Hearted) crowned in Westminster*richard i
In 1271 ---------- king of Bohemia and Poland (1278-1305), born*wenceslas ii
In 1290 Bilbo ---------- (in Shire Reconning), born*baggins
In 1292 Saidi, great ---------- poet (Orchard, Rose Garden) dies*persian
In 1364 Battle of Auray, ---------- forces defeat French at Brittany*english
In 1513 ---------- Nuez de Balboa is the first European to see the Pacific Ocean*vasco
In 1520 King Henry VIII of England orders bowling lanes to be built at---------- , in London*whitehall
In 1520 Martin ---------- publicly burned papal edict demanding that he recant*luther
In 1540 Society of Jesus (Jesuits) founded by Ignatius----------*loyola
In 1541---------- , Chile founded*santiago
in 1543, who published a theory that planets revolve around the sun?*copernicus
In 1583 ---------- Alighieri Day*dante
In 1620 The Mayflower sets sail from ---------- with 102 Pilgrims*plymouth
In 1632 Sir ---------- Wren, England, astronomer/great architect, born*christopher
In 1666 Great London ---------- begins in Pudding Lane. 80% of London is destroyed*fire
In 1672 (Italy) Giovanni Cassini discovers Rhea, a satellite of----------*saturn
In 1686 1st volume of ---------- Newton's "Principia" published*isaac
In 1708 ---------- von Haller, the father of experimental physiology*albrecht
In 1709 Elizabeth, empress of ---------- (to Peter the Great and Catherine I),born*russia
In 1713 Ferdinand VI, king of ---------- (1746-59), born*spain
In 1725 ---------- -Joseph Cugnot, designed and built first automobile, born*nicolas
In 1727 Severe earthquake in----------*new england
In 1731 Henry---------- , English physicist, chemist born*cavendis
In 1732 George---------- , father figure for U.S., President (1789-1796), born*washington
In 1741 Vitus Bering, Dutch ---------- and explorer, died*navigator
In 1752 Nicolas---------- , inventor of food canning, bouillon tablet, born*appert
In 1758 Horatio Nelson Burnham Thorpe, Britain, naval hero at---------- , born*trafalgar
In 1774 John Chapman, alias Johnny---------- , born*appleseed
In 1776 Continental Congress renames "---------- ", "United States"*united colonies
In 1782 ---------- Paganini, Genoa, Italy, composer/violin virtuoso (Princess Lucca), born*niccolo
In 1783 Washington ----------, writer (Rip Van Winkle, Legend of Sleepy Hollow), born*irving
In 1784 Empress of ---------- sets sail on first New York to China route*china
In 1795 Third partition of Poland, between Austria, ---------- and Russia*prussia
In 1798 ---------- 1st emperor of Brazil (1822-31), king of Portugal, born*pedro i
In 1812 Fire of----------*moscow
In 1812 Waltz introduced into English---------- . Most observers consider it disgusting & immoral. No wonder it caught on!*ballrooms
In 1813 German Kingdom of ---------- abolished*westphalia
In 1815 World's first commercial ---------- factory is established in Switzerland*cheese
In 1817 First American school for the ---------- (Hartford, Connecticut)*deaf
In 1820 Susan B.---------- , Woman's suffaregette, born*anthony
In 1821---------- , El Salvador, Guatemala, Honduras and Nicaragua gain independence*costa rica
In 1823 Charles ---------- of Scotland begins selling raincoats (Macs)*macintosh
In 1824 ---------- defies Pele (Hawaiian volcano goddess) and lives*kapiolani
In 1824 Kapiolani defies ---------- (Hawaiian volcano goddess) and lives*pele
In 1825 Hannah Lord ---------- of New York grabs her scissors and creates the first detachable collar on one of her husband's shirts, in order to reduce her laundry load*montague
In 1830 Eadweard ---------- , pioneered study of motion in photography, born*muybridge
In 1833 Ernesto ---------- Moneta, Italian journalist (Nobel Peace Prize 1907) BORN*teodoro
In 1833 Ernesto Teodoro---------- , Italian journalist (Nobel Peace Prize 1907) BORN*moneta
In 1836 Darwin returns to England aboard the HMS----------*beagle
In 1844 Henry ---------- Heinz, founded a prepared-foods company, born*john
In 1848 ---------- (U.S.) and Lassell (England) independently discover Hyperion*bond
In 1848 Mexico sells U.S. Texas, ---------- , New Mexico and Arizona*california
In 1849 ---------- Pavlov, Russia, physiologist/pioneer in psychology, born*ivan
In 1849 Edgar Allen Poe dies in Baltimore at----------*forty
In 1849 Ivan---------- , Russia, physiologist/pioneer in psychology, born*pavlov
In 1852 2nd French empire established; Louis ---------- becomes emperor*napoleon
In 1853 First round-the-world trip by yacht (Cornelius---------- )*vanderbilt
In 1854 Frederick---------- , Arms manufacturer, born*krupp
In 1857 Konstantin---------- , pioneer in rocket and space research, born*tsiolkovsky
In 1858 First electric ---------- is installed in Boston, Mass*burglar alarm
In 1861 C.S.A. President Jefferson ---------- is inaugurated at Montgomery, AL*davis
In 1863 International Committee of the ---------- is founded (Nobel 1917, 1944, 1963)*red cross
In 1865 Charles Proteus ---------- , electronics pioneer, born*steinmetz
In 1865 President Abraham ---------- shot in Ford's Theatre by J.W. Booth*lincoln
In 1866 H(erbert) G(eorge) Wells---------- , England (War of the Worlds), born*bromley
In 1869 Black Friday -- Wall Street panics after Gould and ---------- attempt to corner gold*fisk
In 1869 First postcards are issued in----------*vienna
In 1870 ---------- is founded in New York City*ywca
In 1870 Napoleon ---------- captured at Sedan*iii
In 1872 Darius---------- , composer born*milhaud
In 1872 Emily ---------- authority on social behavior, writer (Etiquette), born*post
In 1873 Ejnar Hertzsprung, ---------- astronomer (Hertzsprung-Russell diagram), born*danish
In 1874 Gertrude---------- , writer, born*stein
In 1875 Violent bread riots at----------*montreal
In 1878 First telephone exchange in ---------- opens with 18 phones*san francisco
In 1879 The first "mobile home" (horse drawn) is used for a journey between London and ----------*cyprus
In 1879 Thomas Edison commercially perfects the----------*light bulb
In 188 ---------- Roman emperor (211-17), born*caracalla
In 1881 William Edward Boeing, founded ---------- company*aircraft
In 1882 James---------- , writer, born*joyce
In 1883 Kahlil ---------- , philosopher, born*gibran
In 1887 Sino-Portuguese treaty recognizes Portugal's control of----------*macao
In 1888 George Eastman patents first rollfilm camera and registers----------*kodak
In 1890 ---------- Island (NYC) opens as a US immigration depot*ellis
In 1890 Edwin---------- , radio pioneer (invented FM) ,born*armstrong
In 1892 Donald Wills ---------- , founded an aircraft company*douglas
In 1892, who raised the marriageable age for girls to 12 years old*italy
In 1894 Japan defeats China in Battle of----------*ping yang
In 1898 ---------- -American War begins*spanish
In 1898 Enzo---------- , car designer and manufacturer, born*ferrari
In 1898 Spanish-American War ends -- U.S. acquires Guam, Puerto Rico, the Phillipines, and ---------- from Spain*cuba
In 1899 The first auto repair shop opens in---------- , MA*boston
In 1900 British annex ---------- (South Africa)*natal
In 1901 Enrico---------- , Italy, nuclear physicist, born*fermi
In 1901 First ---------- Peace Prizes (to Jean Henri Dunant, Frederic Passy)*nobel
In 1901 Pres William ---------- assassinated by Leon Czologosz in Buffalo, New York*mckinley
In 1904 ---------- Horowitz, pianist born*vladimir
In 1904 Federation Internationale de Football Association (---------- ), Soccer's World governing body forms*fifa
In 1904 Sir John ---------- , actor, singer, born*gielgud
In 1905 Felix---------- , U.S. physicist (Nobel 1952), born*bloch
In 1906 James A ---------- circus showman (Barnum & Bailey), dies at 58*bailey
In 1906 Karl ---------- demonstrates the first 'permanent wave' for hair, in London*nessler
In 1907 (USA) For the 1st time a ball drops at ---------- Square to signal the new year*times
In 1908 ---------- unites with Greece*crete
In 1908 Buddy---------- , actor (Beverly Hillbillies, Barnaby Jones), born*ebsen
In 1909 Alberto Romero "Cubby" ---------- film producer, born*broccoli
In 1909 Comte de Lambert of France sets airplane altitude record of ---------- m*three hundred
In 1909 Victor Borge, pianist, ---------- , born*comedian
In 1910 Fritz---------- , writer, born*leiber
In 1911 (US) Gugliemo ---------- sends the first wireless message across the Atlantic*marconi
In 1911 ---------- Burchett, Australian Communist, journalist, writer, born*wilfred
In 1912 Chuck Jones animator (---------- , Daffy Duck), born*bugs bunny
In 1912 RMS ---------- sets sail for its first and last voyage*titanic
In 1912 Yuan ---------- elected the first President of the Republic of China*shik-k'ai
In 1913 ---------- and Atlantic mix as engineers blow Gamboa Dam, opening the Panama Canal*pacific
In 1913 Oleg ---------- Paris France, fashion designer for Jackie Kennedy, born*cassini
In 1914 Cardinal ---------- della Chiesa becomes Pope Benedict XV*giacome
In 1914 Gypsy Rose ---------- (in Seattle, WA), stripper, born*lee
In 1914 St Petersburg, Russia changes name to----------*petrograd
In 1915 ---------- Miller, playwright (Death of a Salesman, The Crucible), born*arthur
In 1915 Lorne---------- , actor (Bonanza, Battlestar Galactica), born*greene
In 1916 First professional ---------- tournament held*golf
In 1917 ---------- the Cat, cartoon character, born*felix
In 1917 Lenin returns to Russia to start ---------- Revolution*bolshevik
In 1917 Mata Hari executed by firing squad outside of----------*paris
In 1918 ---------- President Sidonio Paes is assassinated*portugese
In 1919 ---------- Hayworth (in New York), actor, alzheimer victim, born*rita
In 1919 Art---------- , jazz drummer (Jazz Messengers), born*blakey
In 1919 Volstead Act passed by U.S. Congress, starting----------*prohibition
In 1920 ---------- Warden, actor (Verdict, Brian's Song), born*jack
In 1920 Jack---------- , actor (Verdict, Brian's Song), born*warden
In 1920 Japan receives League of Nations mandate over ---------- islands*pacific
In 1920 Mickey---------- , actor (too many credits to mention), born*rooney
In 1921 ---------- Poston, comedian, actor (Newhart), born*tom
In 1921 Robert---------- , archbishop of Canterbury born*runcie
In 1922 Ava---------- , actress, born*gardner
In 1922 Yvonne---------- , Vancouver BC, actress (Lily Munster in the Munsters), born*de carlo
In 1924 ---------- Bacall (in Staten Island, NY), actor, whistler (Dark Passage, Key Largo, Always), born*lauren
In 1924 ---------- Kollontai of Russia becomes 1st woman ambassador*alexandra
In 1924 Albania becomes a----------*republic
In 1926 ---------- Tunney defeats Jack Dempsey for world heavyweight boxing title*gene
In 1926 Chuck Berry, St Louis, USA, rocker (---------- ), born*roll over beethoven
In 1926 Miles ---------- trumpeter; pioneered cool jazz (Porgy & Bess), born*davis
In 1926 Roger Moore, actor, (---------- , numerous James Bond movies), born*the saint
In 1927 ---------- Grass, German novelist, poet (The Tin Drum) born*gunter
In 1927 Al---------- , singer, born*martino
In 1927 Harvey---------- , actor, born*korman
In 1928 ---------- Mouse makes his screen debut in "Steamboat Willie."*mickey
In 1928 Chiang ---------- becomes president of China*kai-shek
In 1928 George Peppard, actor (---------- , Blue Max, A-Team) born*breakfast at tiffany's
In 1928 Katharine Hepburn makes her New York stage debut in "---------- ."*night hostess
In 1928 Mae West makes her New York City debut in a daring new play, " ---------- "*diamond lil
In 1928 Spanky McFarland, actor, little rascal born
In 1930 A cow is flown (and milked in flight) for first time. Her milk was sealed in paper containers and dropped by ---------- over St. Louis, MO. I knew you'd want to know*parachute
In 1930 Edward ---------- England, actor (Breaker Morant, Equalizer), born*woodward
In 1930 Harold---------- , playwright, born*pinter
In 1930 Synthetic ---------- first produced*rubber
In 1931 ---------- Bancroft AKA Mrs Mel Brooks, Bronx, actress (Graduate), born*anne
In 1931 Al ---------- convicted of tax evasion, sentenced to 11 years in prison*capone
In 1932 ---------- vaccine for humans announced*yellow fever
In 1933 ---------- (in Tokyo, Japan), singer, wife of John Lennon, born*yoko ono
In 1933 Yevgeny V ---------- USSR, cosmonaut (Soyuz 5), born*khrunov
In 1934 1st ---------- rpm recording released (Beethoven's 5th)*33 1/3
In 1934 First " ---------- " (laundromat) is opened, in Fort Worth, Texas*washateria
In 1934 Shirley ---------- appears in her 1st movie, "Stand Up & Cheer"*temple
In 1935 ---------- and Harriet Nelson married*ozzie
In 1935 George ---------- 's "Porgy and Bess" opened in New York City*gershwin
In 1935 George Gershwin's "---------- " opened in New York City*porgy and bess
In 1935 Luciano Pavarotti Moderna Italy, operatic ---------- (Oh Giorgio), born*tenor
In 1936 ---------- Holly singer (Peggy Sue, That'll Be the Day), born*buddy
In 1936 ---------- Oerter, US discus thrower, born*al
In 1936 Buddy ---------- singer (Peggy Sue, That'll Be the Day), born*holly
In 1936 Pumping begins to build---------- , San Francisco*treasure island
In 1938 ---------- Koenig Chicago Ill, actor (Chekov-Star Trek), born*walter
In 1938 ---------- Zeppelin II, world's largest airship, makes maiden flight*graf
In 1938 Christopher Lloyd, actor (Taxi, ---------- , Back to the Future, Addams Family), born*star trek iii
In 1938 Christopher Lloyd, actor (Taxi, Star Trek III, Back to the Future,---------- ), born*addams family
In 1939 ---------- Airport opened in New York City*laguardia
In 1939 ---------- go on sale in the U.S. for the first time*nylon stockings
In 1939 ---------- Tabei Japan, 1st woman to climb Mount Everest, born*junko
In 1939 Britain declares war on Germany. France follows 6 hours later quickly joined by Australia, New Zealand, South Africa and----------*canada
In 1939 Franklin D. Rooseveldt declares "limited national emergency" due to war in----------*europe
In 1939 Jim Bakker, ---------- , con-man, born*televangelist
In 1939 Soviet-German treaty agree on 4th partition of ---------- (WW II) and gives Lithuania to the USSR*poland
In 1940 (England) F. Scott---------- , author, died*fitzgerald
In 1940 ---------- Avalon, singer (Four Seasons), born*frankie
In 1940 ---------- Pele, soccer player extraordinaire, born*edison
In 1941 Jaqueline Bisset (in England), actor (---------- ), born*deep
In 1942 ---------- Funichello (in Utica, NY), mouseketeer, actor, born*annette
In 1942 1st controlled ---------- chain reaction (University of Chicago)*nuclear
In 1942 Japanese occupied*manila
In 1942 Michael---------- , author (Andromeda Strain, Jurrasic Park, Rising Sun), born*crichton
In 1942 Paul---------- , singer (Kodachrome, Graceland), born*simon
In 1943 ---------- North, arms dealer, born*oliver
In 1943 Canadian Army troops arrive in*north africa
In 1943 FDR appoints Gen Eisenhower supreme commander of ---------- forces*allied
In 1944 Patty---------- , singer born*labelle
In 1945 ---------- National Day*hungarian
In 1945 ---------- Peron becomes dictator of Argentina*juan
In 1945 Benito ---------- Fascist leader & mistress captured, tried, & shot*mussolini
In 1945 Japanese forces in the ---------- surrender to Allies*philippines
In 1945 Juan ---------- becomes dictator of Argentina*peron
In 1945 Nazi concentration camp at ---------- liberated by US 80th Division*buchenwald
In 1945 Nazi Himmler committed suicide while in prison at---------- , Germany*luneburg
In 1945 President ---------- announced atomic bomb secret shared with Britain and Canada*truman
In 1946 ---------- J. Jeans, astrophysicist, dies on his 69th birthday*james
In 1946 Hayley ---------- (in London, England), actor, born*mills
In 1946 James J.---------- , astrophysicist, dies on his 69th birthday*jeans
In 1946 Oliver Stone NYC, director (---------- , Good Morning Vietnam, Platoon), born*wall st
In 1946 Richard---------- , musician (Carpenters), born*carpenter
In 1946 Ten Nazi leaders hanged as war criminals after ---------- trials*nuremberg
In 1947 ---------- Smith, actress (Charlie's Angel, Nightkill), born*jaclyn
In 1947 Dan---------- , U.S. Vice-president (1989-1992), alleged twit, born*quayle
In 1947 First instant develop ---------- demonstrated in NY City by E. H. Land*camera
In 1947 Ted---------- , SD Calif, actor (Sam Malone-Cheers, 3 Men & a Baby), born*danson
In 1948 ---------- is established*world health organization
In 1948 ---------- Kidder (in Yellowknife), actor (Superman), born*margot
In 1948 ---------- National Day*burmese
In 1948 Phil ---------- comedian (SNL, Newsradio, Simpsons (Voice of Troy McLure)), born*hartman
In 1949 George Wendt, actor (---------- ), born*cheers
In 1949 West begins ---------- Airlift to get supplies around Soviet blockade*berlin
In 1950 David ---------- , Shirley Jones' kid on TV and real life, born*cassidy
In 1950 Morgan ---------- (in Dallas, TX), actress, born*fairchild
In 1951 ---------- Cougar Mellencamp, singer, born*john
In 1951 ---------- Keaton, actor (Pacific Heights, Batman, Multiplicity), born*michael
In 1951 Jay ---------- patents computer core memory*forrester
In 1951 John ---------- Mellencamp, singer, born*cougar
In 1951 Pam Dawber Detroit, actress (Mindy----------- ), born*mork and mindy
In 1951 Sir ---------- Geldof pop musician (Boomtown Rats, Band Aid), born*bob
In 1952 ---------- Connors, tennis player born*jimmy
In 1952 ---------- Stewart, rocker (Eurythmics-Here Comes the Rain Again), born*dave
In 1954 Dennis ---------- , actor, born*quaid
In 1954 Elvis Presley records his debut single, " ---------- "*that's all right
In 1955 ---------- Dean, actor, died in a car crash (born Feb 08, 1931)*james
In 1955 President Jose Antonio Remon of ---------- assassinated*panama
In 1956 Carrie ---------- (in Beverly Hills), actor (Star Wars, Blues Brothers), born*fisher
In 1956 Elvis ---------- appears on national TV for 1st time (Ed Sullivan)*presley
In 1956 James---------- . actor, born*ingram
In 1957 ---------- Fahey rocker (Bananarama), born*siobhan
In 1957 ---------- king of Norway, dies, Olaf succeeds him*haakon vii
In 1957 Ford Motor Co. introduced the---------- ! (Oh boy !)*edsel
In 1957 Seve ---------- , golfer, born*ballesteros
In 1957 Siobhan Fahey rocker (---------- ), born*bananarama
In 1958 Central African Rep made autonomous member of ---------- Commonwealth (Nat'l Day)*french
In 1959 Guggenheim Museum, designed by Frank Lloyd---------- , opens in New York*wright
In 1959 Princess Sarah 'Fergie'---------- , the Duchess of York, born*ferguson
In 1960 ---------- (REM Lead Singer), born*michael stipe
In 1960 Elvis Presley appears on a Frank ---------- TV special*sinatra
In 1960 Mali (without---------- ) gains independence from France (National Day)*senegal
In 1961 UK grants ---------- independence*sierra leone
In 1962 E. E. Cummings poet, dies at----------*sixty seven
In 1962 TV comedy "---------- " premiered on CBS*the beverly hillbillies
In 1963 "Beatlemania" is coined after the Beatles appear at the----------*palladium
In 1963 ---------- 1st tour (opening act for Bo Diddley and Everly Bros)*rolling stones
In 1963 Treaty banning atmospheric nuclear tests signed by US, ---------- , USSR*uk
In 1964 ---------- and Brezhnev replace Soviet premier Nikita Krushchev*kosygin
In 1964 Launch of Voskhod 1, 1st 3 man crew (---------- , Feokistov, Yegorov)*komarov
In 1964 Shooting begins on "The Cage" the pilot for Star----------*trek
In 1964, who recorded "baby love"*supremes
In 1965 "---------- " premiers*get smart
In 1965 ---------- National Day*gambian
In 1965 Bangladesh windstorm kills ----------*17,000
In 1965 Beatles release "---------- ."*yesterday
In 1965 Charlie Sheen, actor (---------- , Platoon), born*wall st
In 1966 "---------- " premiers on NBC TV*star trek
In 1966 Bechuanaland gains independence from England, becomes----------*botswana
In 1966 Botswana gains independence from ---------- (National Day)*Britain
In 1966 Emperor Haile ---------- (Ethiopia) visits Kingston Jamaica*selassie
In 1966 Lesotho (Basutoland) gains independence from ---------- (National Day)*britain
In 1966, which woman became the first Briton to fly solo around the world*sheila scott
In 1967 ---------- makes fly-by of Venus*mariner 5
In 1967 1st successful test flight of a----------*saturn v
In 1967 BBC bans Beatle's "---------- " (drug references)*a day in the life
In 1967 Che Guevara executed in----------*bolivia
In 1967 Gibraltar votes 12,138 to ---------- to remain British*forty four
In 1968 ---------- "Hey Jude", single goes #1 and stays #1 for 9 weeks*beatles'
In 1968 ---------- Lake actress (Hairspray, Ricki Lake Show), born*ricki
In 1968 Borman, ---------- and Anders first men to orbit moon*lovell
In 1968 Swaziland gains independence from ---------- (National Day)*britain
In 1968, who released 'carnival of life' and 'recital'*lee michaels
In 1969 Beatles release "---------- " album*abbey road
In 1969 Dr. Denton ---------- implants first temporary artificial heart*cooley
In 1969 Levi Eshkol dies, ---------- becomes premier of Israel*golda meir
In 1969 Libyan revolution, Col ---------- Gadhafi deposes King Idris*moammar
In 1970 ---------- Republic (Cambodia) declares independence*khmer
In 1970 Anwar Sadat elected president of Egypt, succeeding Gamal ---------- Nasser*abdel
In 1970 Beatles' " ---------- ," single goes #1 & stays #1 for 2 weeks*let it be
In 1970 George Harrison releases "---------- " single*my sweet lord
In 1970 Janis ---------- dies at age 27*joplin
In 1972 11 ---------- athletes are slain at Munich Olympics*israeli
In 1972 Alyassa---------- , actor (Who's the Boss) ,born*milano
In 1972 Harlow ---------- discoverer of the Sun's position in the galaxy, dies*shapley
In 1972 John Young & Charles ---------- explores Moon (Apollo 16)*duke
In 1973 2 Skylab 3 astronauts walk in space for a record ---------- hours*seven
In 1973 Billy Jean King beats Bobby ---------- in battle-of-sexes tennis match*riggs
In 1973 Elvis and ---------- Presley divorce after 6 years*priscilla
In 1973 The ---------- - Israel's missile boat - is unveiled*reshef
In 1974 ---------- TV host (Ed Sullivan Show), dies at 73*ed sullivan
In 1974 French president Georges ---------- died in Paris*pompidou
In 1974 Soyuz ---------- is launched*fourteen
In 1975 Israel formally signs Sinai accord with----------*egypt
In 1976 John Hathaway completes a bicycle tour of every continent in the world and cycling ---------- miles*50,600
In 1977 Cheryl ---------- replaces Farrah Fawcett on "Charlie's Angels"*ladd
In 1977 US recalls William---------- , ambassador to South Africa*bowdler
In 1978 ---------- Ali beats WBA heavyweight champion Leon Spinks*muhammad
In 1979 ---------- Chung-hee South Korean President is assassinated*park
In 1979 Mother Teresa of ---------- was awarded the Nobel Peace Prize*india
In 1980 ---------- people die when a pair of earthquakes struck NW Algeria*4,500
In 1980 BSD ---------- released*unix 3
In 1980 USA beats ---------- and wins the Olympic Gold Medal (4-2)*finland
In 1980, who recorded "Another One Bites the Dust"*queen
In 1981 "Late Night with David ---------- " premiers*letterman
In 1982 ---------- leaves Lebanon*palestinian liberation organization
In 1982 ---------- Portugal, a Spanish priest with a bayonet is stopped prior to his attempt to attack Pope John Paul II*fatima
In 1982 1st permanent artificial ---------- successfully implanted (U of Utah) in retired dentist Barney Clark; lived 112 days with the Jarvic-7 heart*heart
In 1982 Mt ---------- Observatory first to detect Halley's comet on 13th return*palomar
In 1982 Soyuz T-5 returns to---------- , 211 days after take-off*earth
In 1983 St Christopher----------- gains independence from Britain (Nat'l Day)*nevis
In 1984 Christopher ---------- , FBI's 'most wanted man' accidentally killed self*wilder
In 1984, who sang 'girls just want to have fun'*cyndi lauper
In 1985 Walt Disney World's ---------- -millonth guest*two hundred
In 1986 Andrei Tarkovski, Russian ---------- (Stalker), dies at 54*director
In 1986 Record 23,000 start in a marathon (---------- )*mexico city
In 1986 USSR frees dissident Andrei ---------- from internal exile*sakharov
In 1986 USSR releases US journalist ---------- Daniloff confined on spy charges*nicholas
In 1986, what was the maximum fuel capacity imposed in formula 1 racing*one
In 1987 ---------- Greene actor (Bonanza, Battlestar Galactica), dies at 72*lorne
In 1988 "Naked Gun" premieres, a movie based on TV's "---------- Squad"*police
In 1988 Lillehammer, ---------- upsets Anchorage to host 1994 Winter olympics*norway
In 1988 US-Soviet effort free 2 grey whales from frozen----------*arctic
In 1989 ---------- Chapman, member of the Monty Python team, dies from cancer*graham
In 1989 East Germans begin their flight to the west (via Hungary and---------- )*czech
In 1989 San Francisco is hit by an earthquake (Richter 6.9) at 5:05 p.m. Over 1/2 mile of the upper deck of the Nimitz freeway collapses crushing hundreds of cars. When it was over, 62 people had died and billion in damage had occured
In 1990 ---------- threatens to hit Israel with a new missile*saddam
In 1990 Iraqi Pres Saddam ---------- urges Arabs to rise against the West*hussein
In 1990 Lithuania, Estonia and ---------- hold their 1st joint session*latvia
In 1990 Rocky ---------- boxer, dies at 71, of heart failure*graziano
In 1991 ---------- Montand actor (Lets Make Love, Z), dies at 70*yves
In 1991 Miles ---------- jazz musician, dies at 65 from pneumonia*davis
In 1991 UN Security Council issues formal cease fire with ---------- declaration*iraq
In 1995 Barings Bank disaster. Nick ---------- loses billions of Pounds Sterling in offshore investments, ruining Barings Bank*leeson
In 1995 OJ Simpson acquitted for double murder of his Ex-wife ---------- and Ronald Goldman*nicole brown simpson
In 2161---------- : 8 of 9 planets aligned on same side of sun*syzygy
In 254 St ---------- begins his reign as Catholic Pope*stephen i
In 295 8th recorded ---------- passage of Halley's Comet*perihelion
In 31 BC Battle of Actium; ---------- defeats Mark Antony and becomes Emperor Augustus*octavian
In 43BC The Roman politician, ---------- , is slain*cicero
In 490 B.C. Athenians defeat second Persian invasion of Greece at----------*marathon
In 526 Earthquake kills ---------- in Antioch, Syria*250,000
In 680 ---------- ibn 'Ali, Shi'i religious leader, enters martyrdom*husain
In 70 BC ---------- (Publius Vergilius Maro) (Mantua, Italy), poet (Aeneid), born*virgil
In 742---------- , emperor (Holy Roman Empire), born*charlemagne
In 760 14th recorded ---------- passage of Halley's Comet*perihelion
In 879 Charles III [The Simple], king of ---------- (893-923), born*france
In an average lifetime, the average american eats 84,775 ----------*crackers
In an average lifetime, the average american wears 7,500 ----------*diapers
In cookery, what does the term "Julienne" mean*in strips
In cooking where does 'angelica' come from*plant root
In Ferris Buellers Day Off, who is Cameron going to marry?*The first girl he lays
In football, where are the hashmarks*five-yard lines
In greek mythology whose dogs tore actaeon apart*artemis
In greek mythology, mnemosyne is the mother of the ----------*muses
In greek mythology, what was attributed to athena*owl
In Greek mythology, who defeated Athene in a weaving contest*arachne
In greek mythology, who did jocasta marry*oedipus
In greek mythology, who was condemned to bearing the world on his shoulders for trying to storm the heavens*atlas
In greek mythology, who was jason's wife*medea
In greek mythology, who was medea's husband*jason
In Greek mythology, who was the first woman on Earth, created by Hephaestus at the request of Zeus*pandora
In Holloween, Michael Meyers wore a Halloween mask of what famous character?*Captain Kirk mask
In ice hockey, what name is given to a period of play in which one team has a player temporarily suspended from the game*power play
In India what is 'pachisi'*board game
In Indonesian cookery what name is given to meat kebabs served with a peanut sauce*satay
In italy, as what is mickey mouse known*topolino
In Knight Rider,what does K.I.T.T.'s name stand for?*Knight Industries Two Thousand
In London when was the first cricket match held at Lords*1814
In Mathematics, who devised a triangle to show the probability of various results occurring when any number of coins are tossed*blaise pascal
In military slang which word means to carry heavy equipment on foot over difficult terrain*yomp
In norse mythology, who is the chief of the valkyries*brunhilda
In order for a deck of cards to be mixed up enough to play with properly, at least how many times should it be shuffled*seven times
In physics, process of reduction of matter into a denser form, as in the liquefaction of vapor or steam*Condensation
In relation to its size, which bird has, understandably, the thickest skull*woodpecker
In Roseanne what was Roseanne's gay boss/employee?*Leon
In roulette, what number is green*zero
In Russia, what type of food is a blini or blintze*pancake
In Shakespeare's Hamlet, who is the father of Ophelia*polonius
In Shakespeare's The Merchant of Venice, what is the name of the merchant*antonio
In the 'james bond' books, to who is miss moneypenny secretary*m
In the 1938 film 'Bringing Up Baby', what was Baby*leopard
In the 1960s, Alan Reed and Jean Vander Pyle were the voices of which television husband and wife*fred & wilma flintstone
In the 1990 film 'The Krays', who played Violet Kray, the mother of the Kray brothers*billie whitelaw
In the abbreviation VDU what does the V stand for*visual
In the anglo-saxon poem, who killed grendel*beowolf
In the Bible, who led 10,000 soldiers into battle against the Midianites*gideon
In the Christian calendar, what is the alternative name for the Feast of Pentecost*whitsun
In the contract that gave cuba freedom from the us, what was required*permanent us navy base
In the culinary world, what is passata*sieved tomatoes
In the film 'dragonheart', who did the voice of the dragon*sean connery
In the film 'jurassic park', in which comical place did someone hide when the t-rex escaped*toilet
In the film version of Willy Russell's play, who played Shirley Valentine*pauline collins
In the grounds of which house is the largest private tomb/mausoleum in England*castle howard
In the law of torts, oral defamation or use of the spoken word to injure another's reputation, as distinguished from libel or written defamation*Slander
In the monty python parody 'search for the holy grail', what did patsy say when they reached camelot*it's only a model
In the monty python parody 'search for the holy grail', what was used to kill the rabbit*holy hand grenade of antioch
In the movie "Mall Rats", What famous author was signing comic books*Stan
In the movie "Mall Rats", What famous author was signing comic books?*Stan Lee
In the old gag, where is prince albert*in a can
In the parable of the Good Samaritan, to which city was the Samaritan travelling*jericho
In the TV series 'Absolutely Fabulous, who played the part of 'Bubbles'*jane horrocks
In the tv series 'the adventures of hercules', what is hercules' companion's name*iolos
In the tv series 'the brady bunch', what was cindy's toy doll's name*kitty
In the USA what are the TV equivalent of the Oscars*emmys
In the USA, what is an estate agent known as*realtor
In Welsh place names Llan- is a common feature, what does it mean*church
In what Australian state would you find Hobart*tasmania
In what Australian state would you find Horsham*victoria
In what Australian state would you find Inverell*new south wales
In what Australian state would you find Shellharbour*new south wales
In what book does 'Schahriah' appear*thousand & one nights
In what business are 'angle irons' and 'rolex'*dentistry
In what city was the final of the 1991 Canada Cup played*hamilton
In what country is K2 the world's second-highest mountain*pakistan
In what form are the signals that a normal TV aerial receives*analogue
In what is food surrounded with dry, hot, circulated air*convection oven
In what kind of restaurant might you be offered 'kulfi' as a dessert*indian
In what month is Bastille Day*july
In what profession is a 'ruderal*gardening
In what shaped ring does sumo wrestling take place*circular
In what sort of landscape would you find an erg*desert
In what sport do teams compete for the Swaythling Cup*men's table tennis
In what sport would you use spikes and blocks*athletics
In what state is silicon valley*california
In what year did Alaska become the 49th state of America*1959
In what year did Franco come to power*1937
In what year did Jean-Caude Killy win the Olympic grand slam*1968
In what year did Joseph Stalin die*1953
In what year did Rhodesia declare independence*1965
In what year did Robert the Bruce die*1329
In what year did sychronized swimming first appear in the Olympics*1984
In what year did the author Daniel Defoe die*1731
In what year did The Bayer company begin marketing heroin*1898
In what year did the Cold War begin*1946
In what year was 'Saccharin' discovered*1879
In what year was Fred Astaire born*1899
In what year was Greenpeace founded*1971
In what year was Guy Fawkes arrested*1605
In what year was insulin first used to treat diabetes*1922
In what year was Jane Fonda born*1937
In what year was Micky Mouse created*1928
In what year was NATO formed*1949
In what year was the game Monopoly invented*1929
In what year was the Taj Mahal finished*1658
In what year was Walt Disney's Snow White and the Seven Dwarfs first shown*1937
In which city is the eastern terminus of the Trans-Siberian railway*vladivostock
In which country are 'fajitas' a traditional dish*mexico
In which country do the Ashanti people live in the Province of Ashanti*ghana
In which country is the Calabria region*italy
In which country is the chief range of Drakensberg Mountains*south africa
In which country is the city of Mandalay*burma
In which country is the port of Stravangar*norway
In which country is the US naval base of Guantanamo*cuba
In which country is Tobruk*libya
In which country was film star Ray Milland born*wales
In which country was Graham Greene's novel 'A Burnt Out Case' set*belgian congo
In which country would you find McLaks (grilled salmon sandwich) on the McDonalds menu*Norway
In which country would you find the Pripyet Marshes*belarus
In which English town or city would you find The Christmas Steps*bristol
In which European Palace are the State Apartments called the Hall of Mirrors*versailles
In which film is danny devito the voice of 'phil'*hercules
In which film starring Humphrey Bogart and set in Martinique, did he play a character called Harry Morgan*to have and have not
In which film was Charlie Chaplin first heard to speak*the great dictator
In which film, starring James Cagney, with Pat O'Brien as Father Connolly did he play a character called Rocky Sullivan*angels with dirty faces
In which French island territory would you find the towns Bastia and Calvi*corsica
In which game are there hashmarks on each five-yard line*football
In which John le Carre novel does George Smiley first appear*call for the dead
In which ocean is mauritius*indian ocean
In which opera does leporello entertain a vengeful jilted lover*don
In which Puccini opera of 1896 is the Christmas Duet*la boheme
In which Shakespeare play would you find Constable Elbow*measure for measure
In which sphere of industry or commerce is the name of Arthur Maiden famous*advertising
In which sport would you find turkeys and spares*ten pin bowling
In which television series do the characters Doctor Carter and Doctor Benton appear*e r
In which weight category did John Conteh fight for the world title*light heavyweight
In which year did Lester Piggott ride his first Derby winner*1954
In which year this century were there 3 Popes*1978
In which year was aspirin invented*1899
In which year was the Battle of Copenhagen, where Nelson attacked the Danish fleet*1801
In which year was the Battle of Hastings*1066
In which year was the first artificial satellite launched*1957
In which year was the Gulf War*1991
In which year were the Olympic Games held in St. Louis*1904
Inches who at buckingham palace wears bearskins*guards
Indian song withimprovised usually topical words*calypso
Indiana jones: what creature did indy's father fear*rats
Instrument for measuring radio activity*geiger counter
Instrument for measuring wind force*anemometer
into what bay does the golden gate strait lead?*san francisco bay
Into what body of water does the yukon river flow*bering sea
Into what ocean does the Zambezi river flow*indian
Into which body of water does the river Danube flow*black sea
Iron block on which metals are worked*anvil
Islands what is ice cube's real name*o'shea jackson
Isms: Public ownership of the basic means of production, distribution, and exchange*socialism
Israel Tongue and who else devised the "Popish Plot"*titus oates
Jackdaws and magpies belong to which group of birds*crows
James hunt was disqualified after winning which grand prix*1976 british
Jamestown is the capital and chief port of which Atlantic island*st. helena
Jefferson what can be tulip, balloon or flute*wine glasses
Jimmy Carter once thought he saw a UFO; what was it*Venus
Johnny rivers sang 'secret ---------- man'*agent
Jr how many tunes blared from the 1948 wurlitzer model 1100 jukebox*twenty four
Jr*Lindbergh
K-mart. Definately. Definately K-mart*Rainman
Kainolophobia is the fear of*novelty
Kathisophobia is the fear of*sitting down
Kriss Kristofferson and Barbra Streisand starred in the re-make of which film*a star is born
La Sila lies in which region of Italy*calabria
Lack of what is the cause of the deficiency disease 'kwashiorkor'*protein
Lake Titicaca lies in which two countries*bolivia and peru
Laliophobia is the fear of*speaking
largest, rarest, and most powerful anthropoid ape?*gorilla
Lazy Susans are named after who?*Thomas Edison's daughter
Lee Which US state is known as the "Volunteer State"*tennessee
Les Paul and Charlie Christian were exponents of which musical instrument*guitar
Leukophobia is the fear of*the color white
Lewis 1994 - How many copies has the #3 "Eagles Greatest Hits" album sold*fourteen
Line of hereditary rulers*dynasty
Logophobia is a fear of ----------*words
Long necked long legged wading bird*heron
Louis xvi was guillotined in 1732, 1793 or 1842*1793
Love what does encephalitus affect*brain
Lutraphobia is the fear of*otters
Lygophobia is the fear of*darkness
Lyrics: Always spoke my mind with a gun in my hand*Ride Like the Wind Christopher Cross
Lyrics: And it's true we are immune when fact is fiction and TV reality*Sunday Bloody Sunday U2
Lyrics: Before I put another notch in my lipstick case you better make sure you put me in my place!*Hit Me With Your Best Shot Pat Benatar
Lyrics: Can make it I know I can. You broke the boy in me but you won't break the man*Man in Motion St. Elmo's Fire John Parr
Lyrics: Emotions come I don't why/Cover up love's alibi*Call Me Blondie
Lyrics: Ever since you've been leaving me I've been wanting to cry*Much Too Late For Goodbyes Julian Lennon
Lyrics: From frustration first inclination is to become a monk and leave the situation*Bust a Move Young MC
Lyrics: Good things might come to those who wait but not for those who wait too late*Just the Two of Us Grover Washington Jr. feat. Bill Withers
Lyrics: Had a premonition that he shouldn't of gone alone*Smugglers Blues Glenn Frey
Lyrics: Have some more chicken have some more pie*Eat It Weird Al Yankovic
Lyrics: He wants me but only part of the time*Voices Carry 'Til Tuesday
Lyrics: He's licking his lips he's ready to win on the hunt tonight for love at first sting*Rock You Like a Hurricane Scorpions
Lyrics: I asked the doctor to take your picture so I could look at you from inside as well*Turning Japanese Vapors
Lyrics: I can't help recalling how it felt to kiss and hold you tight*Always Something There To Remind Me Naked Eyes
Lyrics: I don't know what you expect staring into the TV set*Burning Down The House Talking Heads
Lyrics: I find myself telling you things I don't even tell my best friends*Lost in Emotion Lisa Lisa and the Cult Jam
Lyrics: I know a place where we can dance the whole night away underneath electric stars*Rhythm Of The Night DeBarge
Lyrics: I know I've been wearin' crazy clothes/And I look pretty crappy sometimes*YouBetter You Bet The Who
Lyrics: I know your plans don't include me*We've Got Tonite Kenny Rogers & Sheena Easton
Lyrics: I need fifty dollars to make you holler*Wild Thing Tone-Loc
Lyrics: I need some company a guardian angel to keep me warm when the cold winds blow*Take Me Home Tonight Eddie Money
Lyrics: I took a page out of my rule book for you*Perfect Way Scritti Politti
Lyrics: I tried my imagination but I was disturbed*867-5309/Jenny Tommy Tutone
Lyrics: I was wrong now I find just one thing makes me forget*Red Red Wine UB40
Lyrics: I'm not the kind of girl who gives up just like that*Tide Is High Blondie
Lyrics: No chocolate covered candy hearts to give away*I Just Called To Say I Love You by Stevie Wonder
Lyrics: Oh mother dear we're not the fortunate ones*Girls Just Wanna Have Fun Cyndi Lauper
Lyrics: Puts a song in this heart of mine/Puts a smile on my face every time*I Love a Rainy Night Eddie Rabbitt
Lyrics: She showed me the beach gave me a peach and pulled out the suntan lotion*Going Back To Cali LL Cool J
Lyrics: She stepped off the bus out into the city street*Fallen Angel Poison
Lyrics: She's so fine she's all mine the girl is all right!*Legs ZZ Top
Lyrics: Somewhere between the soul and soft machine/Is where I find myself again*Kyrie Mr. Mister
Lyrics: Suckin on chili dogs outside the Tastee-Freez*Jack and Dianne John Cougar Mellencamp
Lyrics: The five years we have had have been such good times*Don't You Want Me? Human League
Lyrics: The in crowd say it's cool to dig this chanting thing*Rock the Casbah The Clash
Lyrics: The moon. Beautiful. The sun. Even more beautiful*Oh Yeah Yello
Lyrics: The school bell rings you know it's my cue I'm gonna meet the boys on floor number two*Smokin' in the Boys' Room by Motley Crue
Lyrics: There's a freeway runnin' through the yard*Free Fallin' Tom Petty and the Heartbreakers
Lyrics: We could dance and party all night and drink some cherry wine*We Don't Have To Take Our Clothes Off Jeramine Stewart
Lyrics: We're even skanking to Bob Marley/ Reggae's expanding with Sly and Robbie*Genius of Love Tom Tom Club
Lyrics: Well I like takin' off don't like burnin' out every time you turn it on makes me wanna shout*Cool the Engines Boston
Lyrics: Well it's all right riding around in the breeze*End Of The Line The Traveling Wilburys
Lyrics: When I'm lost at sea I hear your voice and it carries me*Heaven is a Place On Earth Belinda Carlisle
Lyrics: Where am I to go now that I've gone too far?*Twilight Zone Golden Earring
Lyrics: Where ye goin'? What you lookin' for?*Sister Christian Night Ranger
Lyrics: Who needs a heart when a heart can be broken?*What's Love Got To Do With It? Tina Turner
Lyrics: Will you meet him on the main line or will you catch him on the rebound?*Gloria Laura Branigan
Lyrics: With every breath I'm deeper into you*Crazy For You Madonna
Lyrics: Won't you pack your bags we'll leave tonight*Two Tickets To Paradise Eddie Money
Lyrics: You can say anything you like but you can't touch the merchandise*She's a Beauty The Tubes
Mace is the outer covering of which common spice*nutmeg
Mares' tails are examples of which type of cloud*cirrus
Marie Osmond has only had one UK hit single as a solo artist name it*paper roses
Marinated limbs of fowl*chicken wings
Mark David Chapman was famous for what in 1980?*Shooting John Lennon
Marley Who still receives an estimated 25 pieces of junk mail per year at Walden Pond*Thoreau
Mass murder especially among a particular race or nation*genocide
MDMA is another name for which illegal drug*ectasy
Mechanophobia is the fear of*machines
Megalophobia is the fear of*large things
Member of a fraternity for mutual help with secret rituals*freemason
Mexico city is the capital of ----------*mexico
Michael Jackson sing this song in 1987*smooth criminal
mickey mouse has some nephews what were their names?*mortie and ferdie
Microbiophobia is the fear of*microbes
Milk, cheese and meat are good sources of which nutrient needed for a healthy diet*protein
Mixed diced vegetables in mayonnaise is what sort of salad*russian
Monte Corno, at 9554 feet, is the highest point in which Italian mountain range*apennines
Most salad dressings derive the majority of their calories from----------*fat
Mount Athos is famous for its many monasteries of which religion*greek orthodox
Movies/TV: Who starred in the film "The Ten Commandments"*charlton heston
Movies: Who played andy thompson in The Headmaster*Andy Griffith
Music artists: who did "i'd love to change the world" in 1971*ten years
Music: What band did James Brown tour and record with in the 1950's?*The Famous Flames
Music: What brother and sister duo produced a show in their family studio*donny and marie osmond
Music: what composer and organist was married twice and had 20 children*johann sebastian bach
Music: What song by Frankie Avalon went to #1 in 1959?*Venus
Music: what song of shania twain's was on the notting hill soundtrack*you've
Music: what was Steve Miller's magical incantation in 1982*abracadabra
Music: what were frankie and johnny to each other in the old song*lovers
Music: What year did Chet Atkins release his first solo album*1953
music: who is the late kurt coabain's widow?*courtney love
Music: Who re-recorded 'Secret Agent Man' in 1979?*Devo
Music: who recorded "i want you to want me" on epic records in 1979*cheap
Music: Who recorded 'Be True to your School' in 1963*The Beach Boys
Music: Who recorded 'Cuts Like a Knife' in 1983?*Bryan Adams
Music: who recorded the 1966 hit song "barbara ann"*beach boys
Music: Who recorded the 1969 hit song "Let's Work Together"*wilbert harrison
Music: Who replaced 'Bernie Leadon' of 'The Eagles' in 1975?*Joe Walsh
Music: Who sang "Everybody wants to Rule the World?"*Tears for fears
Mycophobia is the fear of*mushrooms
Mycrophobia is the fear of*small things
Myxophobia is the fear of*slime
n boy meets world,what is the crazy older brother's name?*Eric
Name captain smollett's ship in treasure island*hispaniola
Name either of the two giant stars in the constellation of Orion*rigel
Name given to that part of North America first seen in or about 986 by Bjarni Herjlfsson, who was driven there by a storm during a voyage from Iceland to Greenland?*vinland
Name one of the countries to join the Commonwealth in 1995.cameroon*mozambique
Name one type of insect belonging to the order Hymenoptera*ant
Name that car maker*alfa romeo
Name that car maker*kia
Name that car maker*mitsubishi
Name that car maker*nissan
Name that car:*aston martin
Name that car:*dodge
Name that car:*isuzu
Name that car:*jaguar
Name that car:*mitsubishi
Name that car:*saab
Name that car:*toyota
Name the carnivorous mammal related to the hyena*aardwolf
Name the character played by John Cleese in "A Fish called Wanda"*archie leach
Name the computer which beat World Chess Champion Garry Kasparov in 1997*deep blue
Name the director of the film 'American Beauty'*sam mendes
Name the female British climber while killed trying to climb K2 in 1995*alison hargreaves
Name the little elephant in books by Jean de Brunhoff*babar
Name the only actress with 4 Best Drama Actress awards*Tyne Daly
Name the original comic strip Bill The Cat appeared in*Bloom County
Name the port at the mouth of the River Seine*le havre
Name the primeval supercontinent which split into Gondwanaland and Laurasia between 250 and 300 million years ago*pangaea
Name the singer who won the 1998 Eurovision Song Contest with the song Diva*dana international
Name the stretch of water which lies between New Brunswick, Maine and Nova Scotia*bay of fundy
Name the swimmer who became a Hollywood star in the 1940's and 50s in films such as Bathing beauty and Neptune's Daughter*esther williams
Name the two movies that Michael Crichton made (before Jurrasic Park )about a theme park out of control*"West World" and "Future World"
Name the u.s. state with the smallest population*alaska
Name The Year: (France) Joan of Arc, Domremy, martyr, born*1412
Name The Year: (USA) Theodore Roosevelt dies at his home in Oyster Bay, NY, at 60*1919
Name The Year: 18th Space Shuttle Mission (Discovery 5) returns to Earth*1985
Name The Year: 1st all-talking motion picture shown, in NY*1928
Name The Year: 1st Dutch settlers arrive (from NJ), to colonize Manhattan Island*1610
Name The Year: 1st English Parliament called into session by Earl of Leicester*1265
Name The Year: 1st jet fighter used in combat (Messerschmitt 262)*1944
Name The Year: 1st outbreak of "Legionnaire's Disease" kills 29 in Philadelphia, USA*1976
Name The Year: 25,000 die in Iranian Earthquake*1990
Name The Year: 2nd Balkan War ends, Treaty of Bucharest, Bulgaria loses*1913
Name The Year: 3 cosmonauts die as Soyuz XI depressurizes during reentry*1971
Name The Year: Alfred Hitchcock knighted -- Good Evening*1980
Name The Year: Allies refuse Japan's surrender offer to retain Emperor Hirohito*1945
Name The Year: Anthony Quinn, Mexico, actor (Zorba the Greek, Lawrence of Arabia), born*1915
Name The Year: Apollo 11 returns to Earth*1969
Name The Year: Argentina seized the disputed Falkland (Malvinas) Islands from Britain*1982
Name The Year: Armistice signed ending Korean War*1953
Name The Year: Asaph Hall discovers Mars's moon Deimos*1877
Name The Year: At approx. 1:00 am Baghdad local time, allied forces attacked, beginning Gulf War*1991
Name The Year: At Waynesborough, Gen. Early's army is defeated*1865
Name The Year: Attempting to rid area of Palestine guerrillas Israel invades Lebanon*1978
Name The Year: B.C. Assyrians record total solar eclipse event on clay tablet*763
Name The Year: Bao Dai's Republic of Vietnam gains independence from France*1949
Name The Year: Barnum and Bailey's Greatest Show on Earth opens in Madison Square Garden in New York City*1881
Name The Year: Basil Rathbone Johannesburg S Africa, actor (Sherlock Holmes), born*1892
Name The Year: Battle of Dupplin Moor (in Scotland)*1332
Name The Year: Battle of Midway begins. First naval battle won in the air*1942
Name The Year: Battle of San Jacinto, in which Texas wins independence from Mexico*1836
Name The Year: Beatle Paul McCartney married Linda Eastman in London*1969
Name The Year: Beatles meet Rolling Stones for 1st time*1963
Name The Year: Beatles replace Pete Best with Ringo Starr*1962
Name The Year: Beatles sign a petition in The Times to legalize marijuana*1967
Name The Year: Bhadwan Shree Rajneesh, indian guru, dies at 58*1990
Name The Year: Bill Haley Mich, (andthe Comets-Rock Around the Clock), born*1925
Name The Year: Billy Crystal, comedian, born*1947
Name The Year: Blake Edwards, writer/director (Breakfast at Tiffany's), born*1922
Name The Year: Brian Jones founder of the Rolling Stones, drowns*1969
Name The Year: British naval forces raid Nazi occupied French port of St. Nazaire*1942
Name The Year: British under Adm Horatio Nelson beat French at Battle of Nile*1798
Name The Year: Burundi and Rwanda gain independence from Belgium (National Days)*1962
Name The Year: Captain Cook runs aground on Australian Great Barrier Reef*1770
Name The Year: Carl Lewis runs 100m in 9.86 seconds*1991
Name The Year: Carlos Santana Mexico, musician (Santana-Black Magic Woman), born*1947
Name The Year: Carol Channing, actress (Gentlemen Prefer Blondes), born*1921
Name The Year: Cary Grant actor (Arsenic & Old Lace, North by Northwest), born*1904
Name The Year: Charles Lindbergh, died at his home in Hawaii at the age of 72*1974
Name The Year: Charles VII, Holy Roman emperor (1742-45), born*1697
Name The Year: Charlotte, grand duchess of Luxembourg (1919-64), born*1896
Name The Year: China leases Hong Kong's new territories to Britain for 99 years*1898
Name The Year: Chinese republic proclaimed in Tibet*1912
Name The Year: Chris Young Penn, actor (Bryce Lynch-Max Headroom, Great Outdoors), born*1971
Name The Year: Christopher Isherwood, novelist, playwright (I Am a Camera)*1904
Name The Year: Christopher Wilder, FBI's 'most wanted man' accidentally killed self*1984
Name The Year: Cindy Williams (in Van Nuys, CA), actor (Laverne and Shirley), born*1948
Name The Year: Clement Clarke Moore, American author ('Twas the Night Before Xmas)*1779
Name The Year: Clint Eastwood elected mayor of Carmel, California. It made his day*1986
Name The Year: Communist coup is crushed in USSR in 2 days*1991
Name The Year: Congress creates the Territory of Nevada*1861
Name The Year: Constantius II, Roman emperor (337-61), born*317
Name The Year: Construction of Cologne Cathedral is begun*1248
Name The Year: Copernicus makes his 1st observations of Saturn*1514
Name The Year: Curtis Mayfield, musician, born*1942
Name The Year: Dan Aykroyd, Ottawa Canada, comedian/actor (SNL, Dragnet), born*1952
Name The Year: David Bowie releases "Fame"*1975
Name The Year: David Hasselhoff, actor, born*1952
Name The Year: Dean Martin, singer, actor*1917
Name The Year: Diana Ross (in Detroit, Michigan), singer (The Supremes), born*1944
Name The Year: Dick Sargent Carmel Calif, actor (Darrin-Bewitched), born*1933
Name The Year: Dodge Morgan sailed solo nonstop around the world in 150 days*1986
Name The Year: Donald Sutherland, actor (M-A-S-H), born*1934
Name The Year: Donald Trump master builder (Trump Towers/Plaza/Castle), born*1946
Name The Year: Donus ends his reign as Catholic Pope*678
Name The Year: Douglas MacArthur US general (Pacific theater-WW II), dies at 84*1964
Name The Year: Dr. Albert Sabin, polio vaccine discoverer*1906
Name The Year: Echo I, first passive satellite launched*1960
Name The Year: Edgar Allan Poe, Boston, author (Pit & the Pendulum), born*1809
Name The Year: Edouard Manet, French painter, born*1832
Name The Year: Elizabeth Alexandra Mary Windsor II, queen of England (1952- ), born*1926
Name The Year: Elvis Presley records his debut single, "That's All Right"*1954
Name The Year: Emiliano Zapata, Mexican revolutionary, peasant leader*1879
Name The Year: Emma Lazarus, whose poem was inscribed on the Statue of Liberty, born*1849
Name The Year: English defeat French at Battle of Blenheim*1704
Name The Year: Eric Clapton, guitarist, singer, born*1945
Name The Year: Ernest Hausen of Wisconsin sets chicken-plucking record-4.4 sec*1939
Name The Year: Esther Williams (in Inglewood, CA), actor, swimmer*1923
Name The Year: Esther Williams, swimmer, actor, born*1923
Name The Year: European community proposes a boycott of Iraq*1990
Name The Year: European Space Agency launches Giotto Sattelite to Halley's Comet*1985
Name The Year: Everlasting League forms, basis of Swiss Confederation (Nat'l Day)*1291
Name The Year: F.B.I. begins it's "10 most wanted list"*1950
Name The Year: Federal Bureau of Investigation established*1908
Name The Year: Federico Fellini Italian director (Satyricon, La Dolce Vita), born*1920
Name The Year: Ferdinand Magellan world traveler, killed by Filipino natives*1521
Name The Year: Fidel Castro leads attack on Moncada Barracks, begins Cuban Revolution*1953
Name The Year: Fifteenth Space Shuttle Mission - Discovery 3 returns to Earth*1985
Name The Year: First all-color television station to televise live local programs*1986
Name The Year: First American expeditionary force to land in Africa (WW II)*1942
Name The Year: First Boeing B-29 arrives in China "over the Hump"*1944
Name The Year: First Cable Car is patented by Andrew S. Hallidie*1871
Name The Year: First drinking straw is patented by M.C. Stone in Washington, D.C*1888
Name The Year: First electric razor marketed by Schick, Inc*1931
Name The Year: First Israeli election*1949
Name The Year: First jazz record in United States is cut*1917
Name The Year: First known auto race*1895
Name The Year: First man-powered flight (Bryan Allen in Gossamer Condor)*1977
Name The Year: First pineapples planted in Hawaii*1813
Name The Year: First successful helicopter flight, Stratford, Ct*1940
Name The Year: First telegraph company in Hawaii opens*1901
Name The Year: First televised tennis match*1928
Name The Year: First transatlantic jet passenger trip*1950
Name The Year: Flight 255 out of Metro Airport in Detroit crashes just miniutes after take off, killing all but one small child*1987
Name The Year: France declares war on Austria, Prussia & Sardinia*1792
Name The Year: Frances Drake completres circumnavigation of the world*1581
Name The Year: French-Egyptian forces under Napolean I beat Turks at Battle of Abukir*1799
Name The Year: Gemini 5 returned after 12 days, 7 hours, 11 minutes, 53 seconds*1965
Name The Year: Gen. Douglas MacArthur left Bataan for Australia*1942
Name The Year: Gene Hackman, actor (Target, Uncommon Valor), born*1930
Name The Year: George Cormack, the inventor of "Wheaties" cereal*1870
Name The Year: George III of England, king, born*1738
Name The Year: George Michael (in England), singer, born*1963
Name The Year: George Washington creates the Order of the Purple Heart*1782
Name The Year: Germany declares war on Russia in WW I*1914
Name The Year: Germany declares war on Soviet Union during WW II*1941
Name The Year: Gherman S. Titov, second Russian in space aboard Vostok 2*1961
Name The Year: Gottlieb Daimler, automobile pioneer, born*1834
Name The Year: Gough Whitlam (ALP) Australia, PM (1972-75), born*1916
Name The Year: Grace Kelly marries Prince Rainier of Monaco (civil ceremony)*1956
Name The Year: Ground breaking for Disneyland, the Magic Kingdom, in Anaheim, CA*1954
Name The Year: Helen Keller, blind-deaf author-lecturer*1880
Name The Year: Henry VII king of England (1485-1509), born*1457
Name The Year: Hernando De Soto claims the US state of Florida for Spain*1539
Name The Year: Hitler breaks Treaty of Versailles by sending troops to Rhineland*1936
Name The Year: Homing pigeon completes 11,000 km trip (Namibia-London) in 55 days*1845
Name The Year: Honor Blackman (in London, England), actor (The Avengers), born*1929
Name The Year: Hope Emerson, actress (I Married Joan, Peter Gunn), dies at 62*1960
Name The Year: Hunter S. Thompson, gonzo journalist*1949
Name The Year: Imelda Marcos former 1st lady (Philipines)/shoe collector, born*1930
Name The Year: In accordance with Camp David, Israel completes Sinai withdrawl*1982
Name The Year: International Women's Day*1945
Name The Year: Israel destroys alleged Iraqi plutonium production facility*1981
Name The Year: Israel, Syria, Jordan, Iraq & Egypt end "6-Day War" with UN help*1967
Name The Year: Italy beats Czechoslovakia 2-1 (OT) in soccer's 2nd World Cup at Rome*1934
Name The Year: Jacobite Scottish Highlanders defeat royal force at Killiecrankie*1689
Name The Year: James Cagney died at his farm in Stanfordville, NY, at age 86*1986
Name The Year: James Hetfield heavy metal rocker (Metallica), born*1963
Name The Year: Japanese forces on Okinawa surrender to US during WW II*1945
Name The Year: Jean Rey, of Belgium, president of European Commission (1967-70)*1902
Name The Year: Jesse James shot dead in St. Joseph Mo. by Robert Ford*1882
Name The Year: Jim Belushi, Chicago Ill, comedian (Sat Night Live, Trading Places), born*1954
Name The Year: Joe Frazier, boxer, born*1944
Name The Year: John Astin, actor (Gomez in TV Addams Family), born*1930
Name The Year: John Constable (in England), painter, born*1776
Name The Year: John D. Rockefeller, financier, born*1839
Name The Year: John Dean begins testifying before the Senate Watergate Committee*1973
Name The Year: John L. Sullivan wins by KO in 75 rounds in last bareknuckle bout*1889
Name The Year: John Landis actor (American Werewolf in London), born*1950
Name The Year: John Presper Eckert, co-inventor of first electronic computer (ENIAC), born*1919
Name The Year: Judy Garland, singer/actress, dies at 48 of an alcohol overdose*1969
Name The Year: King Camp Gillette, inventor of the safety razor, born*1855
Name The Year: Kiribati (Gilbert and Ellice Is.) gains independence from Britain*1979
Name The Year: Landslide in Huancavelica Province Peru creates a natural dam*1974
Name The Year: Led Zepplin's Debut Album released*1969
Name The Year: Lenin returns to Russia to start Bolshevik Revolution*1917
Name The Year: Leonard Bernstein conductor/composer/pianist/egotist, born*1918
Name The Year: Leonard Nimoy, actor, director, born*1931
Name The Year: Les Paul, Waukesha Wisconsin, U.S.A., guitarist/inventor (Les Paul guitar), born*1915
Name The Year: Lithuanian SSR is accepted into the USSR*1940
Name The Year: Little Richard gets a star on Hollywood's walk of fame*1990
Name The Year: Louis XIV crowned King of France*1654
Name The Year: Louis XVI, king of France (1774-92); guillotined, born*1754
Name The Year: Ludwig II mad king of Bavaria (1864-86), born*1845
Name The Year: Mahatma Gandhi's 1st arrest, campaigning for Indian rights in S Africa*1914
Name The Year: Mahmud I Ottoman sultan, fought Austrians and Russians, born*1696
Name The Year: Manuel Quezon, first president of Philippine Commonwealth (1935-42), born*1878
Name The Year: March by civil rights demonstrators was broken up in Selma, Alabama*1965
Name The Year: Maria Shriver & Arnold Schwarzenegger marry*1986
Name The Year: Mark Russell, raconteur, born*1932
Name The Year: Michael Fish British TV weatherman, born*1944
Name The Year: Michael Landon actor (Bonanza, Highway to Heaven), dies at 54 from cancer*1991
Name The Year: Miguel Vasquez makes first public quadruple somersault on trapeze*1982
Name The Year: N.Y. Highlander (Yankees) tickets first go on sale*1903
Name The Year: Natalie Wood [Natasha Gurdin], SF, (Gypsy, Rebel Without a Cause), born*1938
Name The Year: National Socialist (Nazi) Party formed in Germany*1919
Name The Year: Nelson Mandella, human rights activist, former political prisoner*1918
Name The Year: New socialist constitution of East Germany takes effect*1968
Name The Year: Nineteenth Space Shuttle Mission - Challenger 8 returns to Earth*1985
Name The Year: Oliver Hardy Harlem Ga, comedy team member (Laurel & Hardy), born*1892
Name The Year: Orville Redenbacher, popcorn king*1907
Name The Year: Patrick Swayze Houston Tx, actor/dancer (Dirty Dancing, Ghost), born*1952
Name The Year: Paul Carrack rocker (Squeeze/Ace-How Long), born*1951
Name The Year: Percy Bysshe Shelley England, romantic poet (Adonais), born*1792
Name The Year: Peter Ilyich Tchaikovsky, Russian composer (1812 Overture), born*1840
Name The Year: Phil Ochs rock producer, dies*1976
Name The Year: Philippines gains independence from US*1946
Name The Year: Pink Floyds' "The Wall" is performed where the Berlin Wall once stood*1990
Name The Year: Pioneer-Venus 2 Multi-probe launched to Venus*1978
Name The Year: Potsdam Conference (Roostevelt, Stalin, Churchill) holds first meeting*1945
Name The Year: Pres. Eisenhower signs into law National Aeronautics and Space Administration (NASA) and Space Act of 1958*1958
Name The Year: Princess Ingeborg of Sweden, born*1878
Name The Year: Queen Mother Wilhelmina Netherlands (1890-1948), born*1880
Name The Year: Ransom Eli Olds, auto (Oldsmobile) and truck (REO) manufacturer, born*1864
Name The Year: Rev. William Archibald Spooner, invented 'spoonerisms' (ie When you get you bords wackwards and you can't palk troperly.), born*1844
Name The Year: Richie Sambora guitarist (Bon Jovi-You Give Love a Bad Name), born*1959
Name The Year: Ringo Starr, Beatles' drummer, born*1940
Name The Year: Roald Amundsen, Norwegian explorer, discoverer of South Pole*1872
Name The Year: Rob Lowe, actor, famous for home-made movies, born*1964
Name The Year: Robert Crumb cartoonist (Father Time), born*1943
Name The Year: Robert Goddard, rocketry pioneer, died*1945
Name The Year: Robert Mitchum, actor, born*1917
Name The Year: Rod Stewart, singer, born*1945
Name The Year: Ron Kovic disabled vietnam vet (Born on 4th of July was based on his life), born*1946
Name The Year: Rose Kennedy, Mother of a President, an Attorney General, and a Senator, born*1890
Name The Year: Royal Air Force established in Britain*1918
Name The Year: Rutger Hauer, actor (Blade Runner, Ladyhawke, Osterman Weekend), born*1944
Name The Year: Samuel Goldwyn, pioneer filmmaker*1882
Name The Year: Sean Connery actor (James Bond, Man Who Would Be King), born*1930
Name The Year: Senegalese National Day begins*1960
Name The Year: Sergei Rachmaninoff (in Novgorod Province, Russia), composer, born*1873
Name The Year: Seve Ballesteros, golfer, born*1957
Name The Year: Sierra Leone becomes a republic (Natl Day)*1971
Name The Year: Sir Henry Havelock British soldier (War in Afghanistan 1838-39), born*1795
Name The Year: Sir Stamford Raffles founded Singapore, born*1781
Name The Year: Sir Walter Scott, Scottish novelist, poet (Lady of Lake, Ivanhoe), born*1771
Name The Year: South Africa passes Group Areas Act segregating races*1950
Name The Year: South Korean President Park Chung-Hee escaped an assassination*1974
Name The Year: Soviet troops enter Berlin*1945
Name The Year: Soyuz 28 is launched*1978
Name The Year: Soyuz T-11 is launched*1984
Name The Year: Soyuz T-9 is launched*1983
Name The Year: Spain declares war on U.S., rejecting ultimatum to withdraw from Cuba*1898
Name The Year: Spanish Civil War end as Madrid fell to Francisco Franco*1939
Name The Year: Speech by Khrushchev blasting Stalin made public*1956
Name The Year: St Paul I ends his reign as Catholic Pope*767
Name The Year: St. Frances Xavier Cabrini (Mother Cabrini), first US saint*1850
Name The Year: St. Vladimir's Day*1918
Name The Year: Steve Tyler, Aerosmith's lead singer, born*1948
Name The Year: Suez Canal reopens (after 6 Day War caused it to close)*1975
Name The Year: Sultan of Turkey Abdul Hamid II is overthrown*1909
Name The Year: Suzannah York, actress, born*1941
Name The Year: Sweden's constitution adopted*1809
Name The Year: Sylvester Stallone NYC, actor/director (Rocky, Rambo, Cobra), born*1946
Name The Year: The first patent is granted for a fire escape ... a wicker basket on a pully and a chain, designed by a London watchmaker*1766
Name The Year: The game "Monopoly" is invented*1933
Name The Year: The Salvation Army of England sends group to U.S. to begin welfare and religious activity here*1888
Name The Year: Theophilus Van Kannel of Philadelphia receives a patent for his revolving door -- described as a storm door structure*1888
Name The Year: Thomas Davenport, invented the first commercially successful electric motor, born*1802
Name The Year: Tony Curtis [Real Name : Bernard Schwartz], Bronx New York, actor (Some Like it Hot), born*1925
Name The Year: Total solar eclipse captured on a daguerreotype photograph*1851
Name The Year: TV game show scandal investigation starts*1958
Name The Year: U.S. launches Pioneer Venus probe*1978
Name The Year: U.S. sub locates missing hydrogen bomb in Mediterranean*1966
Name The Year: U.S. Viking 2 goes into Martian orbit after an 11-month flight from Earth*1976
Name The Year: UN Charter signed by 50 nations in SF*1945
Name The Year: US actress Grace Kelly marries Monaco's Prince Rainier III*1956
Name The Year: US and USSR sign Lend-Lease agreement during World War II*1942
Name The Year: US declares war on Germany (WWI)*1917
Name The Year: US drops second atomic bomb on Japan destroying part of Nagasaki. An estimated 74,000 people died. The original target was Kokura*1945
Name The Year: USSR launches Mars 6*1973
Name The Year: Venera 3, Venus landing*1966
Name The Year: Veronica Lake actress, dies at 58*1973
Name The Year: Victor Borge, pianist, comedian, born*1909
Name The Year: Victoria Principal (in Japan), actor (Dallas), born*1950
Name The Year: Voyager 2 begins a flyby of the planet Neptune*1989
Name The Year: Voyager 2 discovers 2 partial rings of Neptune*1989
Name The Year: Walt Disney's "Sleeping Beauty" released*1959
Name The Year: Walter Baade discovers asteroid Icarus inside orbit of Mercury*1949
Name The Year: Wembley Stadium opens-Bolton Wanderers vs West Ham United (FA Cup)*1923
Name The Year: William Gilbert Grace, Victorian England's greatest cricketer*1848
Name The Year: Winston Churchill resigns as British PM, Anthony Eden succeeds him*1955
Name The Year: Worst nuclear disaster, Chernobyl USSR, 31 die*1986
Name The Year: Yugoslavia elects it's 1st president (Marshal Tito)*1953
Named album of the year in 1981, which pop group's debut album was called "Dare"*human league
Names what portable object is the teleram t-3000*computer
Narcolepsy is the uncontrollable need to ----------*sleep
Narrow trench made by a plough*furrow
National capitals: Costa Rica*san jose
Ncaa: in what year was the heisman memorial trophy first awarded*1935
Neptune was the roman god of the ----------*sea
Ness What are panatelas*cigars
New Zealand's Rugby team is know as the ----------*All Blacks
Newkirk c3p0 is the first character to speak in which film*star wars
Nosocomephobia is the fear of*hospitals
Oaks of dodona what did the white house have before it had an indoor bathroom*telephone
Of what continent is cyprus a part*asia
Of what country is the monetary unit the rupee*india
Of what did robert the bruce, king of scotland, die in 1329*leprosy
Of which country is Amharic an official language*ethiopia
Of which metal is sperrylite the ore*platinum
Of which Spanish province is Seville the capital city*andalucia
Of who did the u.s postal service print 500 million stamps in 1993*elvis
Officers in which army were given copies of 'les miserables'*confederate
Ombrophobia is the fear of*rain
Ommetaphobia is the fear of*eyes
On a dartboard, what number is on top*twenty
On Airwolf, what instrument does Hawke play*cello
On FRIENDS what was the name of Ross's monkey?*Marcel
On Full House,what was Jesse's REAL first name?*Hermes
On Little House on the Prairie,what was Laura's horse's name?*Bunny
On MASH,what was Walter 'Radar' O'Reilley's home town?*Ottumwa, Iowa
On maps, what is the 'you are here' arrow*ideo locator
On Night Court,Harry had a "statue" of what animal in his office?*Armidillo
On the 1976 release, who 'wanted to fly like an eagle'*steve miller band
On three's company,what was Chrissy's father's ocupation?*A Reverend
On what does the firefly depend to find mates*sight
On what scale are there 180 degrees between freezing point & boiling point*fahrenheit scale
On what sea is the crimea*black sea
On what show did Dano get to book the bad guy?*Hawaii 5-0
On what street in new rochelle did rob and laura petrie live*bonnie meadow
On which Caribbean island are the Blue Mountains*jamaica
On which continent would you be standing if you were visiting the Republic of Surinam*south america
On which day of the week is the Moslem Sabbath*friday
On which island are the Troodos mountains*cyprus
On which major river are The Owen Falls dam*nile
On which object would you find a crown, a waist, a sound-bow and a clapper*bell
On which river does Berlin stand*spree
On which U. S. river is the Grand Coulee Dam*columbia
Onassis driving: what country is identified by the letters ma*morocco
One of the worst fires in American history gutted the twenty-six storey MGM Grand Hotel in 1988. In which city was the hotel situated*las vegas
organ of the digestive system?*stomach
Oriental market*bazaar
Original inhabitants of New Zealand, of Polynesian stock*maori
Osteoporosis primarily affects*bones
Other than the U.K. and Eire, name a European country where cars are driven on the left hand side of the road*cyprus
Pagophobia is the fear of*ice
Pants*green jacket grey pants
Paralipophobia is the fear of*neglecting duty
Parton what is the official birthplace of country music*bristol
Pasteur developed a vaccine for rabies in which year*1885
Pathophobia is the fear of*disease
Patsy cline is the most noted with pop-country crossovers. which other singer should not be overlooked for her hits 'break it to me gently' and 'fool no. 1'*brenda lee
Pavarotti popularized Nessun dorma but what does it mean*none shall sleep
Pediophobia is the fear of*dolls
Percent what was the final destination of the first u.s. paddle wheel steamboat, what departed from pittsburgh*new orleans
Philip Pirrip is the main character in which Charles Dickens novel*great expectations
Phineas Barnum opened his circus in what year*1871
Phonophobia is a fear of ----------*voices
Phthiriophobia is the fear of*lice
Plant what city did general sherman burn in 1864*atlanta
Poem or song narrating popular story*ballad
Point Maley is the coast guard cutter in what Disney movie*boatniks
Porphyrophobia is the fear of*the color purple
Port Louis is the capital of which island state in the Indian Ocean*mauritius
President richard m nixon called what songstress an "ambassador of love"*pearl baily
President Roosevelt had a landslide victory in 1932, who did he defeat*herbert hoover
Presley elvis presley appeared on how many stamps in 1993*five hundred million
Promotion of friendly relations between countries*bridge-building
Psychrophobia is the fear of*cold
Pussycat sings 'now the country song forever lost its soul, when the guitar player turned to rock n roll ----------' what's the song title*mississippi
QANTAS, the name of the airline, is an acronym for*queensland and northern territory aerial services
Queen Berengaria never came to England, although she was married to the King. Which King*richard the first
Quite a Year for Plums----------*bailey white
Quotations: "--------- as if everything depended on God, and work as if everything depended upon man."- Cardinal Francis J. Spellman*pray
Quotations: "...do your -------------. You can't lead without knowing what you're talking about..."- George Bush (1925 - )*homework
Quotations: "A billion here, a billion there - pretty soon it adds up to real money."*Everett Dirksen
Quotations: "Christmas is over and Business is Business."- Franklin Pierce Adams*business
Quotations: "Do not wait for leaders; do it alone, person to person."*Mother Teresa
Quotations: "Here is the test to find whether your ------------ is finished: if you're alive, it isn't."- Richard Bach*mission on earth
Quotations: "I don't want to achieve immortality through my work, I want to achieve it through not dying."*Woody Allen
Quotations: "If I were --------------, would I be wearing this one?"- Abraham Lincoln (1809-1865)*two-faced
Quotations: "If it weren't for -------------- I'd have no sex life at all."- Rodney Dangerfield*pickpockets
Quotations: "If men could get -------------, abortion would be a sacrament."- Florence R. Kennedy*pregnant
Quotations: "If someone says It's not the money, it's the -------------,'it's the money.'"- Angelo Valenti*principle
Quotations: "If you cannot get your ---------- to call you, try not paying his bill."- Pete Ferguson*lawyer
Quotations: "The -------- of money is the root of all evil."- The apostle Paul*love
Quotations: "The human race has one really effective weapon, and that is ---------."- Mark Twain (1835 - 1910)*laughter
Quotations: "Women who ----------- are called 'mothers'."- Abigail Van Buren*miscalculate
Quotations: "Work is a necessity for man. Man invented the ---------------."- Pablo Picasso*alarm clock
Quotations: "Years may wrinkle the skin. Lack of ---------- will wrinkle the soul."- Anonymous*enthusiasm
Quotations: "You can observe a lot by ------------."- Yogi [Lawrence Peter] Berra*watching
Quotes: I don't want to achieve immortality through my work, I want to achieve it through not dying*Woody Allen
Quotes: In his private heart no man much respects himself*Mark Twain
Rabbits like ----------*licorice
Reddish-brown colour alluding to hair*auburn
Relating to cookery what are 'lokshen', used in a type of Jewish soup*noodles
Relating to food what is 'halloumi'*cypriot cheese
Relating to or using signals over a range of frequencies*broadband
Republic in southern central America, bounded on the north by Nicaragua, on the east by the Caribbean Sea, on the southeast by Panama, & on the southwest & west by the Pacific Ocean*costa rica
Richard Gere was married to which model*cindy crawford
River Providence is the capital of what state*rhode island
Rustic or awkward person*bumpkin
S.American cowboy*gaucho
Saigon is the capital of ----------*south vietnam
Saintpaulia is the botanical name for which houseplant*african violet
Saturday is named for which planet*saturn
Scoleciphobia is the fear of*worms
Scottish sailor Alexander Selkirk became inspiration for what novel*robinson
Scriptophobia is a fear of ----------*writing in public
Second city: Cheyenne (state)*casper
Serotine, Leislers and Noctule are all varieties of which nianinial*bat
She won the 1979 Nobel peace prize for her work among the poor*mother teresa
Shinguards were introduced into football in which year*1839
Shop selling exotic cooked meats and cheeses*delicatessen
Short legged long bodied dog*dachshund
Sieze control of vehicle*hijack
Six ounces of orange juice contains the minimum daily requirement for which vitamin*vitamin c
Sixty what lives in a fornicary*ants
Slang:A promiscuous woman*slapper
Slave trading was abolished in the british empire in 1807, 1825 or 1855*1807
Soceraphobia is the fear of*parents-in-law
Solar time what's the usual age for a jewish boy to celebrate his "bar mitzvah"*thirteen
Southern Comfort is made from a base of Bourbon whiskey and flavouring from which fruit*peach
Space indiana jones: what did drinking from the grail "grant"*immortality
Spectrophobia is the fear of*specters
Squid, octopus and cuttlefish are all types of what*cephalopods
St christopher the patron saint ----------*travellers
Starring Nigel Hawthorne, which 1994 film was publicised with "His Majesty was all-knowing. But he wasn't quite all there."*the madness of king george
Stoppered glass container for wine or spirits*decanter
Stygiophobia is the fear of*hell
Sudden overthrow of government*coup d'etat
Super glue is used to lift fingerprints from what surfaces*difficult
Supposed paranormal force moving objects at a distance*telekinesis
Sydney 2000 Olympics: This countries medal tally was: 0 Gold, 1 Silver, 0 Bronze, 1 in Total*uruguay
Sydney 2000 Olympics: This countries medal tally was: 0 Gold, 2 Silver, 2 Bronze, 4 in Total*argentina
Sydney 2000 Olympics: This countries medal tally was: 0 Gold, 6 Silver, 6 Bronze, 12 in Total*brazil
Sydney 2000 Olympics: This countries medal tally was: 2 Gold, 1 Silver, 1 Bronze, 4 in Total*finland
Sydney 2000 Olympics: This countries medal tally was: 6 Gold, 5 Silver, 3 Bronze, 14 in Total*poland
Tachophobia is a fear of ----------*speed
Talc is a hydrated silicate of which metal*magnesium
Tarlike mixture of hydrocarbons derived from petroleum*bitumen
The 'love apple' is more commonly known as what*tomato
The 'purple heart' medal was created in 1668, 1701 or 1782*1782
The 1st US minimum wage law was instituted in what year*1938
The actor who played captain sisko in 'star trek deep space nine', played ---------- the 1970's series 'spencer for hire'*hawk
The alcohol found in wine, beer & liquor is known as grain alcohol or what*ethanol
The assassination of what country's Archduke led to World War I?*Austria
The assault on Starfleet by the Borg was at*wolf 359
The Atlanta Hawks basketball team have retired 23 which used to belong to ----------*lou hudson
The basis of all scientific agriculture, what involves six essential practices: proper tillage; maintenance of a proper supply of organic matter in the soil; maintenance of a proper nutrient supply; control of soil pollution; maintenance of the correct soil acidity; & control of erosion*soil management
The bering strait lies between russia and ----------*alaska
The canary islands in the pacific are named after what animal*dog
The childrens story 'The Rose and The Ring' was written by which 19th century novelist*william thackeray
The Chinese ideograph with two women under one roof means what?*Trouble
The coast line around this lake in North Dakota is longer than the California coastline along the Pacific Ocean*Lake Sakakawea
The cocktail "Margarita" contains cointreau, lime and which spirit*tequila
The country name for which bird is 'merle'*blackbird
The date of which christian festival was fixed in 325 ad by the council of nicaea*easter
The Dirty Harry franchise ran to five films what was the title of the final 1988 film*the dead pool
The earths atmosphere & the space beyond is known as ----------*aerospace
The filament of a regular light bulb is usually made of ----------*tungsten
The first charity flag day was held in 1914, 1917 or 1919*1914
The first nude Playboy centerfold was*marilyn monroe
The first person to swim the English Channel did so in what year*1875
The first telephone call was made in what year*1876
The force that brings moving bodies to a halt is ----------*friction
The great gothic cathedral of Milan was started in 1386, & wasn't completed until what year*1805
The Guarani is the unit of currency in which South American country*paraguay
The ice cream soda was invented in what year*1874
The Inquisition forced this person to recant his belief in the Coppernican Theory. Who was he?*galileo
The Irish Province of Connaught contains five counties. Sligo and Galway are two. Name one of the others. leitrim*mayo
The Jeffersons was a spinoff from what show?*All in the Family
The largest internal organ of the human body is*liver
The latin qed spells out in full as*quod erat demonstrandum
The left lung is smaller than the right lung to make room for what*heart
The longest bike weighed how much*more than a ton
The mathematical notation for a summation is designated by what greek letter*sigma
The minimum number of members required to be legal is known as a*quorum
The most abundant metal in the earths crust is what*aluminum
The most famous church in Great Britain, enshrining many of the traditions of the British people*Westminster Abbey
The most northerly point of mainland Africa is in which country*tunisia
The most prominent of the 12 disciples of Jesus Christ, a leader and missionary in the early church, and traditionally the first bishop of Rome*peter
The name of which constellation me 'harp'*lyra
The name of which disease comes from the Italian meaning 'bad air'*malaria
The name of which of the seven hills of Rome is the origin of the word 'palace'*palatine hill
The name of which plant comes from the Greek meaning 'earth-apple'*camomile
The nest of an eagle or bird of prey is an*eyrie
The normal temperature of a cat is ---------- degrees (it's a decimal)*101.5
The northern part of north america lies within the ----------*arctic circle
The observable activity of an "individual.(----------)*behaviour
The olympic motto 'citius, altius, fortius' means what*faster, higher,
the only member of the band zz top without a beard has what last name?*beard
The ore pitchblende is the major source of which element*uranium
The peace of Aix-la-Chapelle was celebrated by which piece of music*music for the royal fireworks
The phillips head screwdriver was invented where*oregon
The plant life in the oceans make up about what percent of all the greenery on the earth*85
The Prince of Demons in the new testament was called ----------*beelzebub
The process of splitting atoms is called*fission
The Sailor Who Fell From Grace With the Sea----------*yukio mishima
The Simplon Tunnel runs between which two countries*italy & switzerland
The skin of which animal is used to make Morocco Leather*goat
The St. Valentine's day massacre took place in this city*chicago
The study of shells*conchology
The telephone was invented in which year*1876
The television detective Banacek was played by whom*george peppard
The treatment of disease by chemical substances which are toxic to the causative micro organisms is called ----------*chemotherapy
the two rival gangs in "west side story" were the sharks and the ----------?*jets
The u.s has never lost a war where they used ----------*mules
The University of Houston once elected what rock star as homecoming queen*alice cooper
The variety of living organisms in a particular habitat or geographic area*biodiversity
The Voyage of the Beagle told of which scientist's discoveries?*Charles Darwin
The word "angel" is derived from the Greek term angelos, from the Hebrew experssion mal'akh, usually translated as what?*Messenger
The word 'boondocks' comes from the tagalog (filipino) word 'bundok,' which means*mountain
The word 'whisky' comes from Gaelic, what does it mean*water of life
The words 'dungarees' and 'jungle' originate from which language*hindi
The' Long John Silver Collection' housed on the Cutty Sark is the nations largest collection of what*ship's figureheads
There are 16 ---------- in a cup*tablespoons
There are 318,979,564,000 possible ways of playing just the first four moves on each side in a game of*chess
There are 45 miles of what in the skin of a human being*nerves
There are more statues of ----------, Lewis & Clarks female indian guide, in the U S than any other person*Sacagawea
There were three Kings of England in 1066. Harold and William, of course, and who else*edward the confessor
These attach muscles to bones or cartilage*Tendons
These rabbits are prized for their long, soft fur, used to make very expensive sweaters*angorra
Thick, light yellow portion of milk from which butter is made*cream
Thinker who sang the 1963 hit 'it's my party'*lesley gore
This animal is found at the beginning of an encyclopedia*aardvark
This company uses the slogan AOL*america on line
This country consumes more coca cola per capita than any other*iceland
This fingerlike projection is attached to the large intestine*appendix
This is the hardest naturally occurring substance*diamond
This island group is off the east coast of southern South America*falkland islands
This island was Ulysses' home*ithaca
This membrane controls the amount of light entering the eye*Iris
This place in Germany is also the name of a (popular) cake*black forest
This science deals with the motion of projectiles*ballistics
This space station killed a cow on re entry into earth's atmosphere*skylab
this teen was sentenced to a public caning in singapore in 1994?*michael fay
This U S state touches 4 of 5 great lakes*michigan
This vegetable is a variety of broccoli*calabrese
This was the site of worse nuclear accident in history*chernobyl
Thomas Magnum's dad was played by what actor?*Robert Pine
Thousand four hundred marconi transmitted radio signals across the atlantic in 1901, 1902 or 1903*1901
Through what were dead Egyptian pharaohs' brains extracted*nasal passages
Thumper was a rabbit from which film*bambi
Time of the Season (1969) was done by what group*zombies
Time ---------- when your having fun*flies
To the nearest minute, how long does it take sunlight to reach earth*eight
To what country would a hiker go to assail mt ararat*turkey
To what do the tendons attach the muscles*bones or cartilage
To what does the original term' cutty sark ' refer*short shift
To what family of vegetables does the popular Zucchini or Courgette belong*gourd or squash
To what instrument family do "french horns" belong*brass
To which country do the Coral Sea Islands belong*australia
To which instrument does an orchestra normally tune*oboe
To which plant family (strictly genus) do jonquils and daffodils belong*narcissus
To which team did marlboro switch its backing from brm in the 1974 season*mclaren
To within 30 feet, how tall is the Eiffel Tower*nine hundred & eighty four
Tom hallick was the first male host of which show*entertainment tonight
Transom, poop and keel are all parts of a what*boat
Trees: which tree has catkins in the spring and edible nuts in the autumn*hazel
Tropical shrub used for making hair dye*henna
Tropical tree bearing edible orange fruit*guava
True or false: contrary to popular belief, a lightbulb actually absorbs darkness*false
TV/Movies: "Dozens of people are dying all the time, thousands, so why not mother"*julian
TV/Movies: "Id like to tame her shrew!!!"*back to school
TV/Movies: "stuart little" was a story about a ----------*mouse
TV/Movies: 1914 - Charlie Chaplin - Starred In This Movie:*tillie's punctured romance
TV/Movies: 1925 - Charlie Chaplin - Starred In This Movie:*the gold rush
TV/Movies: 1925 - Gary Cooper - Starred In This Movie:*tricks
TV/Movies: 1932 - Gary Cooper - Starred In This Movie:*make me a star
TV/Movies: 1932 - Katharine Hepburn - Starred In This Movie:*a bill of divorcement
TV/Movies: 1933 - Katharine Hepburn - Starred In This Movie:*little women
TV/Movies: 1934 - Lucille Ball - Starred In This Movie:*hold that girl
TV/Movies: 1934 - Lucille Ball - Starred In This Movie:*jealousy
TV/Movies: 1934 7th Academy Awards: Best Actress In A Leading Role Was won by Claudette Colbert For The Movie:*it happened One Night
TV/Movies: 1935 - Lucille Ball - Starred In This Movie:*the whole town's talking
TV/Movies: 1936 - Gary Cooper - Starred In This Movie:*the plainsman
TV/Movies: 1937 - Lucille Ball - Starred In This Movie:*stage door
TV/Movies: 1938 - Gary Cooper - Starred In This Movie:*adventures of marco polo
TV/Movies: 1938 - Gary Cooper - Starred In This Movie:*bluebeard's eighth wife
TV/Movies: 1938 - Katharine Hepburn - Starred In This Movie:*holiday
TV/Movies: 1939 - Lucille Ball - Starred In This Movie:*five came back
TV/Movies: 1939 - Lucille Ball - Starred In This Movie:*panama lady
TV/Movies: 1939 - Lucille Ball - Starred In This Movie:*twelve crowded hours
TV/Movies: 1940 - Judy Garland - Starred In This Movie:*little nellie kelly
TV/Movies: 1941 - Judy Garland - Starred In This Movie:*life begins for andy hardy
TV/Movies: 1941 14th Academy Awards: Best Actress In A Leading Role Was won by Joan Fontaine For The Movie:*suspicion
TV/Movies: 1942 - Ingrid Bergman - Starred In This Movie:*casablanca
TV/Movies: 1942 - Katharine Hepburn - Starred In This Movie:*keeper of the flame
TV/Movies: 1942 15th Academy Awards: Best Actress In A Leading Role Was Won By Greer Garson For The Movie:*mrs. miniver
TV/Movies: 1944 17th Academy Awards: Best Actress In A Leading Role Was won by Ingrid Bergman For The Movie:*gaslight
TV/Movies: 1945 - Ingrid Bergman - Starred In This Movie:*the bells of st. mary's
TV/Movies: 1946 - Lucille Ball - Starred In This Movie:*lover come back
TV/Movies: 1947 - Katharine Hepburn - Starred In This Movie:*song of love
TV/Movies: 1948 - Judy Garland - Starred In This Movie:*the pirate
TV/Movies: 1948 21st Academy Awards: Best Actress In A Leading Role Was won by Jane Wyman For The Movie:*johnny belinda
TV/Movies: 1949 - Angela Lansbury - Starred In This Movie:*the red danube
TV/Movies: 1949 - Lucille Ball - Starred In This Movie:*sorrowful jones
TV/Movies: 1950 - Elizabeth Taylor - Starred In This Movie:*the big hangover
TV/Movies: 1950 - Gary Cooper - Starred In This Movie:*dallas
TV/Movies: 1950 - Lucille Ball - Starred In This Movie:*the fuller brush girl
TV/Movies: 1953 - Marilyn Monroe - Starred In This Movie:*gentlemen prefer blondes
TV/Movies: 1954 - Marilyn Monroe - Starred In This Movie:*there's no business like show business
TV/Movies: 1954 27th Academy Awards: Best Actress In A Leading Role Was won by Grace Kelly For The Movie:*the country girl
TV/Movies: 1955 - Angela Lansbury - Starred In This Movie:*please murder me
TV/Movies: 1955 - Shirley MacLaine - Starred In This Movie:*artists and models
TV/Movies: 1956 - Elizabeth Taylor - Starred In This Movie:*giant
TV/Movies: 1956 - Ingrid Bergman - Starred In This Movie:*elena et les hommes
TV/Movies: 1957 - Elizabeth Taylor - Starred In This Movie:*raintree county
TV/Movies: 1958 - Angela Lansbury - Starred In This Movie:*the reluctant debutante
TV/Movies: 1958 - Charlie Chaplin - Starred In This Movie:*the chaplin revue
TV/Movies: 1958 - Shirley MacLaine - Starred In This Movie:*hot spell
TV/Movies: 1958 - Shirley MacLaine - Starred In This Movie:*the sheepman
TV/Movies: 1960 - Angela Lansbury - Starred In This Movie:*a breath of scandal
TV/Movies: 1960 - Elizabeth Taylor - Starred In This Movie:*scent of mystery
TV/Movies: 1960 33rd Academy Awards: Best Actress In A Leading Role Was Won By Elizabeth Taylor For The Movie:*butterfield 8
TV/Movies: 1961 - Ingrid Bergman - Starred In This Movie:*goodbye again
TV/Movies: 1962 - Angela Lansbury - Starred In This Movie:*all fall down
TV/Movies: 1962 - Elvis Presley - Starred In This Movie:*girls! girls! girls!
TV/Movies: 1962 - Elvis Presley - Starred In This Movie:*kid galahad
TV/Movies: 1963 - Jane Fonda - Starred In This Movie:*in the cool of the day
TV/Movies: 1964 - Audrey Hepburn - Starred In This Movie:*my fair lady
TV/Movies: 1964 - Jane Fonda - Starred In This Movie:*sunday in new york
TV/Movies: 1967 - Anthony Hopkins - Starred In This Movie:*the white bus
TV/Movies: 1967 - Audrey Hepburn - Starred In This Movie:*wait until dark
TV/Movies: 1967 - Elvis Presley - Starred In This Movie:*double trouble
TV/Movies: 1967 - Julie Andrews - Starred In This Movie:*thoroughly modern millie
TV/Movies: 1968 - Katharine Hepburn - Starred In This Movie:*the lion in winter
TV/Movies: 1970 43rd Academy Awards: Best Actress In A Leading Role Was Won By Glenda Jackson For The Movie:*women in Love
TV/Movies: 1971 - Sally Field - Starred In This Movie:*hitched
TV/Movies: 1971 - Sally Field - Starred In This Movie:*maybe i'll come home in the spring
TV/Movies: 1975 48th Academy Awards: Best Actress In A Leading Role Was Won By Louise Fletcher For The Movie:*one flew over The Cuckoo's Nest
TV/Movies: 1976 - Jodie Foster - Starred In This Movie:*freaky friday
TV/Movies: 1976 - Sally Field - Starred In This Movie:*sybil
TV/Movies: 1976 - Woody Allen - Starred In This Movie:*the front
TV/Movies: 1976 49th Academy Awards: Best Actress In A Leading Role Was won by Faye Dunaway For The Movie:*network
TV/Movies: 1977 - Anthony Hopkins - Starred In This Movie:*a bridge too far
TV/Movies: 1977 - Arnold Schwarzenegger - Starred In This Movie:*pumping iron
TV/Movies: 1977 - Jane Fonda - Starred In This Movie:*julia
TV/Movies: 1977 - Woody Allen - Starred In This Movie:*annie hall
TV/Movies: 1978 - Anthony Hopkins - Starred In This Movie:*magic
TV/Movies: 1979 - Audrey Hepburn - Starred In This Movie:*bloodline
TV/Movies: 1979 - Dustin Hoffman - Starred In This Movie:*agatha
TV/Movies: 1979 - Lauren Bacall - Starred In This Movie:*h.e.a.l.t.h
TV/Movies: 1979 52nd Academy Awards: Best Actress In A Leading Role Was won by Sally Field For The Movie:*norma rae
TV/Movies: 1980 - Anthony Hopkins - Starred In This Movie:*the elephant man
TV/Movies: 1980 - Jodie Foster - Starred In This Movie:*carny
TV/Movies: 1980 - Jodie Foster - Starred In This Movie:*foxes
TV/Movies: 1980 - Michelle Pfeiffer - Starred In This Movie:*the hollywood knights
TV/Movies: 1981 - Audrey Hepburn - Starred In This Movie:*they all laughed
TV/Movies: 1981 - Demi Moore - Starred In This Movie:*choices
TV/Movies: 1981 - Kirstie Alley - Starred In This Movie:*one more chance
TV/Movies: 1981 - Sigourney Weaver - Starred In This Movie:*eyewitness
TV/Movies: 1982 - Demi Moore - Starred In This Movie:*young doctors in love
TV/Movies: 1982 - Sally Field - Starred In This Movie:*kiss me goodbye
TV/Movies: 1982 - Woody Allen - Starred In This Movie:*a midsummer night's sex comedy
TV/Movies: 1983 - Kurt Russell - Starred in this movie:*silkwood
TV/Movies: 1983 - Robin Williams - Starred In This Movie:*the survivors
TV/Movies: 1984 - Arnold Schwarzenegger - Starred In This Movie:*the terminator
TV/Movies: 1984 - Shirley MacLaine - Starred In This Movie:*cannonball run ii
TV/Movies: 1984 57th Academy Awards: Best Actress In A Leading Role Was won by Sally Field For The Movie:*places in the heart
TV/Movies: 1985 - Anthony Hopkins - Starred In This Movie:*arch of triumph
TV/Movies: 1986 - Drew Barrymore - Starred In This Movie:*babes in toyland
TV/Movies: 1987 - Goldie Hawn - Starred in this movie:*overboard
TV/Movies: 1987 - Jodie Foster - Starred In This Movie:*siesta
TV/Movies: 1987 - Meg Ryan - Starred In This Movie:*innerspace
TV/Movies: 1988 - Arnold Schwarzenegger - Starred In This Movie:*twins
TV/Movies: 1988 - Hugh Grant - Starred In This Movie:*the dawning
TV/Movies: 1988 - Michelle Pfeiffer - Starred In This Movie:*dangerous liaisons
TV/Movies: 1988 - Michelle Pfeiffer - Starred In This Movie:*tequila sunrise
TV/Movies: 1988 - Sharon Stone - Starred In This Movie:*action jackson
TV/Movies: 1989 - Bruce Willis - Starred In This Movie:*in country
TV/Movies: 1989 - Dustin Hoffman - Starred In This Movie:*family business
TV/Movies: 1989 - Lauren Bacall - Starred In This Movie:*dinner at eight
TV/Movies: 1989 - Shirley MacLaine - Starred In This Movie:*steel magnolias
TV/Movies: 1989 62nd Academy Awards: Best Actress In A Leading Role Was Won By Jessica Tandy For The Movie:*driving miss Daisy
TV/Movies: 1990 - Bruce Willis - Starred In This Movie:*look who's talking too
TV/Movies: 1990 - Bruce Willis - Starred In This Movie:*the bonfire of the vanities
TV/Movies: 1990 - Shirley MacLaine - Starred In This Movie:*waiting for the light
TV/Movies: 1991 - Anthony Hopkins - Starred In This Movie:*the silence of the lambs
TV/Movies: 1991 - Demi Moore - Starred In This Movie:*nothing but trouble
TV/Movies: 1991 - Julia Roberts - Starred In This Movie:*hook
TV/Movies: 1991 - Julia Roberts - Starred In This Movie:*sleeping with the enemy
TV/Movies: 1991 - Sharon Stone - Starred In This Movie:*scissors
TV/Movies: 1991 - Shirley MacLaine - Starred In This Movie:*defending your life
TV/Movies: 1991 64th Academy Awards: Best Actress In A Leading Role Was Won By Jodie Foster For The Movie:*the silence of The Lambs
TV/Movies: 1992 - Whoopi Goldberg - Starred In This Movie:*sister act
TV/Movies: 1993 - Anthony Hopkins - Starred In This Movie:*the innocent
TV/Movies: 1993 - Bruce Willis - Starred In This Movie:*striking distance
TV/Movies: 1993 - Whoopi Goldberg - Starred In This Movie:*made in america
TV/Movies: 1993: name one of the major stars in the film house of cards*tommy lee jones
TV/Movies: 1994 - Anthony Hopkins - Starred In This Movie:*the road to wellville
TV/Movies: 1994 - Harrison Ford - Starred In This Movie:*a century of cinema
TV/Movies: 1994 - Hugh Grant - Starred In This Movie:*four weddings and a funeral
TV/Movies: 1994 - Sharon Stone - Starred In This Movie:*intersection
TV/Movies: 1995 - Brad Pitt - Starred In This Movie:*12 monkeys
TV/Movies: 1995 - Bruce Willis - Starred In This Movie:*12 monkeys
TV/Movies: 1995 - Emma Thompson - Starred In This Movie:*sense and sensibility
TV/Movies: 1995 - Julia Roberts - Starred In This Movie:*something to talk about
TV/Movies: 1995 - Sharon Stone - Starred In This Movie:*the quick and the dead
TV/Movies: 1995 68th Academy Awards: Best Actress In A Leading Role Was won by Susan Sarandon For The Movie:*dead man Walking
TV/Movies: 1996 - Cameron Diaz - Starred In This Movie:*the last supper
TV/Movies: 1996 - Demi Moore - Starred In This Movie:*striptease
TV/Movies: 1996 - Lauren Holly - Starred In This Movie:*beautiful girls
TV/Movies: 1996 - Meryl Streep - Starred In This Movie:*before and after
TV/Movies: 1996 - Neve Campbell - Starred In This Movie:*the craft
TV/Movies: 1996 - Robin Williams - Starred In This Movie:*the secret agent
TV/Movies: 1996 - Sally Field - Starred In This Movie:*eye for an eye
TV/Movies: 1996 - Sandra Bullock - Starred In This Movie:*a time to kill
TV/Movies: 1996 69th Academy Awards: Best Actress In A Leading Role Was won by Frances Mcdormand For The Movie:*fargo
TV/Movies: 1996-1997 Movies: A family-values politician targets Dangerfield's vulgarly popular talk show*meet wally sparks
TV/Movies: 1996-1997 Movies: Jim Carrey makes himself Matthew Broderick's friend in this dark comedy*the cable guy
TV/Movies: 1996-1997 Movies: The biography of the Latin-American Tejano music star slain by a fan in 1995*selena
TV/Movies: 1997 - Alicia Silverstone - Starred In This Movie:*batman & robin
TV/Movies: 1997 - Julia Roberts - Starred In This Movie:*conspiracy theory
TV/Movies: 1997 - Shirley MacLaine - Starred In This Movie:*a smile like yours
TV/Movies: 1997, This Movie was Released on April 18 ----------*traveller
TV/Movies: 1997, This Movie was Released on February 28 Donnie ----------*brasco
TV/Movies: 1997, This Movie was Released on January 17 ----------*metro
TV/Movies: 1997, This Movie was Released on January 31 Star Wars: ----------*special edition
TV/Movies: 1997, This Movie was Released on June 13 Speed 2: ----------*cruise control
TV/Movies: 1997, This Movie was Released on June 6 ----------*bliss
TV/Movies: 1997, This Movie was Released on March 28 The 6th ----------*man
TV/Movies: 1997, This Movie was Released on March 7 Jungle 2 ----------*jungle
TV/Movies: 1997, This Movie was Released on May 2 Austin Powers: International ----------*man of mystery
TV/Movies: 1997, This Movie was Released on May 9 Father's ----------*day
TV/Movies: 1998 - Drew Barrymore - Starred In This Movie:*the wedding singer
TV/Movies: 1998 - Jamie Lee Curtis - Starred In This Movie:*halloween: h20
TV/Movies: 1998 - Meg Ryan - Starred In This Movie:*hurlyburly
TV/Movies: 1998 - Sharon Stone - Starred In This Movie:*the mighty
TV/Movies: 1998, This Movie was Released on August 21 Dance ----------*with me
TV/Movies: 1998, This Movie was Released on August 21 Wrongfully ----------*accused
TV/Movies: 1998, This Movie was Released on December 11 Star Trek: ----------*insurrection
TV/Movies: 1998, This Movie was Released on February 6 The Replacement ----------*killers
TV/Movies: 1998, This Movie was Released on July 15 There's something ----------*about mary
TV/Movies: 1998, This Movie was Released on July 31 ----------*baseketball
TV/Movies: 1998, This Movie was Released on June 12 Six Days, ----------*seven nights
TV/Movies: 1998, This Movie was Released on June 26 Gone With ----------*the wind
TV/Movies: 1998, This Movie was Released on June 26 Hav ----------*plenty
TV/Movies: 1998, This Movie was Released on March 27 Wide ----------*awake
TV/Movies: 1998, This Movie was Released on May 1 Dancer, ----------*texas pop. 81
TV/Movies: 1998, This Movie was Released on May 1 Les ----------*miserables
TV/Movies: 1998, This Movie was Released on May 8 Deep ----------*impact
TV/Movies: 1998, This Movie was Released on October 2 Dee Snider's ----------*strangeland
TV/Movies: 1998, This Movie was Released on October 30 Living Out ----------*loud
TV/Movies: 1998, This Movie was Released on September 11 ----------*rounders
TV/Movies: 1998, This Movie was Released on September 11 Without ----------*limits
TV/Movies: 1999 - Drew Barrymore - Starred In This Movie:*never been kissed
TV/Movies: 1999 - Sigourney Weaver - Starred In This Movie:*a map of the world
TV/Movies: 1999, This Movie was Released on August 13 ----------*bowfinger
TV/Movies: 1999, This Movie was Released on August 27 The ----------*muse
TV/Movies: 1999, This Movie was Released on August 6 Mystery ----------*men
TV/Movies: 1999, This Movie was Released on December 22 Man on ----------*the moon
TV/Movies: 1999, This Movie was Released on February 19 Office ----------*space
TV/Movies: 1999, This Movie was Released on June 4 Desert ----------*blue
TV/Movies: 1999, This Movie was Released on March 12 Wing ----------*commander
TV/Movies: 1999, This Movie was Released on May 14 A Midsummer ----------*night's dream
TV/Movies: 1999, This Movie was Released on May 28 Notting ----------*hill
TV/Movies: 1999, This Movie was Released on October 1 American ----------*beauty
TV/Movies: 1999, This Movie was Released on October 15 The Straight ----------*story
TV/Movies: 1999, This Movie was Released on October 8 Random ----------*hearts
TV/Movies: 1999, This Movie was Released on September 24 ----------*guinevere
TV/Movies: 2000, This Movie was Released on April 14 Keeping The ----------*faith
TV/Movies: 2000, This Movie was Released on April 28 The Virgin ----------*suicides
TV/Movies: 2000, This Movie was Released on August 18 The ----------*cell
TV/Movies: 2000, This Movie was Released on August 25 Bring ----------*it on
TV/Movies: 2000, This Movie was Released on December 22 Cast ----------*away
TV/Movies: 2000, This Movie was Released on February 18 The Whole ----------*nine yards
TV/Movies: 2000, This Movie was Released on January 14 Girl, ----------*interrupted
TV/Movies: 2000, This Movie was Released on July 21 ----------*loser
TV/Movies: 2000, This Movie was Released on July 26 Thomas and the ----------*magic railroad
TV/Movies: 2000, This Movie was Released on June 2 Big Momma's ----------*house
TV/Movies: 2000, This Movie was Released on June 9 Gone In Sixty ----------*seconds
TV/Movies: 2000, This Movie was Released on March 10 The Ninth ----------*gate
TV/Movies: 2000, This Movie was Released on March 17 Erin ----------*brockovich
TV/Movies: 2000, This Movie was Released on May 5 I Dreamed of ----------*africa
TV/Movies: 2000, This Movie was Released on November 3 The Legend of ----------*bagger vance
TV/Movies: 2000, This Movie was Released on October 20 ----------*bedazzled
TV/Movies: 2000, This Movie was Released on October 20 The ----------*yards
TV/Movies: 2000, This Movie was Released on October 6 ----------*girlfight
TV/Movies: 2000, This Movie was Released on October 6 Digimon: ----------*the movie
TV/Movies: 2000, This Movie was Released on October 6 Get ----------*carter
TV/Movies: 2000, This Movie was Released on September 1 ----------*whipped
TV/Movies: 2000, This Movie was Released on September 22 Almost ----------*famous
TV/Movies: 2000, This Movie was Released on September 29 ----------*beautiful
TV/Movies: 2000, This Movie was Released on September 29 Best In ----------*show
TV/Movies: 2000, This Movie was Released on September 29 The ----------*exorcist
TV/Movies: 50s Flicks: Which French star appeared in ----------and God Created Woman----------, set in St. Tropez*brigitte bardot
TV/Movies: Academy awards: best actor, clark gable, & best actress, claudette colbert, won for this film, which was best picture*retinol
TV/Movies: Actor originally intended to be Wizard in "Wizard of Oz"*w c fields
TV/Movies: Actress In The Role: Batman ---> Vicki Vale*kim basinger
TV/Movies: Amazing actress who won for ----------The Prime of Miss Jean Brodie----------*maggie smith
TV/Movies: Among animation aficionados, what is generally considered to be Chuck Jones best animated short film*one froggy evening
TV/Movies: B Movies: Bomb that featured Eddie Murphy and Dudley Moore as arms dealers*best defense
TV/Movies: B Movies: Disney's 1979 attempt at a sci-fi flick, with Ernest Borgnine*the black hole
TV/Movies: B Movies: John Singleton hit it big in 1991 with this ghetto tale (spelling...)*boyz n the hood
TV/Movies: Back To The Future: Back to the Future 3 was set in one of Doc's favourite places, the ---------- ----------*wild west
TV/Movies: Bill & ted's excellent adventure: strange things are amuck at the ----------*circle k
TV/Movies: Blade Runner: The phrase "Blade Runner" comes from a book by this author*william s burroughs
TV/Movies: Born April 16, 1889, He Starred In This Movie: Souls For Sale - 1923*charlie chaplin
TV/Movies: Born April 16, 1889, He Starred In This Movie: The Chaplin Revue - 1958*charlie chaplin
TV/Movies: Born April 29, 1958, She starred in this movie: Falling in Love Again - 1980*michelle pfeiffer
TV/Movies: Born April 29, 1958, She starred in this movie: Grease 2 - 1982*michelle pfeiffer
TV/Movies: Born April 29, 1958, She starred in this movie: Ladyhawke - 1985*michelle pfeiffer
TV/Movies: Born April 29, 1958, She starred in this movie: The Age of Innocence - 1993*michelle pfeiffer
TV/Movies: Born April 29, 1958, She starred in this movie: Up Close and Personal - 1996*michelle pfeiffer
TV/Movies: Born Aug 17, 1943, He starred in this movie: Mary Shelley's Frankenstein - 1994*robert de niro
TV/Movies: Born Aug 17, 1943, He starred in this movie: New York, New York - 1977*robert de niro
TV/Movies: Born Aug 17, 1943, He starred in this movie: True Confessions - 1981*robert de niro
TV/Movies: Born Aug 25, 1930, He starred in this movie: Action of the Tiger - 1957*sean connery
TV/Movies: Born Aug 25, 1930, He starred in this movie: Entrapment - 1999*sean connery
TV/Movies: Born Aug 25, 1930, He starred in this movie: Family Business - 1989*sean connery
TV/Movies: Born Aug 25, 1930, He starred in this movie: Five Days One Summer - 1982*sean connery
TV/Movies: Born Aug 25, 1930, He starred in this movie: Goldfinger - 1964*sean connery
TV/Movies: Born Aug 25, 1930, He starred in this movie: Hell Drivers - 1958*sean connery
TV/Movies: Born Aug 25, 1930, He starred in this movie: Highlander - 1986*sean connery
TV/Movies: Born Aug 25, 1930, He starred in this movie: Marnie - 1964*sean connery
TV/Movies: Born Aug 25, 1930, He starred in this movie: Meteor - 1979*sean connery
TV/Movies: Born Aug 25, 1930, He starred in this movie: Playing by Heart - 1998*sean connery
TV/Movies: Born Aug 25, 1930, He starred in this movie: The Rock - 1996*sean connery
TV/Movies: Born Aug 29, 1915, She starred in this movie: Autumn Sonata - 1978*ingrid bergman
TV/Movies: Born Aug 29, 1915, She starred in this movie: Hostsonaten - 1978*ingrid bergman
TV/Movies: Born Aug 29, 1915, She starred in this movie: Rich Girl - 1991*ingrid bergman
TV/Movies: Born Aug 31, 1949, He starred in this movie: Bloodbrothers - 1978*richard gere
TV/Movies: Born Aug 31, 1949, He starred in this movie: Primal Fear - 1996*richard gere
TV/Movies: Born Aug 31, 1949, He starred in this movie: Rhapsody in August - 1991*richard gere
TV/Movies: Born Aug 31, 1949, He starred in this movie: Runaway Bride - 1999*richard gere
TV/Movies: Born Aug 31, 1949, He starred in this movie: Unzipped - 1995*richard gere
TV/Movies: Born Aug 6, 1911, She starred in this movie: Bunker Bean - 1936*lucille ball
TV/Movies: Born Aug 6, 1911, She starred in this movie: Fugitive Lady - 1934*lucille ball
TV/Movies: Born Aug 6, 1911, She starred in this movie: Moulin Rouge - 1934*lucille ball
TV/Movies: Born Aug 8, 1937, He starred in this movie: Death of a Salesman - 1985*dustin hoffman
TV/Movies: Born Aug 8, 1937, He starred in this movie: Dick Tracy - 1990*dustin hoffman
TV/Movies: Born Aug 8, 1937, He starred in this movie: Lenny - 1974*dustin hoffman
TV/Movies: Born Aug 8, 1937, He starred in this movie: Outbreak - 1995*dustin hoffman
TV/Movies: Born Dec 1, 1945, She starred in this movie: Down and Out in Beverly Hills - 1986*bette midler
TV/Movies: Born Dec 1, 1945, She starred in this movie: Get Shorty - 1995*bette midler
TV/Movies: Born Dec 1, 1945, She starred in this movie: Stella - 1990*bette midler
TV/Movies: Born Dec 18 1963, He starred in this movie: 12 Monkeys - 1995*brad pitt
TV/Movies: Born Dec 18 1963, He starred in this movie: Fight Club - 1999*brad pitt
TV/Movies: Born Dec 18 1963, He starred in this movie: Too Young To Die? - 1990*brad pitt
TV/Movies: Born Dec 31, 1937, He starred in this movie: 84 Charing Cross Road - 1987*anthony hopkins
TV/Movies: Born Dec 31, 1937, He starred in this movie: Bram Stoker's Dracula - 1992*anthony hopkins
TV/Movies: Born Dec 31, 1937, He starred in this movie: Chaplin - 1992*anthony hopkins
TV/Movies: Born Dec 31, 1937, He starred in this movie: The Good Father - 1986*anthony hopkins
TV/Movies: Born Dec 31, 1937, He starred in this movie: The Lion in Winter - 1968*anthony hopkins
TV/Movies: Born Dec 31, 1937, He starred in this movie: The Tenth Man - 1988*anthony hopkins
TV/Movies: Born Dec 31, 1937, He starred in this movie: The Trial - 1992*anthony hopkins
TV/Movies: Born Dec 31, 1937, He starred in this movie: Titus - 1999*anthony hopkins
TV/Movies: Born Feb 22, 1975, She starred in this movie: Far From Home - 1988*drew barrymore
TV/Movies: Born Feb 22, 1975, She starred in this movie: Scream - 1996*drew barrymore
TV/Movies: Born Feb 22, 1975, She starred in this movie: See You in the Morning - 1989*drew barrymore
TV/Movies: Born Feb 27, 1932, She starred in this movie: Ash Wednesday - 1973*elizabeth taylor
TV/Movies: Born Feb 27, 1932, She starred in this movie: Cleopatra - 1963*elizabeth taylor
TV/Movies: Born Feb 27, 1932, She starred in this movie: Conspirator - 1949*elizabeth taylor
TV/Movies: Born Feb 27, 1932, She starred in this movie: Hammersmith Is Out - 1972*elizabeth taylor
TV/Movies: Born Feb 27, 1932, She starred in this movie: Poker Alice - 1987*elizabeth taylor
TV/Movies: Born Jan 12, 1955, She starred in this movie: It Takes Two - 1995*kirstie alley
TV/Movies: Born Jan 18, 1955, He starred in this movie: For Love of the Game - 1999*kevin costner
TV/Movies: Born Jan 18, 1955, He starred in this movie: Testament - 1983*kevin costner
TV/Movies: Born Jan 3, 1956, He starred in this movie: Bird on a Wire - 1990*mel gibson
TV/Movies: Born Jan 3, 1956, He starred in this movie: Gallipoli - 1981*mel gibson
TV/Movies: Born Jan 3, 1956, He starred in this movie: Mad Max: Beyond Thunderdome - 1985*mel gibson
TV/Movies: Born Jan 3, 1956, He starred in this movie: Mrs. Soffel - 1984*mel gibson
TV/Movies: Born Jan 3, 1956, He starred in this movie: The Road Warrior - 1982*mel gibson
TV/Movies: Born Jan 8, 1935, He starred in this movie: Flaming Star - 1960*elvis presley
TV/Movies: Born Jan 8, 1935, He starred in this movie: Frankie and Johnny - 1966*elvis presley
TV/Movies: Born Jul 21, 1952, He starred in this movie: Hook - 1991*robin williams
TV/Movies: Born Jul 21, 1952, He starred in this movie: Jumanji - 1995*robin williams
TV/Movies: Born Jul 21, 1952, He starred in this movie: Nine Months - 1995*robin williams
TV/Movies: Born Jul 6, 1946, He starred in this movie: Cobra - 1986*sylvester stallone
TV/Movies: Born Jul 6, 1946, He starred in this movie: F.I.S.T. - 1978*sylvester stallone
TV/Movies: Born July 13, 1942, He starred in this movie: Indiana Jones and the Temple of Doom - 1984*harrison ford
TV/Movies: Born July 26, 1964, She starred in this movie: In Love and War - 1996*sandra bullock
TV/Movies: Born July 3, 1962, He starred in this movie: Jerry Maguire - 1996*tom cruise
TV/Movies: Born July 3, 1962, He starred in this movie: Risky Business - 1983*tom cruise
TV/Movies: Born July 3, 1962, He starred in this movie: The Firm - 1993*tom cruise
TV/Movies: Born July 30, 1947, He starred in this movie: A Century of Cinema - 1994*arnold schwarzenegger
TV/Movies: Born July 30, 1947, He starred in this movie: Conan the Destroyer - 1984*arnold schwarzenegger
TV/Movies: Born July 30, 1947, He starred in this movie: Stay Hungry - 1976*arnold schwarzenegger
TV/Movies: Born July 9, 1956, He starred in this movie: A League of Their Own - 1992*tom hanks
TV/Movies: Born July 9, 1956, He starred in this movie: Every Time We Say Goodbye - 1986*tom hanks
TV/Movies: Born July 9, 1956, He starred in this movie: Forrest Gump - 1994*tom hanks
TV/Movies: Born July 9, 1956, He starred in this movie: Nothing in Common - 1986*tom hanks
TV/Movies: Born July 9, 1956, He starred in this movie: The Green Mile - 1999*tom hanks
TV/Movies: Born Jun 22, 1949, She starred in this movie: Julia - 1977*meryl streep
TV/Movies: Born Jun 22, 1949, She starred in this movie: Marvin's Room - 1996*meryl streep
TV/Movies: Born Jun 22, 1949, She starred in this movie: The Deer Hunter - 1978*meryl streep
TV/Movies: Born June 1, 1926, She starred in this movie: Gentlemen Prefer Blondes - 1953*marilyn monroe
TV/Movies: Born June 1, 1926, She starred in this movie: Monkey Business - 1952*marilyn monroe
TV/Movies: Born June 10, 1922, She starred in this movie: Girl Crazy - 1943*judy garland
TV/Movies: Born Mar 10,1958, She starred in this movie: Bolero - 1982*sharon stone
TV/Movies: Born Mar 10,1958, She starred in this movie: Deadly Blessing - 1981*sharon stone
TV/Movies: Born Mar 10,1958, She starred in this movie: Diary of a Hitman - 1992*sharon stone
TV/Movies: Born Mar 10,1958, She starred in this movie: Last Action Hero - 1993*sharon stone
TV/Movies: Born Mar 10,1958, She starred in this movie: Where Sleeping Dogs Lie - 1993*sharon stone
TV/Movies: Born Mar 14, 1947, He starred in this movie: Deconstructing Harry - 1997*billy crystal
TV/Movies: Born Mar 14, 1947, He starred in this movie: Fathers' Day - 1997*billy crystal
TV/Movies: Born Mar 17, 1951, He starred in this movie: Charley and the Angel - 1973*kurt russell
TV/Movies: Born Mar 17, 1951, He starred in this movie: Swing Shift - 1984*kurt russell
TV/Movies: Born Mar 19, 1955, He starred in this movie: 12 Monkeys - 1995*bruce willis
TV/Movies: Born Mar 19, 1955, He starred in this movie: Look Who's Talking Too - 1990*bruce willis
TV/Movies: Born Mar 19, 1955, He starred in this movie: Mercury Rising - 1998*bruce willis
TV/Movies: Born May 4, 1929, She starred in this movie: Always - 1989*audrey hepburn
TV/Movies: Born May 4, 1929, She starred in this movie: How To Steal a Million - 1966*audrey hepburn
TV/Movies: Born May 4, 1929, She starred in this movie: Nous irons a Monte Carlo - 1951*audrey hepburn
TV/Movies: Born May 4, 1929, She starred in this movie: One Wild Oat - 1951*audrey hepburn
TV/Movies: Born May 4, 1929, She starred in this movie: Roman Holiday - 1953*audrey hepburn
TV/Movies: Born May 4, 1929, She starred in this movie: The Lavender Hill Mob - 1951*audrey hepburn
TV/Movies: Born May 4, 1929, She starred in this movie: They All Laughed - 1981*audrey hepburn
TV/Movies: Born May 7, 1901, He Starred In This Movie: Alice in Wonderland - 1933*gary cooper
TV/Movies: Born May 7, 1901, He Starred In This Movie: Dick Turpin - 1925*gary cooper
TV/Movies: Born May 7, 1901, He Starred In This Movie: For Whom the Bell Tolls - 1943*gary cooper
TV/Movies: Born May 7, 1901, He Starred In This Movie: Lest We Forget - 1937*gary cooper
TV/Movies: Born May 7, 1901, He Starred In This Movie: Mr. Deeds Goes to Town - 1936*gary cooper
TV/Movies: Born May 7, 1901, He Starred In This Movie: Nevada - 1944*gary cooper
TV/Movies: Born May 7, 1901, He Starred In This Movie: Task Force - 1949*gary cooper
TV/Movies: Born May 7, 1901, He Starred In This Movie: The Pride of the Yankees - 1942*gary cooper
TV/Movies: Born May 7, 1901, He Starred In This Movie: The Virginian - 1929*gary cooper
TV/Movies: Born May 7, 1901, He Starred In This Movie: The Westerner - 1940*gary cooper
TV/Movies: Born May 7, 1901, He Starred In This Movie: Tricks - 1925*gary cooper
TV/Movies: Born May 7, 1901, He Starred In This Movie: Wolf Song - 1929*gary cooper
TV/Movies: Born May 7, 1901, He Starred In This Movie: You're in the Navy Now - 1951*gary cooper
TV/Movies: Born Nov 11, 1962, She starred in this movie: No Small Affair - 1985*demi moore
TV/Movies: Born Nov 19, 1961, She starred in this movie: A Century of Cinema - 1994*meg ryan
TV/Movies: Born Nov 19, 1961, She starred in this movie: Armed and Dangerous - 1986*meg ryan
TV/Movies: Born Nov 19, 1961, She starred in this movie: Courage Under Fire - 1996*meg ryan
TV/Movies: Born Nov 19, 1961, She starred in this movie: French Kiss - 1995*meg ryan
TV/Movies: Born Nov 19, 1961, She starred in this movie: Hanging Up - 2000*meg ryan
TV/Movies: Born Nov 19, 1961, She starred in this movie: The Presidio - 1988*meg ryan
TV/Movies: Born Nov 19, 1962, She starred in this movie: Hotel New Hampshire - 1984*jodie foster
TV/Movies: Born Nov 19, 1962, She starred in this movie: O'Hara's Wife - 1982*jodie foster
TV/Movies: Born Nov 19, 1962, She starred in this movie: One Little Indian - 1973*jodie foster
TV/Movies: Born Nov 19, 1962, She starred in this movie: Shadows and Fog - 1992*jodie foster
TV/Movies: Born Nov 19, 1962, She starred in this movie: The Accused - 1988*jodie foster
TV/Movies: Born Nov 19, 1962, She starred in this movie: The Hotel New Hampshire - 1984*jodie foster
TV/Movies: Born Nov 21, 1945, She starred in this movie: Cactus Flower - 1969*goldie hawn
TV/Movies: Born Nov 21, 1945, She starred in this movie: Deceived - 1991*goldie hawn
TV/Movies: Born Nov 21, 1945, She starred in this movie: The Girl From Petrovka - 1974*goldie hawn
TV/Movies: Born Oct 1, 1935, She starred in this movie: Little Miss Marker - 1980*julie andrews
TV/Movies: Born Oct 28, 1967, She starred in this movie: Mystic Pizza - 1988*julia roberts
TV/Movies: Born Oct 28, 1967, She starred in this movie: Runaway Bride - 1999*julia roberts
TV/Movies: Born Oct 29, 1971, She starred in this movie: Square Dance - 1987*winona ryder
TV/Movies: Born Oct 29, 1971, She starred in this movie: The Age of Innocence - 1993*winona ryder
TV/Movies: Born Oct 3, 1973, She starred in this movie: Scream 3 - 2000*neve campbell
TV/Movies: Born Oct 3, 1973, She starred in this movie: Three to Tango - 1999*neve campbell
TV/Movies: Born Oct 4, 1946, She starred in this movie: He'll See You Now - 1984*susan sarandon
TV/Movies: Born Oct 4, 1946, She starred in this movie: Mortadella/Lady Liberty - 1972*susan sarandon
TV/Movies: Born Oct 4, 1946, She starred in this movie: The Player - 1992*susan sarandon
TV/Movies: Born Oct 4, 1946, She starred in this movie: Women of Valor - 1986*susan sarandon
TV/Movies: Born Oct 8, 1949, She starred in this movie: 1492: Conquest of Paradise - 1992*sigourney weaver
TV/Movies: Born Oct 8, 1949, She starred in this movie: Eyewitness - 1981*sigourney weaver
TV/Movies: Born Oct 8, 1949, She starred in this movie: Jeffrey - 1995*sigourney weaver
TV/Movies: Born Oct 8, 1949, She starred in this movie: One Woman or Two - 1987*sigourney weaver
TV/Movies: Born Sep 15, 1946, He starred in this movie: Jackson County Jail - 1976*tommy lee jones
TV/Movies: Born Sep 15, 1946, He starred in this movie: The Betsy - 1978*tommy lee jones
TV/Movies: Born Sep 15, 1946, He starred in this movie: The Eyes of Laura Mars - 1978*tommy lee jones
TV/Movies: Born Sep 28, 1934, She starred in this movie: Le Fils de Caroline Cherie - 1955*brigitte bardot
TV/Movies: Born Sep 28, 1934, She starred in this movie: Le Trou Normand - 1952*brigitte bardot
TV/Movies: Born Sep 28, 1934, She starred in this movie: Les Bijoutiers du Clair de Lune/The Night Heaven Fell - 1958*brigitte Bardot
TV/Movies: Born Sep 28, 1934, She starred in this movie: Two Weeks in September - 1967*brigitte bardot
TV/Movies: Born Sep 28, 1934, She starred in this movie: Une Ravissante Idiote/A Ravishing Idiot - 1964*brigitte bardot
TV/Movies: Born Sep 28, 1934, She starred in this movie: Voulez-vous danser avec moi?/Come Dance With Me! - 1959*brigitte Bardot
TV/Movies: Born Sep 5, 1940, She starred in this movie: 100 Rifles - 1969*raquel welch
TV/Movies: Born Sep 5, 1940, She starred in this movie: L'Animal/Stuntwoman - 1977*raquel welch
TV/Movies: Born Sep 5, 1940, She starred in this movie: Naked Gun 33 1/3: The Final Insult - 1994*raquel welch
TV/Movies: Born Sep 5, 1940, She starred in this movie: Scandal in a Small Town - 1988*raquel welch
TV/Movies: Born Sep 5, 1940, She starred in this movie: The Four Musketeers - 1975*raquel welch
TV/Movies: Born Sep 9, 1960, He starred in this movie: Impromptu - 1990*hugh grant
TV/Movies: Cartoons: Fred Flintstone's boss*mr. slate
TV/Movies: Cartoons: Name of the creator of Ren and Stimpy*john kricfalusi
TV/Movies: Category: "Why, I guess you don't know everything about women yet"*the empire strikes back
TV/Movies: Category: 50s Flicks: Which 1957 film had the whistled "Colonel Bogey" as its theme*the bridge on the river kwai
TV/Movies: Category: 80s Movies: 1983 film that brings Jeff Goldblum, Glenn Close, and William Hurt together*the big chill
TV/Movies: Category: 80s Movies: 1987 thriller hyped with the line "she mates and she kills"*black widow
TV/Movies: Category: 80s Movies: Featured the lines "I was born to love you. I was born to lick your face."*caddyshack
TV/Movies: Category: 80s Movies: Roman Polanski's 1980 film adaptation of a Thomas Hardy novel*tess
TV/Movies: Category: 90s TV: In which program do Mulder and Scully appear*x-files
TV/Movies: Category: 90s TV: Show about comedian-turned 'tool-man' Tim Allen and his sidekick Al*home improvement
TV/Movies: Category: Academy Award Directors: Elia Kazan was awarded in 1954 for this movie*on the waterfront
TV/Movies: Category: Academy Award Directors: John Schlesinger was awarded in 1969 for this movie*midnight cowboy
TV/Movies: Category: Academy Award Directors: Michael Curtiz was awarded in 1943 for this movie*casablanca
TV/Movies: Category: Academy Award Directors: Robert Wise was awarded in 1965 for this movie*the sound of music
TV/Movies: Category: Academy Award Directors: Vincente Minnelli was awarded in 1958 for this movie*gigi
TV/Movies: Category: Academy Awards: Who is the only person to win 3 Best Supporting Actor awards*walter brennan
TV/Movies: Category: Actors Common Ground 1: The Donna Reed Show, Pete and Gladys, M-A-S-H*harry morgan
TV/Movies: Category: Actors In Film: Dennis Quaid, Louis Gossett Jr*enemy mine
TV/Movies: Category: Actors In Film: Sean Penn, Al Pacino, John Leguzamo*carlitos way
TV/Movies: Category: Actors In Film: Tom Berenger, Sidney Poitier, Clancy Brown*shoot to kill
TV/Movies: Category: ActPersons: He was born Allen Konigsberg*woody allen
TV/Movies: Category: ActPersons: The birthplace (city) of Jack Lemmon*boston
TV/Movies: Category: ActPersons: The birthplace (city) of Raul Julia*san juan
TV/Movies: Category: Actress In The Role: The Grifters ---> Lily Dillon*anjelica huston
TV/Movies: Category: American Cartoonists: Cartoon created by C.C. Beck*captain marvel
TV/Movies: Category: American Cartoonists: Cartoon created by Tom Wilson*ziggy
TV/Movies: Category: Animaniacs: This is the argumentative "good feather"; also a popular pasta topping*pesto
TV/Movies: Category: Anime: What is the other (not Japanese) nationality of Asuka Soryuu Langley*german
TV/Movies: Category: B Movies: John Singleton hit it big in 1991 with this ghetto tale (spelling...)*boyz n the hood
TV/Movies: Category: B Movies: Steve McQueen's first starring role was in this 1958 B-Movie*the blob
TV/Movies: Category: B Movies: Sting and Jennifer Beals in this remake of The Bride of Frankenstein*the bride
TV/Movies: Category: Beverly Hills 90210: The numbers 90210 are what*zip code
TV/Movies: Category: Blazing Saddles: Who played the German Chanteuse*madeline kahn
TV/Movies: Category: Bleeding Heart Movies: "Farm boy, fetch me that pitcher."*the princess bride
TV/Movies: Category: Bleeding Heart Movies: Will Smith knifes himself to get aquantied with high society folks*six degrees of Seperation
TV/Movies: Category: Blues Brothers: Who played the jilted fiancee*carrie fisher
TV/Movies: Category: Cartoon Sidekicks: Buttons*mindy
TV/Movies: Category: Cartoon Sidekicks: Calvin and ------*hobbes
TV/Movies: Category: Cartoon Sidekicks: Disney: Dale*chip
TV/Movies: Category: Cartoons: Boris Badanov's nationality*pottsylvanian
TV/Movies: Category: Cheers Trivia: Norm's favorite restaurant*hungry heifer
TV/Movies: Category: Classic Movies: He played Johnny Boy, small-time gambler in big-time debt to loan sharks*robert deniro
TV/Movies: Category: Commitments: "Decko the bus conductor Is that ------ ----- or ------ -----"*single decko double decko
TV/Movies: Category: Commitments: In what year was this film released*1991
TV/Movies: Category: Defining Roles: Back to the Future--> Marty McFly*michael fox
TV/Movies: Category: Defining Roles: Batman [the movie]--> Batman*michael keaton
TV/Movies: Category: Defining Roles: Butch Cassidy and the Sundance Kid--> the Sundance Kid*robert redford
TV/Movies: Category: Defining Roles: M.A.S.H. [the movie]--> Benjamin Franklin "Hawkeye" Pierce*donald sutherland
TV/Movies: Category: Defining Roles: The King and I--> The King of Siam*yul brynner
TV/Movies: Category: Dick Van Dyke Show: What phrase did Laura utter whenever she was in trouble or upset*oh rob
TV/Movies: Category: Disney Afternoon: DD What is the name of DarkwingUs pilot and sidekick*launchpad mcquack
TV/Movies: Category: Doctor Who: Multi-purpose gadget which helps the Doctor out of many sticky situations*sonic screwdriver
TV/Movies: Category: Famous Celebrities: What is the surname of 8-times married actress Zsa Zsa*gabor
TV/Movies: Category: Film 101: The relationship of the frame's width to its height*aspect ratio
TV/Movies: Category: Film Roles: Dirty Harry*clint eastwood
TV/Movies: Category: Film Roles: Superman*christopher reeve
TV/Movies: Category: Film Top Cops: Jean-Claude travels to the past*time cop
TV/Movies: Category: Flicks: Three vignettes dealing with two hitmen, a renegade boxer, and a bad date*pulp fiction
TV/Movies: Category: Flicks: Tom Hanks and Dan Akroyd track down the crimes of P.A.G.A.N*dragnet
TV/Movies: Category: Full House: Name Jesse and Rebecca's children (use "and")*nicky and alex
TV/Movies: Category: Get Smart: Siegfried's faithful sidekick and dummkopf*shtarker
TV/Movies: Category: Grease: She got her name because she was the best dancer at St. Bernadette's*chacha
TV/Movies: Category: Highlander: What was the name of Duncan MacLeod's monk mentor?*darius
TV/Movies: Category: Hollywood: He was a circus acrobat before acting*burt lancaster
TV/Movies: Category: Hollywood: She had children out of wedlock with Roger Vadim and Marcello Mastroianni*catherine deneuve
TV/Movies: Category: Hollywood: She is the mother of Jason Gould*barbra streisand
TV/Movies: Category: Hollywood: TV/Casino star who wrote much poetry, including Touch Me, & Touch Me Again*suzanne somers
TV/Movies: Category: Independent Films: 1991 Richard Linklater film which helped define "Generation X"*slacker
TV/Movies: Category: Independent Films: David Lynch's 1976 film filled with bizarre ideas and nightmare imagery*eraserhead
TV/Movies: Category: Independent Films: Mark Wahlberg allegedly wore a prosthetic device in this 1997 film*boogie nights
TV/Movies: Category: Indiana Jones: What kind of scientist is Indiana Jones*archaeologist
TV/Movies: Category: Indiana Jones: What possession does Indy never quite lose*hat
TV/Movies: Category: James Bond: Letter name for James Bond's superior*m
TV/Movies: Category: Kids in the Hall: The Kid that played Satan*mark mckinney
TV/Movies: Category: Last Names: Mr. Belvidere/The family's last name*owens
TV/Movies: Category: Letterman: According to Dave,not a man,woman,or child doesn't enjoy a cool refreshing ----------*beverage
TV/Movies: Category: MASH: Klinger wears dresses to earn himself one of these*section 8
TV/Movies: Category: Mostly Older Movies: Anne Francis and Robby The Robot starred in this SciFi classic:*forbidden planet
TV/Movies: Category: Mostly Older Movies: Musical with Richard Harris and Lynn Redgrave:*camelot
TV/Movies: Category: Movie Bombs: Gene Hackman, David Janssen, and Gregory Peck: Apollo13-ish space bomb*marooned
TV/Movies: Category: Movie In Which: Tom Cruise plays a high school football player*all the right moves
TV/Movies: Category: Movie Lines: The shark still looks fake*back to the future 2
TV/Movies: Category: Movie Musicals: The final sequence of this movie featured Kenny Jones instead of Keith Moon:*the kids are Alright
TV/Movies: Category: Movie Tag Lines 2: 1955: Teenage terror torn from today's headlines*rebel without a cause
TV/Movies: Category: Movie Tag Lines 2: 1995: Gluttony, greed, sloth, envy, wrath, pride, lust*se7en
TV/Movies: Category: Movie Tag Lines2: 1940: Walt Disney's Technicolor FEATURE triumph!*fantasia
TV/Movies: Category: Movie Tag Lines: 1995: What kind of man would defy a king?*braveheart
TV/Movies: Category: Movie Tag Lines: 1997: A message from deep space. Who will be the first to go?*contact
TV/Movies: Category: Movie That Features: A wife for a million*indecent proposal
TV/Movies: Category: Movie Trivia: Actor originally intended to be Wizard in "Wizard of Oz"*w c fields
TV/Movies: Category: Movie Trivia: Dean Martin's real last name*crocetti
TV/Movies: Category: Movie Trivia: Movie that featured the line "Here's looking at you kid."*casablanca
TV/Movies: Category: Movies: In The Two Jakes, he plays the part of Jake Gittes*jack nicholson
TV/Movies: Category: Muppet Mania: Name of the muppet that throws boomerang fish*lew
TV/Movies: Category: Muppet Mania: What lived in Ernie's garden (Sesame Street)*twiddle bugs
TV/Movies: Category: Music Movie Trivia: "Arnold" solos this Cole Porter song in Torch Song Trilogy*love for sale
TV/Movies: Category: Music Movie Trivia: This Gypsy song's footage is thought to be lost forever*together we go
TV/Movies: Category: Name That Actor: ...who played Mozart in "Amadeus"*tom hulce
TV/Movies: Category: Name That Actor: ...who played Obi Wan Kenobi*alec guiness
TV/Movies: Category: Name That Actor: ...who played The Nutty Professor*jerry lewis
TV/Movies: Category: Name Their Job: Michelangelo Antonioni*film director
TV/Movies: Category: Name Their Network: Sandy Rinaldo*ctv
TV/Movies: Category: PBS TV: What Mystery! series has had the most separate runs (9 through 1996)*inspector morse
TV/Movies: Category: Pop Music On Film: The name of "The Dating Game" show theme by Herb Alpert*spanish flea
TV/Movies: Category: Pop Music On Film: This singer can currently be seen on film in "Little Buddah"*chris isaak
TV/Movies: Category: Power Rangers: Identity of the Green Ranger*tommy
TV/Movies: Category: Pulp Fiction: book that vega read while expunging bodily wastes*madame blair
TV/Movies: Category: Pulp Fiction: color fabienna's toothbrush*red
TV/Movies: Category: Quality Movies: Produced "Howards End" and "The Remains of the Day"*ismael merchant
TV/Movies: Category: Quantum Leap: In "The Beast Within", Sam meets this legendary creature*bigfoot
TV/Movies: Category: Quantum Leap: The main character, Dr. Sam Beckett, was played by this actor*scott bakula
TV/Movies: Category: Relatives: Clark Gable's wife*carol lombard
TV/Movies: Category: Robotech: Japanese-translated name of the third storyline arc*sdb mospeada
TV/Movies: Category: Rush Limbaugh: The name of Rush's newletter*limbaugh letter
TV/Movies: Category: Sci Fi Movies: He played the captain of the guard in Dune*patrick stewart
TV/Movies: Category: Simpsons Cartoon: What is Lisa's future occupation according to the CANT test*homemaker
TV/Movies: Category: Sitcoms: City in which Mary Richards lived*minneapolis
TV/Movies: Category: Sports Actors: Who played Lou Gehrig in Pride of the Yankees*gary cooper
TV/Movies: Category: Star Trek Classic: This race is related to the vulcans*romulans
TV/Movies: Category: Star Trek Deep Space 9: Name of the trill inside Jadzia*dax
TV/Movies: Category: Star Trek Next Gen Chars: The 'bar' on the ship is located on deck*ten
TV/Movies: Category: Star Trek Next Gen Chars: This character was killed by a black oil slick*yar
TV/Movies: Category: Star Trek Next Gen Tech: Navigational deflector shields are generated by this*main deflector dish
TV/Movies: Category: Star Trek Next Generation: First encounter with Moriarty occurs in this episode*elementary dear data
TV/Movies: Category: Threes Company: Which one of the roommates was a medical professional*terri
TV/Movies: Category: TV Last Names: Family Matters/Carl and Harriet*winslow
TV/Movies: Category: UK TV: Which Eamonn spoke the first words on GMTV*holmes
TV/Movies: Category: Wizard Of Oz: Chevy Chase movie which takes place in 30's Hollywood*under the rainbow
TV/Movies: Category: World Of Disney: What color is the fairy that grants Geppetto's wish to bring Pinocchio to life?*blue
TV/Movies: First Class Flicks: Jailed rebel has personality conflict with chain gang warden, plans escape*cool hand luke
TV/Movies: Fox TV: On 90210: Who are the twins (name1 and name2)*brandon and brenda
TV/Movies: Game Shows: He hosted High Rollers, the Question, Concentration, Wizard of Odds*alex trebek
TV/Movies: Game Shows: Name of the buzzer on "Truth or Consequences"*beulah
TV/Movies: Gen X TV: Name the two curmudgeons in the balcony on "The Muppet Show" (Use and)*statler and waldorf
TV/Movies: Gen X TV: What cousin appeared in the last season making "The Brady Bunch" more annoying*oliver
TV/Movies: Get Smart: Said "People hate; robots love."*hymie
TV/Movies: Get Smart: Smart calls this Chinese KAOS agent "The Craw"*the claw
TV/Movies: Get Smart: This is issued by Control in different flavors every month*suicide pill
TV/Movies: Good Morning Vietnam stared what actor as Adrian Cronauer*robin williams
TV/Movies: Grease: "Renegade" star who played Tom, characterized as a dumb jock*lorenzo lamas
TV/Movies: He played the part of King Arthur in "Camelot."*richard harris
TV/Movies: Herschel bernardi starred in this 70's sitcom*arnie
TV/Movies: Highlander: What is the name of the female police officer in the first movie*brenda wyatt
TV/Movies: Hollywood: Actress who said, "You know how to whistle...put your lips together-and blow"*lauren bacall
TV/Movies: Hollywood: Musical superstar played a saloon pianist in South Sea Sinner*liberace
TV/Movies: How many stories did the towering inferno have*one hundred & thirty eight
TV/Movies: How old was Sally Field when she starred in Gidget*eighteen
TV/Movies: In "The Maltese Falcon," what was Peter Lorre's character named*joel cairo
TV/Movies: In Ian Fleming's James Bond series, what does the 00 in 007 stand for*license to kill
TV/Movies: In October of 1962, Johnny Carson succeeds him on the Tonight show*jack parr
TV/Movies: In the cartoon 'marine boy' what was marine boy's dolphin named*splasher
TV/Movies: In the episode "trilogy part 3", who did sam leap into*larry stanton
TV/Movies: In the movie "the manchurian candidate", how many people are feared dead in the hurricane*twenty
TV/Movies: Independent Films: 1996 Steve Buscemi film about suburban losers & their neighborhood bar*trees lounge
TV/Movies: Indiana Jones: Name the first challenge*breath of god
TV/Movies: Katharine hepburn & humphrey bogart played in what classic 1951 movie*african queen
TV/Movies: Lisa Loopner and Todd DiLaMuca are this*nerds
TV/Movies: Lone NYC cop against a highrise full of baddies*die hard
TV/Movies: Marx Movies: Whatever it is I am against it*horsefeathers
TV/Movies: MASH: Klinger wears dresses to earn himself one of these*section 8
TV/Movies: MASH: What was the name of the final episode*goodbye farewell & amen
TV/Movies: Mash: who said, "thanks for the eight pound baby nose"*klingers father
TV/Movies: Meg Ryan switches bodies with Sydney Walker*prelude to a kiss
TV/Movies: Monty Python: Lancelot's servant*concord
TV/Movies: Monty Python: The deformed knight had this many heads*three
TV/Movies: Monty Python: What was the name of the scholary knight at the witch trial*sir bedivere
TV/Movies: Moronic Duo 1: B&B called a 1-900 phone sex line and listened to a woman's all night*butt
TV/Movies: Mostly Classic Movies: Willis O'Brien, special-effects wizard, did this follow-up to King Kong*mighty joe young
TV/Movies: Mother Maybelle & the Carter family were regulars in this variety show*johnny cash show
TV/Movies: Movie Actresses: 40s star married to both Charlie Chaplin & Burgess Meredith*paulette goddard
TV/Movies: Movie Actresses: She starred in Casablanca, Gaslight, & Anastasia*ingrid bergman
TV/Movies: Movie Bombs: Gamera fights a stupid-looking monster with a head like a giant blade*gamera vs guiron
TV/Movies: Movie In Which: Billy Crystal and Gregory Hines play Chicago police detectives*running scared
TV/Movies: Movie Quotes: "Wendy....I'm home!"*the shining
TV/Movies: Movie set in 1950's Ireland about three friends and their first loves*circle of friends
TV/Movies: Movie That Features: Virginia Madsen, Peter O' Toole, Vincent Spano, and Muriel Hemingway*creator
TV/Movies: Movie Theme Songs: Movie that first featured "Whatever Will Be, Will Be"*man who knew too much
TV/Movies: Movie titles: around the world in ---------- days*eighty
TV/Movies: Movie titles: ---------- my way*going
TV/Movies: Movies & Music Jazzy Jeff & The Fresh Prince did a parody of this horror flick with Englund*nightmare on elm street
TV/Movies: Name one of the major stars in "nashville"*scott glenn
TV/Movies: Name the disney cartoon in which the character "belle" appears*beauty & the beast
TV/Movies: Name the only Elvis Presley film in which Elvis did not star*love me tender
TV/Movies: On Captain Midnight, what was Ichabod Mudd's nickname*icky
TV/Movies: On what british sitcom was 'all in the family' based*steptoe & son
TV/Movies: PBS TV: What was the 1989 Mystery! series about US-Soviet intelligence in Germany*game set and match
TV/Movies: Phrase coined to describe Errol Flynn's unmatched success with women*in like flynn
TV/Movies: Pierce Brosnan is a soviet agent (from the Forsyth novel)*the fourth protocol
TV/Movies: Pop Kulture: What is the longest running soap opera still on the air*guiding light
TV/Movies: Power Rangers: Name the robot the Green Ranger controlled*dragonzord
TV/Movies: Quantum Leap: In ep. 87, Sam Leaped into this celebrity who counseled Al about his sex life*dr. ruth westheimer
TV/Movies: Robotech: Robotech name of the third storyline arc*the invid invasion
TV/Movies: Seinfield: George's name if he was a porno star*buck naked
TV/Movies: Simpsons: what is the secret ingredient in a flaming moe*cough syrup
TV/Movies: Six man jewelry heist gone bad. Characters with colorful names*reservoir dogs
TV/Movies: Southpark: what is printed on chef's apron*chef
TV/Movies: Star Trek Deep Space 9: Who seeks the advice of a hologram to pursue Kira*odo
TV/Movies: Star Trek Next Gen Chars: This omnipotent letter makes live interesting on the old 1701D*q
TV/Movies: Star Wars: Original owner of C3PO and R2D2*biggs
TV/Movies: Tag Lines: 1997: Off the record, on the QT, and very hush-hush*l.a. confidential
TV/Movies: Tarantino: Harvey Keitel had a role similar to the one he had in PF in which film*point of no return
TV/Movies: the bunkers had neighbors who got their own series. Their last name was what*jefferson
TV/Movies: The Hollywood walk of fame star was awarded to "toby wing" for ----------*movies
TV/Movies: Theme Songs: Movie that featured "End of the Road"*boomerang
TV/Movies: Theme Songs: Movie that featured "Kiss"*under the cherry moon
TV/Movies: This Clark Gable/Claudette Colbert film kicked butt in 1934*it happened one night
TV/Movies: This is a classic film about a huge gorilla*king kong
TV/Movies: This is the black counterpart to American Bandstand*soul train
TV/Movies: This movie starring sam neill won best movie in the 1993 cannes film festival*piano
TV/Movies: TV Actors: Roseanne regular and host of Fernwood 2-Night*martin mull
TV/Movies: TV Actresses: Erica Kand on All My Children*susan lucci
TV/Movies: TV Actresses: Krystle on Dynasty*linda evans
TV/Movies: TV Actresses: Pamela Barnes Ewing on Dallas*victoria principal
TV/Movies: TV Roles: George Jefferson*sherman hemsley
TV/Movies: TV Top Cops: Bochco's classic emmy winning ensemble piece*hill street blues
TV/Movies: TV Transpo: Name T.C.'s helicopter tour company on "Magnum P.I."*island hoppers
TV/Movies: Twin Peaks: A burning smell is present whenever who is nearby*bob
TV/Movies: What 1988 film sequel brings its characters back to Earth from Antarea*cocoon: the return
TV/Movies: What color was Bullitt's car*green
TV/Movies: What comedy duo appeared on the first broadcast of the toast of the town"*dean martin & jerry lewis
TV/Movies: What does Mel Blanc's headstone say*that's all folks
TV/Movies: What is radar o'reilly's favorite drink*grape nehi
TV/Movies: What is the name of batman & robin's secret hiedout*batcave
TV/Movies: What is the name of the actor who plays "q" in the james bond films*desmond llewelyn
TV/Movies: What is the name of the movie that horace pinker appeared in*shocker
TV/Movies: What is the real name of bo derek*cathleen collins
TV/Movies: What is WC Fields' full name*william claude dunkenfield
TV/Movies: What late filmmaker was notorious for timing employees' trips to the soft drink machine*walt disney
TV/Movies: What lodge did Fred Flintstone & Barney Rubble belong to*the royal order of water buffalo
TV/Movies: What looney toons character used a univac to solve a mystery*porky pig
TV/Movies: What mechanical cop was described as "the thinking man's Terminator"*robocop
TV/Movies: What movie shows robin williams lining up to buy toilet paper*moscow on the hudson
TV/Movies: What popular youth oriented actor made a cameo appearance in Star Trek VI: The Undiscovered Country*christian slater
TV/Movies: What soul great appears in the flick Ski Party*james brown
TV/Movies: What transporter room aboard the Enterprise is Chief O'brien's favorite*three
TV/Movies: What was Archie Bunkers wife's name*edith
TV/Movies: What was barbara streisand's first film*funny girl
TV/Movies: What was Keanu Reeves' first big film*point break
TV/Movies: What was Samantha's mother's name in Bewitched*endora
TV/Movies: What was the first spin off in tv history*the andy griffith show
TV/Movies: What was the name given to robin hood's men in sherwood forest*merry men
TV/Movies: What was the name of lisa's pony in the simpsons*princess
TV/Movies: What was the name of the bartender in the tv series cheers*sam malone
TV/Movies: What was the name of the family in the "blondie" movies*bumstead
TV/Movies: what was the old tarts name in "the golden girls"*blanche
TV/Movies: What was the sequel to 'going my way'*the bells of st mary's
TV/Movies: What was the shape of lolita's sunglasses in the 1962 film*hearts
TV/Movies: What's the first video game to become a television show*pacman
TV/Movies: Whats the first name of batmans butler*alfred
TV/Movies: Where does George Jetson work*spacely space sprockets
TV/Movies: Which 'tarzan' swimmer was the first man to swim a hundred yards in less than a minute*johnny weismuller
TV/Movies: Which 1993 film starred Richard Attenborough*jurassic park
TV/Movies: which famous actor is michael douglas' father*kirk douglas
TV/Movies: Which film featured harold lloyd dangling from a clock tower*safety last
TV/Movies: Which motel chain paid connie frances million as a settlement because she was raped in one of its motels*howard johnson
TV/Movies: Which movie starred glenn close, jeff bridges & robert loggia*jagged edge
TV/Movies: which tv series was david cassidy in*partridge family
TV/Movies: Who created Maudie Frickett*jonathan winters
TV/Movies: who directed 'The Shining'*stanley kubrick
TV/Movies: Who got million to provide the voice of a baby for a 1990 movie*bruce willis
TV/Movies: Who is featured on 1984s the best of annette*annette funicello
TV/Movies: Who is melanie griffiths married to*antonio banderas
TV/Movies: Who is the anchorperson for channel 17 news*sally vacuum
TV/Movies: Who killed laura palmer (in "twin peaks")*leland palmer
TV/Movies: Who played "Meg Magrath" in Crimes of the Heart*jessica lange
TV/Movies: who played commander riker in 'star trek'*jonathon frakes
TV/Movies: who played god in 'oh god, book ii'*george burns
TV/Movies: who played lestat in 'interview with the vampire'*tom cruise
TV/Movies: Who played lulu hogg on dukes of hazzard*pearl shear
TV/Movies: Who played Major Pappy Boyington in the war drama Baa Baa Black Sheep*robert conrad
TV/Movies: Who played pontius pilate in "the last days of pompeii"*basil rathbone
TV/Movies: who played samantha in "bewitched"*elizabeth montgomery
TV/Movies: Who played sherlock holmes on the pbs series*basil rathbone
TV/Movies: Who played the character of jim rockford in the series the rockford files*james garner
TV/Movies: Who played the role of richard blaine in casablanca*humphrey bogart
TV/Movies: Who played the title role in the 1957 comedy the sad sack*jerry lewis
TV/Movies: who plays kelly in "married with children"*christina applegate
TV/Movies: Who plays ralph furley on threes company*don knotts
TV/Movies: Who starred with john travolta in the movie "Broken Arrow"*Christian Slater
TV/Movies: Who threw a badly aimed tomahawk on tv's "Tonight Show"*Ed Ames
TV/Movies: Who was born Marion Morrison*john wayne
TV/Movies: Who was Clark Kent*superman
TV/Movies: Who was Daisy the dog's owners*blondie & dagwood bumstead
TV/Movies: Who was nominated for best actress in 1944*greer garson
TV/Movies: who was pinocchio's father*geppetto
TV/Movies: Who was Robin Hood's girlfriend*maid marion
TV/Movies: Who was the voice of Rocket J. Squirrel & Natasha Fatale*june foray
TV/Movies: who were lucy & ricky's next door neighbours & best friends*fred & ethel
TV/Movies: Who won the only Oscar given to the hilarious ----------A Fish Called Wanda----------*kevin kline
TV/Movies: Who won the oscar for best actor in 1931 32*wallace beery
TV/Movies: Who won the Oscar for best ACTRESS in a Leading Role in 1928/1929 2nd Academy Awards for the Movie COQUETTE*mary pickford
TV/Movies: Who won the Oscar for best ACTRESS in a Leading Role in 1940 13th Academy Awards for the Movie KITTY FOYLE*ginger rogers
TV/Movies: Who won the Oscar for best ACTRESS in a Leading Role in 1944 17th Academy Awards for the Movie GASLIGHT*ingrid Bergman
TV/Movies: Who won the Oscar for best ACTRESS in a Leading Role in 1947 20th Academy Awards for the Movie THE FARMER'S DAUGHTER*loretta young
TV/Movies: Who won the Oscar for best ACTRESS in a Leading Role in 1957 30th Academy Awards for the Movie THE THREE FACES OF EVE*joanne woodward
TV/Movies: Who won the Oscar for best ACTRESS in a Leading Role in 1959 32nd Academy Awards for the Movie ROOM AT THE TOP*simone signoret
TV/Movies: Who won the Oscar for best ACTRESS in a Leading Role in 1968 41st Academy Awards for the Movie THE LION IN WINTER*katharine hepburn
TV/Movies: Who won the Oscar for best ACTRESS in a Leading Role in 1971 44th Academy Awards for the Movie KLUTE*jane Fonda
TV/Movies: Who won the Oscar for best ACTRESS in a Leading Role in 1973 46th Academy Awards for the Movie A TOUCH OF CLASS*glenda jackson
TV/Movies: Who won the Oscar for best ACTRESS in a Leading Role in 1975 48th Academy Awards for the Movie ONE FLEW OVER THE CUCKOO'S NEST*louise fletcher
TV/Movies: Who won the Oscar for best ACTRESS in a Leading Role in 1982 55th Academy Awards for the Movie SOPHIE'S CHOICE*meryl streep
TV/Movies: Who won the Oscar for best ACTRESS in a Leading Role in 1984 57th Academy Awards for the Movie PLACES IN THE HEART*sally field
TV/Movies: Who won the Oscar for best ACTRESS in a Leading Role in 1985 58th Academy Awards for the Movie THE TRIP TO BOUNTIFUL*geraldine page
TV/Movies: Who won the Oscar for best ACTRESS in a Leading Role in 1986 59th Academy Awards for the Movie CHILDREN OF A LESSER GOD*marlee matlin
TV/Movies: Who won the Oscar for best ACTRESS in a Supporting Role in 1938 11th Academy Awards for the Movie JEZEBEL*fay Bainter
TV/Movies: Who won the Oscar for best ACTRESS in a Supporting role in 1939 12th Academy Awards for the Movie GONE WITH THE WIND*hattie mcdaniel
TV/Movies: Who won the Oscar for best ACTRESS in a Supporting Role in 1944 17th Academy Awards for the Movie NONE BUT THE LONELY HEART*ethel barrymore
TV/Movies: Who won the Oscar for best ACTRESS in a Supporting Role in 1945 18th Academy Awards for the Movie NATIONAL VELVET*anne revere
TV/Movies: Who won the Oscar for best ACTRESS in a Supporting role in 1956 29th Academy Awards for the Movie WRITTEN ON THE WIND*dorothy malone
TV/Movies: Who won the Oscar for best ACTRESS in a Supporting Role in 1956 29th Academy Awards for the Movie WRITTEN ON THE WIND*dorothy malone
TV/Movies: Who won the Oscar for best ACTRESS in a Supporting role in 1957 30th Academy Awards for the Movie SAYONARA*miyoshi umeki
TV/Movies: Who won the Oscar for best ACTRESS in a Supporting Role in 1959 32nd Academy Awards for the Movie THE DIARY OF ANNE FRANK*shelley winters
TV/Movies: Who won the Oscar for best ACTRESS in a Supporting Role in 1960 33rd Academy Awards for the Movie ELMER GANTRY*shirley jones
TV/Movies: Who won the Oscar for best ACTRESS in a Supporting role in 1963 36th Academy Awards for the Movie THE V.I.P.S*margaret rutherford
TV/Movies: Who won the Oscar for best ACTRESS in a Supporting Role in 1975 48th Academy Awards for the Movie SHAMPOO*lee Grant
TV/Movies: Who won the Oscar for best ACTRESS in a Supporting Role in 1981 54th Academy Awards for the Movie REDS*maureen Stapleton
TV/Movies: Who won the Oscar for best ACTRESS in a Supporting Role in 1983 56th Academy Awards for the Movie THE YEAR OF LIVING DANGEROUSLY*linda hunt
TV/Movies: Who won the Oscar for best ACTRESS in a Supporting Role in 1988 61st Academy Awards for the Movie THE ACCIDENTAL TOURIST*geena davis
TV/Movies: Who's brother pulled strings to get her her first movie role*julia roberts
TV/Movies: Whose figure did Walt Disney use as a model for Tinkerbell*marilyn monroe
TV/Movies: Whose sister, Pamela appeared in the 1988 movie Sleepaway Camp 2*bruce springsteen
TV/Movies: Winning for Mrs Miniver, she took longer than 1.5 hours to say ----------thanks----------*greer garson
TV/Movies: Winona Ryder: Complete the title: "Welcome Home ---- ----------"*roxy carmichael
TV/Movies: ---------- street is where big bird & oscar the grouch live*sesame
TV: barbara bel geddes won a emmy in 1980 as outstanding lead actress in what drama series?*dallas
Tv: dean martin and jerry lewis appeared on the first telecast of this long-running variety show*the ed sullivan show
TV: What does ALF stand for*Alien Life Form
Tv: what was the favourite dish at mel's diner in alice*chili
Tv: who plays "roz" on night court*marsha warfield
Two 747's collided here in 1977*canary islands
Two short words are combined to give the name of which small stand with several shelves or layers for displaying ornaments*whatnot
under what name did norma jean mortenson become famous?*marilyn monroe
Unfriendly, cold and sexually unresponsive*frigid
Unscramble the letters of the words "no stamp" into a single english word*postman
Unscramble this Word: a a g d e m n*managed
Unscramble this Word: a a j o n v*navajo
Unscramble this Word: a a p n a m*panama
Unscramble this Word: a c l i i t*italic
Unscramble this Word: a d e s h g n*gnashed
Unscramble this Word: a d t i s n e*instead
Unscramble this Word: a e e e t g n*teenage
Unscramble this Word: a e r s o*arose
Unscramble this Word: a e s n i t p e l*palestine
Unscramble this Word: a i a f n r o c i l*california
Unscramble this Word: a i a s l f c*facials
Unscramble this Word: a k i d n m n*mankind
Unscramble this Word: a l i a m d r*admiral
Unscramble this Word: a n a e w b n*wannabe
Unscramble this Word: a n n a t m s t a h*manhattans
Unscramble this Word: a n r k o*akron
Unscramble this Word: a p s r n s w e e p*newspapers
Unscramble this Word: a r i k e t c*tackier
Unscramble this Word: a r k o e*korea
Unscramble this Word: a r l b e b b*babbler
Unscramble this Word: a r o s m r t*mortars
Unscramble this Word: a s e p r r p*rappers
Unscramble this Word: a s e t r r t*restart
Unscramble this Word: a s l p e d p*dapples
Unscramble this Word: a s o s n r m*ramsons
Unscramble this Word: a s t n u p e*peanuts
Unscramble this Word: a t i m s b p*baptism
Unscramble this Word: a t o b u*about
Unscramble this Word: a t s h r o w e m r*earthworms
Unscramble this Word: a y f t i b e*beatify
Unscramble this Word: b a r u m*burma
Unscramble this Word: b r t n o i a o*abortion
Unscramble this Word: b s a e n*beans
Unscramble this Word: b s i e r a*rabies
Unscramble this Word: c d d s a e u*adduces
Unscramble this Word: c e r e t l a*treacle
Unscramble this Word: c e r u g l a r u t i*agriculture
Unscramble this Word: c i e e p*piece
Unscramble this Word: c k t b t u o*buttock
Unscramble this Word: c p e r e*creep
Unscramble this Word: c r e a d*cedar
Unscramble this Word: c s e l r s a*scalers
Unscramble this Word: c s r h e a r*archers
Unscramble this Word: c u p d e i o c*occupied
Unscramble this Word: d a h e r*heard
Unscramble this Word: d a m y d o r r e*dromedary
Unscramble this Word: d c e d d e i*decided
Unscramble this Word: d e s e d*deeds
Unscramble this Word: d o m e c s m r o o*commodores
Unscramble this Word: d o o s c e n*secondo
Unscramble this Word: d t e s b e i*betides
Unscramble this Word: e a d k b*baked
Unscramble this Word: e b u t s c j*subject
Unscramble this Word: e d d b n u l*bundled
Unscramble this Word: e d t m t o l*mottled
Unscramble this Word: e d u d p e t*deputed
Unscramble this Word: e e a m n g r*germane
Unscramble this Word: e e c s t r f*refects
Unscramble this Word: e e l v t w*twelve
Unscramble this Word: e e r f t c a b o n*benefactor
Unscramble this Word: e e r v n*never
Unscramble this Word: e e t h r*there
Unscramble this Word: e h a s w h e*heehaws
Unscramble this Word: e i a s j r l*jailers
Unscramble this Word: e i t s e r c*recites
Unscramble this Word: e l a r l e b*relabel
Unscramble this Word: e l i m s r a*realism
Unscramble this Word: e l i s m w d*mildews
Unscramble this Word: e l l y r a*really
Unscramble this Word: e l p e s*sleep
Unscramble this Word: e m a d i m r*mermaid
Unscramble this Word: e m i u e r q*requiem
Unscramble this Word: e m u s d r p*dumpers
Unscramble this Word: e n f e c*fence
Unscramble this Word: e n i w r t*winter
Unscramble this Word: e o o c r l*cooler
Unscramble this Word: e o r w p*power
Unscramble this Word: e o s n a e h l d m*lemonheads
Unscramble this Word: e p e d a r d o s*desperado
Unscramble this Word: e p i s r r p*rippers
Unscramble this Word: e r e e n t r*terrene
Unscramble this Word: e r f o t*forte
Unscramble this Word: e r i t e m a*meatier
Unscramble this Word: e r t h e*three
Unscramble this Word: e r z t e s l*seltzer
Unscramble this Word: e s l t e b a*beatles
Unscramble this Word: e s o s n n l*nelsons
Unscramble this Word: e s t a s b r*breasts
Unscramble this Word: e s y a k l n*alkynes
Unscramble this Word: e t a n n r g*regnant
Unscramble this Word: e t i t o m n y s*testimony
Unscramble this Word: e u e s n r t*neuters
Unscramble this Word: e v o m s i*movies
Unscramble this Word: e w o k r n t*network
Unscramble this Word: e x b r o g a*gearbox
Unscramble this Word: e z o l y b n*benzoyl
Unscramble this Word: f e m s y l*myself
Unscramble this Word: f y i t f*fifty
Unscramble this Word: f y l i l*filly
Unscramble this Word: g a i r h e*hegira
Unscramble this Word: g i b r o n*boring
Unscramble this Word: g r e u u n d o n d r*underground
Unscramble this Word: g s e r a s*gasser
Unscramble this Word: g y h t e i*eighty
Unscramble this Word: h a e r h e t*heather
Unscramble this Word: h n g s t i*things
Unscramble this Word: h n m u a*human
Unscramble this Word: h n y r l e c*lyncher
Unscramble this Word: h o e e t l n e p*telephone
Unscramble this Word: h r s i s c o i r*cirrhosis
Unscramble this Word: h t h o a p g p r o*photograph
Unscramble this Word: h u l a g*laugh
Unscramble this Word: i a l g f n w*flawing
Unscramble this Word: i c u t i b s*biscuit
Unscramble this Word: i d e t t s h*shitted
Unscramble this Word: i e m l s*smile
Unscramble this Word: i e s e r i f d*fireside
Unscramble this Word: i g i t n j l*jilting
Unscramble this Word: i h e r e t*either
Unscramble this Word: i h i m u l t*lithium
Unscramble this Word: i l e d a m s*mislead
Unscramble this Word: i l t l s*still
Unscramble this Word: i n e g v n t*venting
Unscramble this Word: i n n s u g o d*sounding
Unscramble this Word: i p e e r l s t*reptiles
Unscramble this Word: i s a g t n t*tasting
Unscramble this Word: i s p s r e u r*surprise
Unscramble this Word: i s r d e r a*raiders
Unscramble this Word: i s t r e d a*diaster
Unscramble this Word: i y c t t s s s i l*stylistics
Unscramble this Word: j n a a p*japan
Unscramble this Word: k a b l c*black
Unscramble this Word: k a o l a*koala
Unscramble this Word: k o r c s o*crooks
Unscramble this Word: k s a u m n s*unmasks
Unscramble this Word: l a i r d f o*florida
Unscramble this Word: l d e s r e u*eluders
Unscramble this Word: l e p s e e s r*sleepers
Unscramble this Word: l e s s a s m o*molasses
Unscramble this Word: l g i s g o o*gigolos
Unscramble this Word: l h n u c*lunch
Unscramble this Word: l i n y l p a*plainly
Unscramble this Word: l n d r e b o*blonder
Unscramble this Word: l o d u c*could
Unscramble this Word: l s o d r*lords
Unscramble this Word: l s t m r o a*mortals
Unscramble this Word: l t a s d c w i*wildcats
Unscramble this Word: l w e r f o*flower
Unscramble this Word: l w l o y e*yellow
Unscramble this Word: m d g n e s i*smidgen
Unscramble this Word: m i n c t o a u c e m*communicate
Unscramble this Word: m l s c i i a*islamic
Unscramble this Word: m r e o t r e*remoter
Unscramble this Word: m s l b d e a*bedlams
Unscramble this Word: m t g i h*might
Unscramble this Word: n a n e l g e d t*entangled
Unscramble this Word: n d l h l o a*holland
Unscramble this Word: n d n l a f i*finland
Unscramble this Word: n e a s t u s*unseats
Unscramble this Word: n e t l e a r*eternal
Unscramble this Word: n g d g a r i*grading
Unscramble this Word: n g k g w a i*gawking
Unscramble this Word: n g p p p o i*popping
Unscramble this Word: n g t b t u i*butting
Unscramble this Word: n g t f t i i*fitting
Unscramble this Word: n h k a s*shank
Unscramble this Word: n h k i t*think
Unscramble this Word: n l b r e i*berlin
Unscramble this Word: n l d s d y u e*suddenly
Unscramble this Word: n l o n w u b*unblown
Unscramble this Word: n n c n a o*cannon
Unscramble this Word: n o r a i o t*ontario
Unscramble this Word: n r e g i r a*rangier
Unscramble this Word: n r h t e a o*another
Unscramble this Word: n r k e s i*sinker
Unscramble this Word: n r v e e*never
Unscramble this Word: n s e u v*venus
Unscramble this Word: n s s w r o e*worsens
Unscramble this Word: n t d e o r e*erodent
Unscramble this Word: n t o r c e u*recount
Unscramble this Word: n t p n t e r i u e*turpentine
Unscramble this Word: n u l m d u p e*pendulum
Unscramble this Word: n y r g a h u*hungary
Unscramble this Word: o a l t c b*cobalt
Unscramble this Word: o e n d e c z*cozened
Unscramble this Word: o e n p e p r*propene
Unscramble this Word: o e o n c r h*coehorn
Unscramble this Word: o f e r h o*hoofer
Unscramble this Word: o f s u o u c c i n*confucious
Unscramble this Word: o g a o o z n g l r*gorgonzola
Unscramble this Word: o g i s n d w*dowsing
Unscramble this Word: o i c l i c d*codicil
Unscramble this Word: o i s r e p n r s*prisoners
Unscramble this Word: o k s c r c e w r*corkscrew
Unscramble this Word: o m r b d e o*bedroom
Unscramble this Word: o n l t e m o*moonlet
Unscramble this Word: o r a s p t r*parrots
Unscramble this Word: o r k e s t*stoker
Unscramble this Word: o r s d e c u*coursed
Unscramble this Word: o s l w a a v*avowals
Unscramble this Word: o t e r l o*looter
Unscramble this Word: o t h d a e g r*goatherd
Unscramble this Word: o t h u g b r*brought
Unscramble this Word: o t l n g e i i u l*guillotine
Unscramble this Word: o u c s k o t n k*knockouts
Unscramble this Word: p s n e i*penis
Unscramble this Word: q e e e z s u*squeeze
Unscramble this Word: r a g l t u p o*portugal
Unscramble this Word: r b o y s l e*soberly
Unscramble this Word: r c t l a f a*fractal
Unscramble this Word: r d g g e d a*dragged
Unscramble this Word: r d w f r o a*forward
Unscramble this Word: r e s d o t a u p*outspread
Unscramble this Word: r e t a h b e*breathe
Unscramble this Word: r f a w s e*wafers
Unscramble this Word: r h a t e*earth
Unscramble this Word: r i n d e d a*drained
Unscramble this Word: r l e s e a*reales
Unscramble this Word: r m p e l t a*trample
Unscramble this Word: r n c e f a*france
Unscramble this Word: r p k a s*spark
Unscramble this Word: r p p d e g i*gripped
Unscramble this Word: r p r o h p s e t e o*troposphere
Unscramble this Word: r r e i d d e*derider
Unscramble this Word: r r r a e*rarer
Unscramble this Word: r t e d p e e*petered
Unscramble this Word: r t p c r a o*carport
Unscramble this Word: r w e s l c e*crewels
Unscramble this Word: r w r o a*arrow
Unscramble this Word: r y g u b*rugby
Unscramble this Word: s a r d p o e d e*desperado
Unscramble this Word: s c b u k*bucks
Unscramble this Word: s e c p o r*copers
Unscramble this Word: s e l p e a*please
Unscramble this Word: s e n c i h e*chinese
Unscramble this Word: s i t d m*midst
Unscramble this Word: s l b e l*bells
Unscramble this Word: s m e r u m*summer
Unscramble this Word: s m t i e*times
Unscramble this Word: s o b o k*books
Unscramble this Word: s o d o r*doors
Unscramble this Word: s p r o e*ropes
Unscramble this Word: s r o w o r*sorrow
Unscramble this Word: s s o h e*shoes
Unscramble this Word: s t e w p*swept
Unscramble this Word: t a s r w*warts
Unscramble this Word: t d i e r*tired
Unscramble this Word: t d o m i e h*ethmoid
Unscramble this Word: t e d r i n i*nitride
Unscramble this Word: t e p c m o e*compete
Unscramble this Word: t e r d a o u*outdare
Unscramble this Word: t e r e c x e*excrete
Unscramble this Word: t k h n a*thank
Unscramble this Word: t l e d r e a*related
Unscramble this Word: t l n a o e h*ethanol
Unscramble this Word: t n a r i*train
Unscramble this Word: t n w s e s e e y i*eyewitness
Unscramble this Word: t o l o e r w a*waterloo
Unscramble this Word: u c k n a s q d i*quicksand
Unscramble this Word: u d e n k c l*clunked
Unscramble this Word: u e g m u o b l r x*luxembourg
Unscramble this Word: u e g r a c o*courage
Unscramble this Word: u g i r e p d*pudgier
Unscramble this Word: u i a s l b r*burials
Unscramble this Word: u i n u t m o p l*plutonium
Unscramble this Word: u l a s w o t*outlaws
Unscramble this Word: u r e l a n c*nuclear
Unscramble this Word: u s i b l i o*bilious
Unscramble this Word: u y i g n b s*busying
Unscramble this Word: v e t e t e n n s h e*seventeenth
Unscramble this Word: v l c o a*vocal
Unscramble this Word: v s n e r t a*taverns
Unscramble this Word: v s u s e r*versus
Unscramble this Word: w d o l u*would
Unscramble this Word: w h t a c*watch
Unscramble this Word: w r d r e a a*awarder
Unscramble this Word: w t r o s*worst
Unscramble this Word: y f f i t*fifty
Unscramble this Word: y g u o n*young
Until 1947, what did 'gripe water' contain*opium
Upon his death in 1931, all non essential lights in the U S were turned off for one minute in his honor*thomas edison
Upon which river did Babylon stand*euphrates
US Captials - Califorina*Sacramento
US Captials - Hawaii*Honolulu
US Captials - Maine*Augusta
US Captials - Minnesota*St. Paul
US Captials - Mississippi*Jackson
US Captials - Missouri*Jefferson City
US Captials - New Hampshire*Concord
US Captials - Oregon*Salem
Useless Facts: A Dutch study indicated that 50 percent of the adult Dutch population have never flown in an airplane, and ---------------- percent admitted a fear of flying*twenty eight
Useless Facts: About 60 percent of all American babies are named after -------------*close relatives
Useless Facts: According to a recent survey, --------- percent of people who play the car radio while driving also sing along with it*seventy five
Useless Facts: An American Animal Hospital Association survey revealed that -------------- percent of dog owners sign letters or cards from themselves and their dogs*sixty two
Useless Facts: Banging your head against a wall can burn up to ----------- calories per hour*one hundred and fifty
Useless Facts: Butterflies taste with their*feet
Useless Facts: During the Spanish American War in 1898, there were 45 stars on the -------------*american flag
Useless Facts: Every time you lick a stamp, you're consuming 1/10 of a*calorie
Useless Facts: In 1977, according to the American Telephone and Telegraph Company, there were 14.5 telephone calls made for every 100 people in the ---------------------*entire world
Useless Facts: In 1977, less than 9 percent of physicians in the U.S. were -----------*women
Useless Facts: In ancient Egypt, Priests ---------- EVERY hair from their bodies, including their eyebrows and eyelashes*plucked
Useless Facts: It has been estimated that the typical American will spend an average of -------- years of his/her life reading newspapers*two
Useless Facts: King Francis I of France is reported to have paid master artist Leonardo da Vinci 4,000 gold crowns for his masterpiece -------------- but the king did not get immediate possession. Da Vinci kept the painting hanging on a wall of his chateau to the day he died*mona lisa
Useless Facts: Monaco boasts the highest per capita ownership of -------------- in the world. An early 1990s survey put the figure at one for every 65 people*rolls royces
Useless Facts: Since its introduction in February 1935, more than 150 million ----------- board games have been sold worldwide*monopoly
Useless Facts: The ------------------ were the first Asian colony to become independent following World War II. Today the country's population is approximately 60,000,000, and is comprised of many ethnicity. Many citizens are of Malay, Chinese, or Spanish descent*philippines
Useless Facts: The average American will eat 35,000 --------- during their life span*cookies
Useless Facts: The most memorable kiss in a motion picture was in -------------------, according to 25 percent of those polled*gone with the wind
Useless Facts: Two out of three adults in the United States have*hemorrhoids
Useless Facts: Zip code 12345 is assigned to -------------- in Schenectady, New York*general electric
Useless Trivia: "Crack" gets its name because it ---------- when you smoke it*crackles
Useless Trivia: ---------- newborn babies will be dropped in the next month*2,500
Useless Trivia: ---------- of all road accidents in Canada involve a Moose*0.3%
Useless Trivia: ---------- percent of the American population has never visited a dentist*forty
Useless Trivia: ---------- phone calls will be misplaced by telecoms service every minute*1,314
Useless Trivia: ---------- stands for ' Electrical and Musical Instruments'*emi
Useless Trivia: ---------- was once appointed Special Agent of the Bureau of Narcotics and Dangerous Drugs*elvis presley
Useless Trivia: 10% of ---------- fans replace the lenses on their glasses every 5 years whether they need to or not*star trek
Useless Trivia: A ---------- is a village without a church and a town is not a city until it has a cathedral*hamlet
Useless Trivia: A canton is the blue field behind the----------*stars
Useless Trivia: A female swine, or a sow, will always have a even number of---------- , usually twelve*nipples
Useless Trivia: A group of ---------- is called a Charm*finches
Useless Trivia: A group of ---------- is called a knot*toads
Useless Trivia: A large flawless emerald is worth more than a similarly large flawless----------*diamond
Useless Trivia: A lifetime supply of all the vitamins you need weighs only about ---------- ounces*eight
Useless Trivia: A man named ---------- Peterson is the inventor of the Egg McMuffin*ed
Useless Trivia: A man's ---------- contains between 7000 and 15,000 hairs*beard
Useless Trivia: A necropsy is an autopsy on----------*animals
Useless Trivia: A pack-day smoker will approx. lose 2 ---------- every ten years*teeth
Useless Trivia: A pig always sleeps on its ---------- side*right
Useless Trivia: A red-haired man is more likely to go ---------- than anyone else*bald
Useless Trivia: A square mile of fertile earth has ---------- earthworms in it*32,000,000
Useless Trivia: About ---------- years ago, most Egyptians died by the time they were 30*300
Useless Trivia: According to Genesis 1:20-22, the chicken came before the----------*egg
Useless Trivia: According to the Gemological Institute of America, up until 1896, India was the only source for ---------- to the world*diamonds
Useless Trivia: America media mogul Ted Turner owns 5% of----------*new mexico
Useless Trivia: Americans eat ---------- bananas a year*12 billion
Useless Trivia: An ---------- can stay under water for twenty-eight minutes*iguana
Useless Trivia: Annual growth of ---------- traffic is 314,000%*www
Useless Trivia: Apples, not---------- , are more efficient at waking up in the morning*caffeine
Useless Trivia: Assuming Rudolph was in front, there are 40320 ways to rearrange the other ---------- reindeer*eight
Useless Trivia: At any given time, there are ---------- thunderstorms in progress over the earth's atmosphere*1,800
Useless Trivia: At birth a panda is smaller than a mouse and weighs about ---------- ounces*four
Useless Trivia: At the equator the Earth spins at about ---------- miles per hour*1,000
Useless Trivia: Australian Rules football was originally designed to give ---------- something to play during the off season*cricketers
Useless Trivia: Average age of top ---------- executives in 1994: 49.8 years*gm
Useless Trivia: Before 1850, golf balls were made of leather and were stuffed with----------*feathers
Useless Trivia: Bernard Clemmens of London managed to sustain a ---------- for an officially recorded time of 2 mins 42 seconds*fart
Useless Trivia: Between 1902 and 1907 the same ---------- killed 436 people in India*tiger
Useless Trivia: By ---------- years old, Americans have watched more than nine years of television*sixty five
Useless Trivia: Cattle are the only mammals that are retro-mingent (they pee---------- )*backwards
Useless Trivia: Chances of a white ---------- in New York: 1 in 4*christmas
Useless Trivia: City with the most Roll Royces per capita:----------*hong kong
Useless Trivia: Clark Gable used to shower more than ---------- times a day*four
Useless Trivia: Cranberries are sorted for ripeness by bouncing them; a fully ripened cranberry can be dribbled like a----------*basketball
Useless Trivia: Dairy products account for about ---------- of all food consumed in the U.S*29%
Useless Trivia: Despite accounting for just one-fiftieth of body weight, the ---------- burns as much as one-fifth of our daily caloric intake*brain
Useless Trivia: Driving at ---------- miles per hour, it would take 258 days to drive around one of Saturn's rings*seventy five
Useless Trivia: During a lifetime, one person generates more than 1,000 pounds of ---------- blood cells*red
Useless Trivia: During the Cambrian period, about ---------- years ago, a day was only 20.6 hours long*500 million
Useless Trivia: During the---------- , banks first used Scotch tape to mend torn currency*depression
Useless Trivia: Eosophobia is the fear of----------*dawn
Useless Trivia: Everytime Beethoven sat down to write music, he poured ---------- water over his head*ice
Useless Trivia: Flamingo ---------- were a common delicacy at Roman feasts*tongues
Useless Trivia: Giant flying foxes that live in ---------- have wingspans of nearly six feet*indonesia
Useless Trivia: Gorillas often sleep for up to ---------- hours a day*fourteen
Useless Trivia: Hairstylist Anthony Silvestri cuts hair while----------*underwater
Useless Trivia: Howdy Doody had ---------- freckles*forty eight
Useless Trivia: Human ---------- is estimated to grow at 0.00000001 miles per hour*hair
Useless Trivia: Humans are the only primates that do not have ---------- in the palms of their hands*pigment
Useless Trivia: Hydroxydesoxycorticosterone and hydroxydeoxycorticosterones are the largest----------*anagrams
Useless Trivia: If ---------- imported just 10% of it's rice needs- the price on the world market would increase by 80%*china
Useless Trivia: If you travel across the Russia, you will cross ---------- time zones*seven
Useless Trivia: If you went out into space, you would explode before you ---------- because there's no air pressure*suffocated
Useless Trivia: Iguanas, ---------- and Komodo dragons all have two penises*koalas
Useless Trivia: In ---------- exists a tribe of tall. white people whose parrots are a warning sign against intruders*irian jaya
Useless Trivia: In 1936, American track star Jesse Owens beat a ---------- over a 100-yard course. The horse was given a head start*race horse
Useless Trivia: In Hartford, Connecticut, it is illegal for a husband to kiss his wife on----------*sundays
Useless Trivia: In Miami, Florida, roosting vultures have taken to snatching ---------- from rooftop patios*poodles
Useless Trivia: In most---------- , including newspapers, the time displayed on a watch is 10:10*advertisements
Useless Trivia: In the Great Fire of London in 1666, half of London was burnt down but only ---------- people were injured*six
Useless Trivia: In the summer, ---------- get a tan*walnuts
Useless Trivia: In---------- , the colour of mourning is violet*turkey
Useless Trivia: Issac Asimov is the only author to have a book in every ---------- -decimal category*dewey
Useless Trivia: It is illegal to hunt ---------- in the state of Arizona*camels
Useless Trivia: It takes about a half a gallon of water to cook macaroni, and about a ---------- to clean the pot*gallon
Useless Trivia: Jackals have one more pair of chromosomes than ---------- or wolves*dogs
Useless Trivia: Jill St. John, Jack Klugman, ---------- , Carol Burnett and Cher have all worn braces as adults*diana ross
Useless Trivia: John ---------- has entered over 5000 contests...and never won anything*bellavia
Useless Trivia: Li Hung-chang is the ---------- of Chop Suey*father
Useless Trivia: Mae West was once dubbed 'The statue of----------*libido
Useless Trivia: Male ---------- will try to attract sex partners with orchid fragrance*bees
Useless Trivia: Medical researchers contend that no disease ever identified has been completely----------*eradicated
Useless Trivia: Mice, whales, ---------- , giraffes and man all have seven neck vertebra*elephants
Useless Trivia: Michael Jordan makes more money from ---------- annually than all of the Nike factory workers in Malaysia combined*nike
Useless Trivia: Money is made of woven---------- , not paper*linen
Useless Trivia: Mongolia is the largest ---------- country*landlocked
Useless Trivia: Mongooses were brought to Hawai'i to kill rats. This plan failed because rats are ---------- while the mongoose hunts during the day*nocturnal
Useless Trivia: Most gemstones contain several elements, except the diamond; its all----------*carbon
Useless Trivia: Native Americans never actually ate turkey; killing such a timid bird was thought to indicate----------*laziness
Useless Trivia: Pennsylvania was the first colony to legalize----------*witchcraft
Useless Trivia: Rhinos are in the same family as---------- , and are thought to have inspired the myth of the unicorn*horses
Useless Trivia: Smelling ---------- or green apples can help you lose weight*bananas
Useless Trivia: The ---------- is illegal as a high school sport in all states except Rhode Island*hammerthrow
Useless Trivia: The monastic hours are matins, ---------- , prime, tierce, sext, nones, vespers and compline*lauds
Useless Trivia: The monastic hours are matins, lauds, prime, tierce, sext, nones, ---------- and compline*vespers
Useless Trivia:---------- , whales, elephants, giraffes and man all have seven neck vertebra*mice
Venus has how many moons*zero
Video Games: 'Secret of Evermore' was entirely produced in which country?*U.S.A
Video Games: The Nintendo 64 was titled under what name during production?*Project Reality
Video Games: Which character was introduced in 'Super Street Fighter II'?*Cammy
Video Games: Who is Mega Man's creator?*Dr. Light
Video Games: Who is the main character in the 'DeathQuest' series?*Lucretzia
Video Games: Who said "All life begins and ends with Nu...at least this is my belief for now..."?*Nu
Vientiane is the capital of ----------*laos
Vincent Van Gogh sold exactly one painting while he was alive, what was it*red vineyard at arles
Visual representation of individual people, distinguished by references to the subject's character, social position, wealth, or profession*portraiture
Weapon consisting of a long, sharp edged or pointed blade fixed in a hilt (a handle that usually has a protective guard at the place where the handle joins the blade)*sword
Weight-loss guru ---------- brings his fitness ideas to the little screen*richard
What 1,300- foot column of basalt do wyoming indians want to keep people from climbing*devil's tower
What 1995 movie was initially banned in malasyia because pigs are offensive to muslims*babe
What 2 countries border the Dead Sea*israel and jordan
What 80's cartoon was a showcase for 'New Wave' Music videos?*Kidd Video
What 80's spin off of a 70's tv show did Martin Lawrence play on?*What's Happening Now
What actor played seven roles in no way to treat a lady*rod steiger
What actor played the lead in the remake of breathless*richard gere
What actor was stung in "the sting"*robert shaw
What actress had made a million dollars by the age of 10*shirley temple
What actress played mrs margaret williams in the danny thomas show*jean
What airport in Uganda was the scene of a rescue drama in 1977*entebbe
What animal has red patches on its rear*mandrill
What animal is represented by the constellation Lacerta*lizard
What animal is represented by the constellation Monoceros*unicorn
What animal is thought to have inspired the myth of the unicorn*rhinocerous
what animals did hannibal lead over the alps for the first time?*elephants
What arabian peninsula nations recently merged under communist leadership*yemen
What are a group of gulls called*colony
What are elementary particles originating in the sun and other stars, that continuously rain down on the earth*cosmic rays
What are people encouraged to kiss under*mistletoe
What are scallops*shellfish
What are the Amish also known as*pennsylvania dutch
What are the annual awards for the best billboards (obies) named after*obelisks
What are the Boyoma and Tugela*waterfalls
What are the Christian names of the novelist P D James*phyllis dorothy
What are the clouds of magellan*galaxies
What are the only canines whose hair has a hook (or barb) on each individual follicle*dalmatians
What are the only two london boroughs that start with the letter 'e'*ealing
What are the separators on a guitar neck called*frets
What are the three main types of Greek columns*doric, ionic & corinthian
What are the two christian names of HE Bates*herbert ernest
What are the world's tallest trees*coast redwoods
What are tiny cracks in the glaze of pottery*crackle
What artist cut off his right ear*vincent van gogh
What australian food was discovered by john macadam*macadamia nuts
what averted an arab boycott of the 1948 summer olympics?*israel's exclusion
What bird is associated with lundy island*puffin
What bird is the offspring of a cob and a pen*swan
What body of water is fed from the south by the Wadi Araba & from the north by the river Jordan*the dead sea
What body organs did mae west say could be an asset if you hide them*brains
What body parts are oversized in a man suffering from gynecomastia*breasts
What Boston craftsman made George Washington's false teeth*paul revere
What boxer played the lead in the broadway musical buck white*muhammad ali
What branch of mathematics was devised by Sir Isaac Newton*calculus
What brand of footwear is endorsed by dr j*converse
What came down on jesus' head after he was baptised*dove
What can be measured in angstroms*wavelengths
What can't roosters do if they can't fully extend their necks*crow
What canadian city was carling beer first brewed in*toronto
What car was used in 'back to the future'*de lorean
What caused a separation of Baja, California and the rest of Mexico*The San
What character did Michael J Fox play in the film Back to the Future*marty mcfly
What children's book did Forrest Gump keep in his suitcase?*Curious George
What city boasts a world of coca cola pavilion featuring futuristic soda fountains*atlanta
What city has the world's largest black population*new york
What city in Nepal translates as "wooden temples"*Katmandu
What city is the setting for the US sitcom Cheers*boston
What city was originally called edo*tokyo
What civil war was fought between 1936 and 1939*spanish civil war
What cocktail does bourbon, sugar and mint make*mint julep
What cocktail is made from vodka and kahlua*black russian
What color is the blood of an octopus*pale bluish-green
What colour does a chameleon turn when its angry*black
What colours was the ferrari formula 1 car in the 1964 u.s.a grand prix*blue
What committee eventually developed a standard for the 'c' programming language*ansi
What company was founded by Sir Allan Lane in 1935*penguin books
What completed a journey of 19,500 miles with only three stops in August 1929*the graf zeppelin
What continent boasts the greatst number of Roman Catholics*south america
What continent is submerged*atlantis
What counrty would you visit to ski in the Dolomites*italy
What countries are known as the abc powers*argentina brazil chile
What country does Paul Hogan come from*australia
What country has the third most satellites in orbit*france
What country is the world's deepest mine located*South Africa
What country lies north of france and south of holland*belgium
What country officially limits women to one child*china
What country saw the origin of lawn tennis*england
What country would a Bulgarian with a good sense of direction walk through to reach Armenia by foot*turkey
What country's currency is the bolivar*venezuela
What country's people developed the crossbow*china
What craft uses a kiln and a kick wheel*pottery
What creatures call an apiary home*bees
What creatures do the Galapagos islands take their name from*Tortoises
What describes one complete turn of a rotating object*revolution
What did adolphe sax invent*saxophone
What did Americans call the first Cuban in space*castronaut
What did aristotle believe the heart was*seat of intelligence
What did dan aykroyd and john belushi quit 'saturday night live' to become*blues brothers
What did denmark sell to the u.s*virgin islands
What did Dr Godfrey invent in 1762*fire extinguisher
What did Gabriel Fahrenheit invent*thermometer
What did Grace Kelly become in 1956*princess
What did Moldavia & Walachia unite to become*romania
What did My Favorite Martian have to do before he could become invisible*raise his antenna
What did Neptune hold in his hand*trident
What did Sir Arnold Lunn begin in Switzerland*slalom skiing
What did the "P" in Roscoe P. Coltrane (from Dukes of Hazzard) stand for?*Purvis
What did the name 'battenberg' become*mountbatten
What did the Oshkosh steamer win*first automobile race
What did the shire's reeve become when the concept was brought to the u.s*sheriff
What disease is carried by the tsetse fly*sleeping sickness
What do diners in a restaurant use to take away their leftovers*doggy bag
What do people use to propel kayaks*paddles
What do Spanish dancers hold in their hands*castanets
What do spiders and ticks have in common*eight legs
What do table tennis players change after five points?*Service
What do the French call la manche*the english channel
What do the letters 'r.e.m.' stand for*rapid eye movement
What do the locals call the cloud that covers Table Mountain in Cape Town*tablecloth
What do the skunk, magpie and otter have in common they are all*black and white
What do trees get 90% of their nutrients from*air
What do you call a person whose iq is between 110-120*superior
What do you call an emasculated ram, whether or not he wears a bell*wether
What do you call the act of putting a word inside another (ie: abso bloody lutely.)*tmesis
What document is needed for one to enter a foreign country*passport
What does 'majuba' mean*place of rock pidgeons
What does 'n.b.a' mean*national basketball association
What does 'rio de janeiro' mean in portuguese*january river
What does 3 d mean*three dimensional
What does a 'postman' normally receive in kids' party games*kisses
What does a botanist study*plants
What does a brandophile collect*cigar bands
What does a chromophobic fear*certain colors
What does a i stand for*artificial intelligence
What does a person look like if described as 'wan'*pale-faced
What does a phyllophagus animal eat*leaves
What does a taxidermist do*stuff animals
What does a.n.c stand for*african national congress
What does an anemologist study*wind
What does an anthropophagist eat*people
What does an oologist study*eggs
What does an optician make*spectacles
What does bette davis' headstone say*she did it the hard way
What does blt stand for*bacon, lettuce, tomato
What does BMW stand for?*Bavarian Motor Works
What does breaking the sound barrier cause*a sonic boom
what does britain lose the lease on in 1997?*hong kong
What does cobol stand for*common business oriented language
What does jefferson davis' headstone say*at rest, an american soldier and
What does lacrimal fluid lubricate*eyes
What does the acronym "cpu" stand for*central processing unit
What does the computer acronym IKBS stand for*intelligent knowledge based system
What does the navajo term 'kemo sabe' mean*soggy shrub
What does the pancreas produce*insulin
What does the word chicane mean in the context of a game of bridge*a hand without any trumps
What does v.s.o.p. stand for on a bottle of brandy*very superior old pale
What does VAX stand for*Virtual Access eXtension
What dog shares his owner with Garfiled the Cat?*Odie
What drug is obtained from the poppy plant*opium
What eighties TV show starred Bruce Willis in a detective agency?*Moonlighting
What event was the interview of the Natural Born Killer, Mickey, to be held during?*The Superbowl
What ex-girl friend of prince andrew appeared naked on screen*koo stark
What falls out with phalacrosis*hair
What famous classical composer continued to compose great music after becoming deaf*Beethoven
What famous classical composer continued to compose great music after becoming deaf?*Beethoven
What FBI agent tracked Charles "Pretty Boy" Floyd to Ohio, where Floyd died*melvin purvis
What firm markets the B25 microcomputer*burroughs
What flag flies over gibraltar*union jack
What flower produces pink and white flowers in alkaline soil*hydrangea
What football team was previously known as the frankford yellow jackets*philadelphia eagles
What foreign country's phone book is alphabetized by first name*Iceland
What form of light comes at wavelengths below 360 nanometers*ultra violet
What french painter was the subject of somerset maugham's 'the moon and sixpence'*gauguin
What french phrase means 'well informed'*au courant
What fruit is usually used to make Marmalade?*Orange
What game challenges you to "double in" & "double out"*darts
What game of chance was originally called 'Beano'*Bingo
What game tiles were first made with a pocket knife*scrabble
What German city is best known for its Oktoberfest*munich
What german philosopher claimed morality required a belief in god and freedom*immanuel kant
What have over 80% of boxers suffered*brain damage
What hit lp did rockpile release in 1980*seconds of pleasure
What hobby was developed by the palmer paint company*painting by numbers
What indian tribe is associated with "the trail of tears"*cherokee
What instrument do doctors usually have around their necks*stethoscope
What instrument on a car measures distance*odometer
What is 'sapodilla' a type of*fruit
What is 240 minutes in hours*four
What is a 'crossbuck'*an x
What is a 'niblick'*golfer's nine iron
What is a 'tandoor'*clay oven
What is a bridge hand with no cards in one suit called*void
What is a chihuahua named after*mexican state
What is a conundrum*riddle
What is a female ferret*jill
What is a flat, round hat sometimes worn by soldiers*beret
What is a flowering plant that lives three or more years called*perennial
What is a group of ants*colony
What is a group of bass*shoal
What is a group of boars*singular
What is a group of buffalo*gang
What is a group of cockroaches*intrusion
What is a group of curs*cowardice
What is a group of grouse*pack
What is a group of hyenas*cackle
What is a group of larks called*exaltation
What is a group of leopards called*leap
What is a group of monkeys*troop
What is a group of pheasant*nest
What is a group of ponies*string
What is a group of roe deer*bevy
What is a group of this animal called: Ape*shrewdness
What is a group of this animal called: Chicken*brood
What is a group of this animal called: Dove*dule
What is a group of this animal called: Greyhound*leash
What is a group of this animal called: Jellyfish*smack
What is a group of this animal called: Swan*bevy
What is a group of this animal called: Swine*sounder
What is a group of this animal called: Woodpecker*descent
What is a leech a type of*worm
What is a male sheep*ram
What is a male swine called (giggle no ex boyfriends names...)*boar
What is a male whale called*bull
What is a mamba*a snake
What is a myocardial infarct*heart attack
What is a nibong a type of*palm tree
What is a noggin*a small cup
What is a portuguese man o' war*jellyfish
What is a pregnant goldfish*twit
What is a pyrotechnic display*fireworks
What is a regurgitation of acid from the stomach into the aesophagus*heartburn
What is a spat*baby oyster
What is a tightrope walker*funambulist
What is a tombstone inscription called*epitaph
What is a triangle with a 90 degree angle in it called*right angled triangle
What is a turkey's wishbone*furcula
What is a two-humped dromedary*camel
What is a young goose called*gosling
what is a young lion called?*cub
What is ambergis used in the making of*perfume
What is an angle greater than 90 degrees*obtuse
What is an extra lane on an uphill stretch of motorway provided for slow-moving vehicles called*crawler lane
What is an instrument for indicating the depth of the sea beneath a moving vessel called*bathometer
What is another name for crude oil*black gold
What is bed-wetting*enuresis
What is black gold*crude oil
What is borscht*soup
What is BSE in humans called*cjd
what is epidaurus famous for?*greek theatre
What is halloween*all hallow's eve
What is having a hole drilled through the cranium supposedly enabling people to reach a higher state of consciousness*trepanning
What is interpol*international criminal police
What is known as the " Palace of the Peak"*chatsworth house
What is MacGyver's first name?*Stace
What is made with a mix of charcoal, saltpetre and sulphur*gunpowder
What is Mr. Roger's first name*fred
What is name of the tubes that connect the ear & throat*eustachian
What is one of the items that the wood of the sycamore tree is used for*boxes
What is podobromhidrosis*smelly feet
What is Pogonophobia the fear of*Beards
what is raku?*japanese pottery
What is Rapec*type of snuff
What is schizophrenia*hallucinations & delusions
What is the 'pound' or 'number' symbol on the telephone*octothorpe
What is the actual vat in Romania*19%
What is the address Donald Duck lives at*1313 webfoot walk, duckburg, calisota
What is the approximate speed of light*186,000 miles per second
What is the aquatic nickname of Schubert's Piano Quintet in A*the trout quintet
What is the astrological sign for death*pluto
What is the average lifespan of a major league baseball*five to seven
What is the average temperature (f) at the South Pole*minus fifty six
What is the base twenty numbering system*vigesimal
What is the basic flavouring of kahlua*coffee
What is the better known name of writer Madame Dudevant*george sand
What is the birthplace (city) of the late John Candy*Toronto
What is the capital city of the Middle East state of Qatar*doha
What is the capital of albania*tirana
What is the capital of australia*canberra
What is the capital of Australia*canberra
What is the capital of connecticut*hartford
What is the capital of Florida*tallahassee
What is the capital of ireland*dublin
What is the capital of luxembourg*luxembourg
What is the capital of morocco*rabat
What is the capital of norway*oslo
What is the capital of Tasmania*hobart
What is the capital of tennessee*nashville
What is the capital of the Canadian province of British Columbia*victor1a
What is the capital of the US state of Delaware*dover
What is the capital of Venezuela*caracas
What is the Capital of: American Samoa*pago pago
What is the Capital of: Benin*porto-novo
What is the Capital of: Brazil*brasilia
What is the Capital of: Cambodia*phnom penh
What is the Capital of: Costa Rica*san jose
What is the Capital of: Czech Republic*prague
What is the Capital of: Ecuador*quito
What is the Capital of: Falkland Islands (Islas Malvinas)*stanley
What is the Capital of: Fiji*suva
What is the Capital of: Gibraltar*gibraltar
What is the Capital of: Guadeloupe*basse-terre
What is the Capital of: Hungary*budapest
What is the Capital of: Jamaica*kingston
What is the Capital of: Libya*tripoli
What is the Capital of: Mauritania*nouakchott
What is the Capital of: Netherlands*amsterdam
What is the Capital of: Oman*muscat
What is the Capital of: Portugal*lisbon
What is the Capital of: Slovenia*ljubljana
What is the Capital of: Sudan*khartoum
What is the Capital of: Swaziland*mbabane
What is the Capital of: Tonga*nuku'alofa
What is the Capital of: Turkey*ankara
What is the chemical name for water*hydrogen oxide
What is the chief monetary unit of Croatia*kuna
What is the circle of the earth at 0 degrees latitude*equator
What is the closest relative of the manatee*elephant
What is the collective noun for a group of crows*murder
What is the collective noun for a group of tigers*an ambush
What is the colour of the maple leaf on the Canadian flag*red
What is the common name for corporations formed to act as trustees according to the terms of contracts known as trust agreements?*Trust companies
What is the common name for the marine animals asteroidea*starfish
What is the common name given to the larvae of a crane fly*leatherjackets
What is the common term for the condition monochromatism*colour blindness
What is the connection between Good Times and Different Strokes?*Janet Jackson
What is the connection between Jeffersons and Good Times?*Janet Dubois
What is the correct name for an animal's pouch*marsupium
What is the criminal number of sideshow bob in 'the simpsons'*24601
what is the currency of venezuela?*bolivar
What is the current vat rate in south africa*14%
What is the discharge of a liquid from a surface, usually pores or incisions*exudation
What is the drink 'Southern Comfort' flavoured with*peaches
What is the drug that is used to treat Parkinson's disease*dopamine
What is the drummer's name in 'the muppet show'*animal
What is the eighth month of the year*august
What is the English statute of 1689 guaranteeing the rights & liberty of the individual subject*bill of rights
What is the fastest fish in the world*sailfish
What is the fear of being tickled by feathers known as*pteronophobia
What is the fear of certain fabrics known as*textophobia
What is the fear of children known as*pedophobia
What is the fear of clouds known as*nephophobia
What is the fear of computers known as*logizomechanophobia
What is the fear of darkness known as*lygophobia
What is the fear of dreams known as*oneirophobia
What is the fear of fears known as*pantophobia
What is the fear of flying known as*pteromerhanophobia
What is the fear of food or eating known as*sitophobia
What is the fear of germs known as*spermophobia
What is the fear of ghosts known as*phasmophobia
What is the fear of glass known as*nelophobia
What is the fear of hospitals known as*nosocomephobia
What is the fear of ice or frost known as*pagophobia
What is the fear of illness known as*nosemaphobia
What is the fear of light flashes known as*selaphobia
What is the fear of lockjaw, tetanus known as*tetanophobia
What is the fear of many things known as*polyphobia
What is the fear of mice known as*musophobia
What is the fear of movement or motion known as*kinetophobia
What is the fear of one thing known as*monophobia
What is the fear of outer space known as*spacephobia
What is the fear of pregnancy or childbirth known as*tocophobia
What is the fear of rabies known as*kynophobia
What is the fear of rabies or of becoming mad known as*lyssophobia
What is the fear of satan known as*satanophobia
What is the fear of sexual perversion known as*paraphobia
What is the fear of shellfish known as*ostraconophobia
What is the fear of suffering and disease known as*panthophobia
What is the fear of technology known as*technophobia
What is the fear of wet dreams known as*oneirogmophobia
What is the first letter of the Russian alphabet*a
What is the first name of the inventor of braille*louis
What is the first name of Webster, the man who published a dictionary still used today*noah
What is the flower that stands for: affection*mossy pear
What is the flower that stands for: affection*mossy saxifrage
What is the flower that stands for: aversion*indian single pink
What is the flower that stands for: betrayal*judas tere
What is the flower that stands for: boldness*pink
What is the flower that stands for: bonds*convolvulus
What is the flower that stands for: concealed love*motherwort
What is the flower that stands for: dangerous pleasures*tuberose
What is the flower that stands for: decrease of love*yellow rose
What is the flower that stands for: devotion*heliotrope
What is the flower that stands for: difficulties that i surmount*mistletoe
What is the flower that stands for: divine beauty*american cowslip
What is the flower that stands for: early friendship*blue periwinkle
What is the flower that stands for: early youth*primrose
What is the flower that stands for: elegance and grace*yellow jasmine
What is the flower that stands for: envy*crane's bill
What is the flower that stands for: poverty*evergreen clematis
What is the flower that stands for: remembrance*rosemary
What is the flower that stands for: retaliation*scotch thistle
What is the flower that stands for: silliness*fool's parsley
What is the flower that stands for: sincerity*fern
What is the flower that stands for: splendid beauty*amarylis
What is the flower that stands for: strength*cedar
What is the flowering shrub Syringa usually called*lilac
What is the former name of Istanbul*constantinople
What is the former name of the Russian city Volgograd*stalingrad
What is the French term for "d day"*j
What is the french word for 'mistake'*faux pas
What is the full name of the creator of "Jeeves & Wooster"*pelham grenville wodehouse
What is the heaviest element*uranium
What is the honeymoon capital of the world*niagara falls
what is the international telephone code for the uk?*forty four
What is the largest (in population) state/territory in Australia*new south wales
What is the largest city in Texas*houston
What is the largest gland in the human body*liver
What is the largest inhabited castle*windsor castle
What is the largest island in Asia*borneo
What is the largest item on any menu in the world*roast camel
What is the largest lake in Central America*lake nicaragua
What is the largest lake in the u.s*superior
What is the largest landlocked country*mongolia
What is the largest lizard*komodo dragon
What is the largest ocean*pacific ocean
What is the latin phrase meaning 'in the original arrangement'*in situ
What is the leaf of a fern called*frond
What is the longest English word that only has one vowel*strengths
What is the longest insect*walking stick
What is the longest river in Scotland*tay
What is the longest river in the world*nile
What is the longest running race at the olympic games*marathon
What is the longest strait in the world*malacca
What is the longest thing an "abseiler" carries with him*rope
What is the main ingredient in an omelet*egg
What is the main ingredient of most shampoos*water
What is the maximum number of degrees in an obtuse angle*one hundred and
What is the more popular narne of the plants belonging to the genus galanthus*snowdrop
What is the Morse code representation for the letter T*single dash
What is the most air polluted city in the united states*los angeles
What is the most commonly spoken language in India*hindi
What is the most essential tool in astronomy*telescope
What is the most important mineral for strong bones & teeth*calcium
What is the most westerly county of Ireland*kerry
What is the name for a sexual disorder in which a person obtains gratification by receiving physical pain or abuse*masochism
What is the name for music that is transmitted orally or aurally (taught through performance rather than with notation, and learned by hearing)*folk
What is the name given to Indian food cooked over charcoal in a clay oven*tandoori
What is the name given to the fortified gateway of a castle*barbican
What is the name given to thin pieces of crisp toast*melba
What is the name given to young deer*fawns
What is the name of a device used to stem the flow of blood?*tourniquet
What is the name of a formal, written accusation of crime against a person, presented by a grand jury to a court, and upon which the accused person is subsequently tried*indictment
What is the name of Jonny Quest's Dog*Bandit
What is the name of Mr. Krane's dog on Frasier?*Eddie
What is the name of the capital of Alberta (Canada)*edmonton
What is the name of the capital of Saskatchewan (canada)*regina
What is the name of the detective in john dickson carr novels*gideon fell
What is the name of the Dukes of Hazzards car?*General Lee
What is the name of the Freelings' dog in "Poltergeist"?*Ebuzz
What is the name of the fruit that looks like a hairy lychee*rambutan
What is the name of the group of Muslim scholars who have fought for control of Afghanistan in recent years*taliban
What is the name of the island that separates the two waterfalls at Niagara*goat island
What is the name of the man who gave his name to the World Cup Trophy*david rimet
What is the name of the official residence of the president of France*the elysee palace
What is the name of the pig that Jim Davis draws*orson
What is the name of the spaceship in the film 'Alien'*nostromo
What is the name of the Tokyo Stock Market Index*nikkei
What is the next-to-next-to-last event*antepenultimate
What is the nickname for Alaska*land of the midnight sun
What is the nickname for Kentucky*bluegrass state
What is the nickname for North Dakota*sioux state
What is the number of blue razor blades a given beam can puncture*gillette
What is the occupation of Mary Poppins*nanny
What is the official language of new caledonia*french
What is the only 'real food' astronauts can take into space*pecan nuts
What is the only bird that can fly backwards?*Hummingbird
What is the only country with a bible on its flag*dominican republic
What is the only English word formed by the first three letters of the alphabet*cab
What is the only English word that ends in the letters "mt"*dreamt
What is the only female animal that has antlers*caribou
What is the only metal that is liquid at room temperature*mercury
What is the only word in the English language that ends in the letters 'mt'*dreamt
What is the point value of the 'f' in scrabble*four
What is the proper term for a guinea-pig*cavy
What is the ratio of the speed of an object to the speed of sound in the surrounding medium*mach speed
What is the real name of the 'man of arms' in 'He man and the Masters of the Universe'*Duncan
What is the relatively constant, but dynamic internal environment necessary for life*homeostasis
What is the roughly circular hollow feature on the top of a volcano called*caldera
What is the sacred animal of India*cow
What is the saltiest sea in the world*dead sea
What is the seventh day of the week*saturday
What is the shape of the pasta 'tortlloni' based on*venus's navel
What is the shape of the US President's office*oval
What is the shortest and bloodiest of Shapespeare's plays*macbeth
What is the significance of the moth found in the Harvard Mark I computer*First computer "bug"
What is the singular of dice*die
What is the slogan on New Hampshire license plates*live free or die
What is the slowest moving land mammal*sloth
What is the speed of light*186,000 miles per second
What is the study of animals known as*zoology
What is the study of mankind called*anthropology
What is the study of prehistoric plants & animals called*paleontology
What is the study of word origins*etymology
What is the substance obtained from acacia trees that is used in medicine*gum arabic
What is the sum of 2y + 32y + 56y*ninety y
What is the sum of 9685z + 235z - 1800z + 2z*8122z
What is the symbol of the democratic party*donkey
What is the throwing event making up part of the ancient greek pentathlon, in which a circular object had to be thrown*discus
What is the traditional trade of aspiring bullfighters*bricklaying
What is the transformation of inhospitable planets into hospitable ones*terraforming
What is the tribal african word for dowry*lobola
What is the unit of currency in Hungary*forint
What is the US equivalent of the S.A.S*delta force
What is the widest-ranging ocean bird*albatross
What is the world's deepest lake?*Lake Baikal
What is the world's largest rodent*capybara
What is the young of this animal called: Rat*kitten
What is the young of this animal called: Shark*cub
What is the young of this animal called: Turkey*poult
What is the young of this animal called: Zebra*foal
What is tina turner's real name*annie mae bullock
What is tuberculosis*consumption
What is uruguay's chief port*montevideo
What is usually served at bedouin feasts*roast camel
What is Venezuela named after?*Venice
What is William Hague's middle name*jefferson
What island group is off the east coast of southern South America*falkland islands
What Italian city is considered the fashion capital?*Milan
What keeps one from crying when peeling onions*chewing gum
What kind of 'mate' produces a tie in a chess game*stalemate
What kind of animal has a tail pinned on it in a birthday party game*donkey
What kind of animal was Rikki Tikki Tavi in The Jungle Book*mongoose
What kind of animals are impalas, elands & kudus*antelopes
what kind of cat is used in purina(tm) commercials?*white persian
What kind of clay can potters heat to a higher temperature earthenware or stoneware*stoneware
What kind of condition is 'protanopia'*colour blindness
What kind of creature is a funnel web*spider
What kind of creature is a Lorikeet*a parrot
what kind of dog was "rin tin tin"?*german shepherd
what kind of gun does the movie's "dirty harry" pack?*magnum
What kind of music does an "MOR" radio station play*middle of the road
What kind of nuts are ground up to make marzipan*almonds
What kind of pain is a migraine*headache
What kind of shoe is nailed above the door for good luck*horseshoe
What kind of sword did Thundar the Barbarian have?*A Sun Sword
What kind of tradesman uses a 'plunger'*a plumber
What late television commentator closed with "good night and good luck"*edward r murrow
What legendary monster does Seattle secretary Katie Martin believe to be the father of her furry faced son*bigfoot
what legendary us magazine publisher was born in tengchow, china?*henry luce
What loaded gaming devices were found in the ruins of Pompei*dice
What lollies are well known for rolling down the aisles at the movies*jaffas
What made up the Bouquet in the 70's TV series starring Susan Penhaligon*barbed wire
What magazine was the first to be distributed widely through grocery stores*family circle
What major city is served by Gatwick Airport*london
What major law was violated in the movie Smokey and the Bandit?*Smuggling beer
What make and model of car does Nash Bridges drive?*A 1971 Plymouth Barracuda convertible.
What makes a solution saline*salt
What makes brown bread healthier than white bread*wholemeal
What mammal moves so slowly that green algae can grow undisturbed on it's fur*sloth
What media format did the denon company help pioneer*compact discs
What metal forms one twelfth of the earth's crust*aluminium
What mixture is used to calm crying babies
What motto ends merrie melodies cartoons*that's all folks
What movie featured Reece's Pieces as a crucial part of the story, because the director couldn't obtain the rights to use M&M's?*E.T.
What muscles provide about 200 pounds of force*jaw muscles
What name is given to an isolated mountain peak protruding through an ice sheet*nunatuk
What name is given to the blend of Black China and Darjeeling teas, flavoured with oil of Bergamot*earl grey
What name is given to the broad gap between the outermost and the brightest of Saturn's rings*cassini division
What name is given to the effect that the Earth is gradually becoming warmer*global warming
What name is given to the study of living things in their environment*ecology
What name is popularly applied to twins congenitally united in a manner not incompatible with life or activity?*siamese twins
What name was given to the 8th century Muslim invaders of Spain*moors
What nation is nicknamed the 'regaa boyz'*jamaica
What nationality is designer Karl Lagerfield*german
What nationality is the keyboards wizard Vangelis*greek
What nationality was actress Greta Garbo*swedish
What nationality was the first person who walked in space*russian
What natural disasters are ranked in severity by the Saffir Simpson scale*hurricanes
What neighbouring country did Iraq go to war with in 1980*iran
what new york city avenue divides the east side from the west side?*fifth avenue
What New York street is famous for its theatres*broadway
What New Zealand native invented bungee jumping?*AJ Hackett
What New Zealand native was the first man to climb Mt. Everest*Sir Edmund
What NFL team was formerly known as the Portsmouth Spartans*detroit
What Northeastern European country's capital is Tallinn?*Estonia
What northern country Helsinki the capital of*finland
What novel by Geoffrey Household was about an attempt to kill Hitler*rogue male
What novel was alexandra ripley hired to pen a sequel to*gone with the wind
What number does VII mean in roman numerals*seven
What number is at 6 oclock on a dartboard*three
What operating system in used on an ibm as400*os400
What organ of the body is particularly affected by hepatitis*liver
What organ will most often suffer permanent damage if you have amoebic dysentery*the liver
What organization helped defend earth in "ultra man"*science patrol
What organization was given the only Nobel Peace Price awarded during World War I?*The Red Cross
What other common name is given to a rook in chess*castle
What part of the body does arthritis particularly affect*the bone joints
What part of the body has a crown, a neck & a root*tooth
What part of the body is particularly affected by pneumonia*lungs
What part of your body is elastic, waterproof, washable & fits you very well*skin
What peace treaty ended WWI*treaty of versailles
What peninsula does Mexico occupy*yucatan peninsula
What percentage of alcohol is contained in a 100 proof mixture*fifty
What period is the age of fish*devonian
What physical disability is also known as nanism*the condition of being a dwarf
What piano man used to play for Bette Middler and then went on to his own career and made Hits like "Mandy" and "Copacabana"*barry manilow
What piece of music commemorates military action that took place on 16th May 1943*the dam busters
What plane did Aerospatiale of France & the British Aircraft Corp. develop*the concord
What portuguese territory will revert to china in 1999*macao
What position does a sloth spend its day in*upside down
what president's hobbies included pitching hay, fishing, and golf?*calvin coolidge
What presidential ticket was dubbed bozo and the pineapple*gerald ford and
What product is sold with "just for the taste of it"'*diet coke
What protein makes blood red*haemoglobin
what queen did edmund spenser dedicate his faerie queene to?*elizabeth i
What race's runners refer to the noisy section along wellesley college as the "screech tunnel"*the boston marathon
What relation was Queen Victoria to George III*granddaughter
What releases an explosive charge of air that moves at speeds up to 60 mph*cough
What religion follows the teachings of the prophet Mohammed*islam
What religious movement did joseph smith found*mormonism
what religious movement did joseph smith found?*mormonism
What replaced English as the official language of Kenya in 1974*swahili
What represent the body and blood of Christ in the service of Holy Communion*bread and wine
What reptilian feature evolved in feathers*scales
What rifle accessory originated in Bayonne, France, in 1641*bayonet
what river divides the dutch capital of amsterdam in two?*amstel
What river does the Grand Coulee Dam dam*columbia
What river had 40 million fish killed by insecticide in 1969*rhine
What river was Francisco de Orellano the first to travel the length of*amazon
What rock group uses roman numerals on all of its album covers*chicago
What scandinavian country owned iceland from 1262 to 1944*denmark
What sea creature uses its chest as a table while floating on its back*sea
What sea is between italy and yugoslavia*adriatic
What seaport's name is spanish for 'white house'*casablanca
What sentence uses every letter of the alphabet*the quick brown fox jumps over the lazy dog
What sequence is this the start of: 2 4 6 8 10 12 14*even numbers
What serious umderwater ailment was named after a Victorian notion of chic posture*the bends
What sexually ambigious prisonmate was often dubbed "Mrs. Hitler"*rudolf hess
What shakespearean king was actually king of scotland for 17 years*macbeth
What Shakespearean play features Iago*othello
What shape are playing cards in India*round
What short lived TV western did Rod Serling produce after Twilight Zone*loner
What show did Claire Danes get her start on?*My So-Called Life
What show was a spin-off of Transformers?*Go-Bots
What show/game has characters such as bulbasaur and pikachu*pokemon
What sinatra hit did he dooby dooby do in*strangers in the night
What singer's February 6th birthday is a national holiday in Jamaica?*Bob Marley
What sitcom was "The Facts of Life" a spinoff of?*Different Strokes
What small Arctic rodents are said to, but don't, commit suicide in mass plunges into the sea*lemmings
What song did bobby hebb sing to his brother in 1966*sunny
What soprano simply titled her autobiography beverly*beverly sills
What soul great appears in the flick Ski Party*james brown
What South American country produces the most coffee*brazil
What Spanish artists surrelistic paintings feature items such as clock faces*salvador dali
What Spanish islands are Gomera, Hierro & Lanzarote a part of*canary islands
What sport do the Kansas City monarchs participate in*baseball
What sport does "FISA" govern*auto racing
What sport uses barrier stalls*horse racing
What sport was observed by Captain James Cook in 1771*surfing
What state has the most workers employed by the travel & tourism industry*california
What state is 'the hoosier state'*indiana
What state is mount mckinley in*alaska
What state is only part of the U S by treaty*texas
What stretch of water seperates Australia from Tasmania*bass strait
What struck honshu island, japan in 1934 killing 4,000 people*typhoon
What structure in the back of the brain governs motor control*cerebellum
What style of dancing was popularized with rap music?*Break Dancing
What subject did Mr. Chips teach*latin
What talk show hostess gave her guests the fewest opportunities to speak, according to a 1996 msu survey*oprah winfrey
What technique did Patrick Steptoe and Robert Edwards pioneer*in vitro fertilization
What ten volume tome did Victor Hugo give the world in 1862*les miserables
What term describes the study of the behaviour of materials and substances at very low temperatures*cryogenics
What term is given to that part of the Earth which can support life*biosphere
What term is used to describe the process of extracting poison from snakes*milking
What term was used from 1914 onwards to describe music emanating from New Orleans*jazz
What Texan slammed back more bourbon and branch water than any character in TV history?*j. r. ewing
What the most north-eastern state of the contiguous u.s*maine
What three words mean the same as 5,880,000,000,000 miles*one light year
What tropic passes through Mexico*cancer
What TV personality did Doritos commercials?*Jay Leno
What two characters from Sesame Street got their names from the movie "It's a Wonderful Life"*bert and ernie
What two countries were known as "the yellow peril" in the 1890's*china & japan
What type of animal was selected to test the first electric toothbrush*the dog
What type of craft is the u.s's airforce one*boeing 747
What type of insect performs a waggle dance*hive bee
What type of metal is used in the filament of an electric light bulb*tungsten
What type of number describes the ratio of the speed of a plane to the speed of sound*mach
What type of scientific equipment was named after the german Bunsen*burner
What type of storm has a central calm area, called the eye, which has winds spiraling inwardly*hurricane
What U S state was once an independent republic*texas
What u.s. vice-president said 'some newspapers dispose of their garbage by printing it'*spiro agnew
What units are used to measure the size of pearls*grains
What US state includes the telephone area code 503*oregon
What US state includes the telephone area code 504*louisiana
What us state includes the telephone area code 612*minnesota
What us state includes the telephone area code 615*tennessee
What US state includes the telephone area code 703*virginia
What vehicles are involved in the 'Tour de France'?*Bicycles
What vitamin deficiency causes rickets*vitamin d
What war did Joan of Arc's inspirational leadership help end*the hundred
What was 1990s most populous U S state*california
What was A.A. Milne's first name*alan
What was Al Bundy's nickname during the dream sequence with all of the women in his bedroom?"Al night long
What was Al Capone's favorite bullet proof car*cadillac
What was al capone's favorite bullet-proof car*cadillac
What was ALF's girlfriend from Melmac's name?*Rhonda
What was astronaut edwin aldrin's nickname*buzz
What was discovered at Sutter's Mill, California in 1848*gold
What was Donald Fagen's first solo album title (1982)?*The Nightfly
What was elvis presley's twin brother's first name*garon
What was first sold at the 1904 St Louis worlds fair*ice cream cones
What was formerly called the Christian Revival Association and the East London Christian Mission*salvation army
What was JFK's nickname for his daughter Caroline?*Buttons
What was Lestat's mother's name?*gabrielle
What was Louise Joy Brown the first of*test tube baby
What was Maggie Seaver's maiden name on Growing Pains?*Maggie Malone
What was Massachusetts' logical choice for an official state dessert, in 1996*boston cream pie
What was originally called the pluto platter*frisbee
What was Potsie's last name on Happy Days*weber
What was rembrandt's surname*van rijn
what was richard bach's best selling book?*jonathan livingston seagull
What was Rizzo's real name in Grease?*Betty
What was robert montgomery's profession*actor
what was Rocky Balboa's nickname in the ring?*the italian stallion
What was st. paul's trade before he converted*tent-maker
What was Terry's surname in the television series Minder*mccann
What was the challanging method of catching a fly in Karate Kid?*Using chopsticks
What was the country of Botswana called before 1966*bechuanaland
What was the famous line uttered by an old woman in Wendy's ads?*Where's The Beef?
What was the first commercial readymix food*pancake mix
What was the first computer software company to go public on the New York Stock Exchange?*Cullinet
what was the first disney film to feature stereophonic sound?*fantasia
What was the first motion picture to have a synchronized musical score*Don
What was the first name of the baby girl who fell down the well?*Jessica
What was the first u.s consumer product sold in the soviet union*pepsi cola
What was the first video Mtv played?*Video Killed the Radio Star
What was the former German name of the Czech town of Ceske Budejovice*budweis
What was the former name of Burkina Faso in Africa*upper volta
What was the leading cause of death in the late 19th century*tuberculosis
What was the name (4 letters) of the New York night club that helped launch the career of several early new wave groups?*CBGB's
What was the name of Eddie Murphy's character in Beverly Hills Cop?*Axel Foley
what was the name of jacques cousteau's research ship?*calypso
What was the name of jim henson's muppet hound on the jimmy dean show*rowlf
What was the name of King Arthur's sword*excalibur
What was the name of Norman Beaton's barber's shop which was also the title of the TV series*desmond's
What was the name of the actress who played "Melonie" on the show "Webster and Melonie"?*Heather O' Rourke
What was the name of the bar that the characters from "Three's Company" frequented?*Regal Beagle
What was the name of the bartender on The Love Boat?*Isaac Washington
What was the name of the detective agency in Moonlighting?*Blue Moon Detective Agency
What was the name of the first synthetic plastic made in 1908*bakelite
What was the name of the helicopter service that was the cover for Airwolf?*Santini Air
What was the name of the home that Sofia Patrillo lived in before moving in with her daughter on the Golden Girls*Shady Pines
What was the name of the I.B.M. computer which played Chess against Gary Kasparov*deep blue
What was the name of the monster that attacked Luke in the trash compactor in Star Wars?*A dianogaIn
What was the name of the movement founded by the Pole Lech Walesa*solidarity
What was the name of the multi-colored cube you had to re-organize?*Rubik Cube
What was the name of the operatic diva who gave her name to a peach dessert*dame nellie melba
What was the name of the Other short-lived spinoff of "Three's Company"*"Three's a Crowd"
What was the name of the owner of the talking horse, Mr. Ed on TV*wilbur post
What was the name of the party dog that that was Budwiser's mascot in the late eighties?*Spuds McKenzie
What was the name of the police character played by Roy Scheider in the film Jaws*martin brody
What was the name of the South African Prime Minister murdered in 1966*hendrik verwoerd
What was the name of the submarine which sank the General Belgrano during the Falklands conflict*hms conqueror
What was the name of the Titanic's sister ship*Olympic
What was the nickname of Charles Heidsick, the 19th Century French wine producer*champagne charlie
What was the number of the squadron which flew the Dambusters mission in 1943*617
What was the profession of Lancelot 'Capability' Brown*landscape gardener
What was the royal residence after st james court*buckingham palace
What was the screen name of the lead character in The Untouchables*elliot
What was the title of Jung Chang's account of growing up in China*wild swans
What was the world's principal Christian city before it fell to the Ottoman Turks in 1453*constantinople
What was Webster's adopted mom and dad's name*Poppadouupalus
What was William H. Bonney's nickname*Billy the kid
What was willie mosconi famed for shooting*pool
What weapon is tattooed on Glen Campell's arm*dagger
What were Club Nouveu originally known as?*Timex Social Club
What were dachshunds bred to hunt*badgers
What were the first names of T E Lawrence, known as Lawrence of Arabia*thomas edward
What woman is the wife of prince phillip, the mother of anne, andrew, charles and edward, and the daughter of george vi*elizabeth ii
What woman is thought of as the greatest trick shot artist of all time?*annie oakley
What word is used for a female fox*vixen
what word is used for a female fox?*vixen
What word links these: bar, cereal, continental*breakfast
What word links these: battery, rain, test*acid
What word links these: brother, gloves, skin*kid
What word links these: cab, frequency, station*radio
What word links these: cake, tea, egg*cup
What word links these: cavalry, chore, name*household
What word links these: centre, certificate, record*medical
What word links these: comic, singer, soap*opera
What word links these: contract, dodger, proposal*draft
What word links these: detector, polish, scrap*metal
What word links these: face, round, time*about
What word links these: growth, policy, recovery*economic
What word links these: meal, set, work*piece
What word may be used to refer to a group of gnats*horde
What would you do with "ackee" in jamaica*eat it
What would you expect to find in a vespiary*wasps
What ws Balki Bartokamus' occupation when he lived in Mypos?*Sheep Herder
What year did Chernobyl explode*1986
What year did the first nudist colony open*1903
What year was a U2 pilot shot down for spying*1960
What year was film introduced to replace glass in making photographic negatives*1891
What year was The Bible printed using moveable type*1455
What year was the first tooth extraction under anaesthetic performed*1846
What year was the last woman hung in England*1955
What's a 10-20 to a police officer*location
What's a dead body of an animal called*carcass
What's a microchip made of*silicon
What's a natatorium*swimming pool
What's Krypton's state at standard temperature & pressure*gaseous
what's the circulation of winds around a low pressure system called?*cyclone
What's the fastest sea dwelling mammal*the dolphin
What's the highest mountain in the 48 contiguous u.s. States*mount whitney
What's the international radio code word for the letter "J"*juliet
What's the longest river in the U S*mississippi river
What's the Malayan sun bear's main claim to fame*smallest bear
What's the most common name in nursery rhymes*jack
What's the most common term of endearment in the U.S*honey
What's the most frequently ingested mood altering drug*caffeine
What's the most valuable crop in burma, laos and thailand*poppy
What's the name of the counting system in which four is written "100"*binary
What's the name of the dragon in the Ivor the Engine stories*idris
What's the name of the second book in the Bible*exodus
What's the nickname of the Iowa state football team*cyclones
What's the official state sport of alaska*dog sledding
What's the only property an orthodox Hindu woman can own?*Jewelry
What's the second most spoken language on earth*english
What's the sky king's home, near the town of grover, called*flying crown
What's the square root of one-quarter*one half
What's the world's largest fresh-water island*manitoulin
What's white sugar mixed with to make brown sugar*molasses
Whats the computer term "bit" short for*binary digit
Whats the name of the large wooded area in which Robin Hood was supposed to have lived*sherwood forest
Whats the nearest galaxy to our own*andromeda
Whats the official language of Morocco*arabic
When did Henry Ford build his first car*1896
When does a full moon always rise*sunset
When is Saint George's day celebrated*april 23rd
When not fighting crime, what did Underdog do for a living*shoeshine boy
When someone is clumsy or awkward, especially with their hands, they are often said to be "all ...." These*thumbs
When was george jones inducted into the country music hall of fame*1992
When was the first credit card issued*1900
When was the first jet aircraft flown*1941
When was the first Mad Max film released*1979
When was the first play staged at Londons Globe Theatre*1599
When was the first toothbrush with bristles invented*1498
When was the incandescent lamp invented*1879
When was the quadruplex telegraph invented*1864
When was the rechargable storage battery invented*1859
When was the shortest war in history*1896
Where are phalanges*hand
Where are the Guiana Highlands*northern south america
Where did Stalin, Churchill, Attlee and Truman meet in 1945 to determine the future of Germany after their unconditional surrender*potsdam
Where did the bay of pigs take place*cuba
Where did the birkenhead sink*danger point
Where did the incas live*peru
Where did the mafia originate*sicily
Where do Grand Prix drivers put their cars at the beginning of a race*grid
Where do the english monarchs live*buckingham palace
Where does Dilbert think of inventions*In the bathtub
Where does the Iditarod dog sled race take place*alaska
Where in Italy is the wine Marsala made*sicily
Where in London was the Great Exhibition of 1851 held*hyde park
Where in the body are the Haversian canals*inside bones
Where in the body is the axilla*armpit
Where is antofagasta*chile
Where is bill gates' company based*redmond, washington
Where is charlottetown*prince edward island
Where is crystal palace*london
Where is eurodisney*paris, france
Where is frostbite falls*minnesota
Where is Huracan stadium*buenos aires
Where is it polite to stick your tongue out at your guests*tibet
Where is Mount Rushmore*south dakota
Where is mount vesuvius*italy
Where is the Bernabau stadium*madrid, spain
Where is the Devil's Tower*wyoming usa
Where is the fictional television station bdrx located*bedrock
Where is the guggenheim museum*new york city
where is the space needle?*seattle
Where is the wailing wall*jerusalem
Where was the last major american indian resistance to white settlement*wounded knee
Where was the record for most snowfall in a day, on february 7 1916*alaska
Where was the septuagint written*alexandria
Where would you find vox humana and vox angelica together*on an organ
where's the 19th hole on a golf course?*clubhouse
Where, in 1955, was one of the worst accidents in motor racing history, when 82 spectators were killed*le mans
Where, on a horse are its withers*shoulder
Which "daring young man on the flying trapeze" gave his name to a garment*jules leotard
Which 'first lady of jazz' died in June 1996*ella fitzgerald
Which 1978 film from the book of the same name by Ira Levin, tell of the cloning of Adolf Hitler*the boys from brazil
Which 50's Actress was born Vera Jayne Palmer*jane mansfield
Which 60's folk artist sang the lyrics "god told abraham kill me your son. abe said man you must be puttin me on"*bob dylan
Which 9-fingered pop pianist starred in the film Its all Happening?*Russ Conway
Which acid builds up in the body during excessive exercise*lactic
Which actor was born Maurice Micklewhite*michael caine
Which actor's autobiography is entitled Dear Me*peter ustinov
Which actress starred in the film Love Story*ali mcgraw
Which African capital city is named from the Greek meaning 'three towns'*tripoli
Which airline has the registration prefix 'vr'*cathay pacific
Which animal floats in water*porcupine
Which Australian author wrote Illywhacker and Oscar and Lucinda*peter carey
Which Australian state capital was named in honour of a British Prime Minister*melbourne
Which author created Svengali*georges du maurier
Which band had members Robert palmer, Andy and John Taylor, and Tony Thompson?*The Power Station
Which bank did the jailed Nick Leeson work for and ruin*barings
Which best selling car with a production spanning some 30 years is to be replaced by the "Focus"*ford escort
Which bone in the human body is at the front but sounds like it should be at the back*sternum
Which book by James Joyce takes palce on a single Dublin day in June 1904*ulysses
Which British town is famous for its cutlery production*sheffield
Which Canadian city was originally called Bytown*ottawa
Which car company makes the 'Avensis'*toyota
Which cellular structures are composed of DNA*chromosomes
Which character was portrayed by Robert Redford in the film Out of Africa*dennis finch hatton
Which chemical element has the ancient name Stannum*tin
Which chemical element is named after the 1959 winner of the Nobel Prize for Physics*laurencium
Which city's airport is the home base for Cathay Pacific Airlines*hong kong
Which classic dish contains strips of steak cooked in a wine sauce with sour cream*stroganoff
Which comedian created the character of maude frickert*jonathan winters
Which comic actor who died in 1977 entered a competition to find his look alike, anonymously, and only came third*charlie chaplin
Which country blew up a greenpeace ship in new zealand*france
Which country has won the most Olympic gold medals at 10,000 metres*finland
Which country is known as the Hashemite Kingdom*jordan
Which country is the biggest consumer of wine*france
Which country is widely acknowledged to have the largest Jewish population*united states
Which country produces Dao wine*portugal
Which country saw the Mau Mau uprising?*Kenya
Which country was invaded by Soviet troops in August 1968*czechoslovakia
Which country was the first to legalise abortion*iceland
Which country was the setting for The Flame Trees of Thika*kenya
Which country's name means "equator"*ecuador
Which country's national flag consists only of a green field*libya
Which county lies between the north sea and greater london*essex
Which creature do Eskimos (or Inuit) call a nanook*polar bear
Which department of the us government did eliot ness work for*treasury
Which detective was played by Jack Webb in Dragnet*sgt joe friday
Which dog was originally bred to hunt badgers*dachshund
Which drink does Melanie Sykes advertise on TV*boddingtons
Which drug can be extracted from the bark of the cinchona tree*quinine
Which element has the chemical symbol Cs; capital C lower-case s*caesium
Which English composer was born near Worcester in 1857 and died in 1934*edward elgar
Which English King met Francis I of France on the 'Field of the Cloth of Gold'*henry viii
Which English king's coronation was postponed because he was suffering from appendicitis*edward vii
Which European city is served by Fiumicino airport*rome
Which european country will lose its independence if there is no heir to the throne*monaco
Which famous artist took up painting with his left hand when he lost the use of his right hand at the age of sixty*Leonardo da Vinci
Which famous film actor, who died of lung cancer in 1957, used his real name but dropped his middle name of de Forest*humphrey bogart
Which famous museum is in paris, france*louvre
Which famous sporting venue is above NewYorks Pennsylvania Station*madison square gardens
Which film director's films include "Midnight Express" and "Bugsy Malone"*alan parker
Which film links novelist Ira Levin and Sharon Stone*sliver
Which film star is the real life husband of Goldie Hawn*kurt russell
Which film, directed by Sydney Pollack, won the 1985 Academy Award for Best Picture*out of africa
Which food product did Henry Cooper advertise in 1984*shredded wheat
Which forename, deriving from the Germanic 'rulehard', has been held by three English kings*richard
Which former 'Neighbours' star had a hit with 'Any Dream Will Do'*jason donovan
Which French athlete won both the 200m and the 400m on the track at the 1996 Atlanta Olympic Games*maria-jose perec
Which french dramatist's works include Phedre and Andromaque*jean racine
Which French mathematician, "the father of Modem Mathematics", invented analytical or co-ordinate geometry*rene descartes
Which fungal plant disease particularly affects brassicas*club root
Which game uses the largest ball*earthball
Which German actress appeared in the film "Witness for the Prosecution*marlene dietrich
Which Gloucestershire town, famous for its abbey, lies on the confluence of the Severn and Avon*tewkesbury
Which golfer has won the British Open most times since 1945*tom watson (5)
Which great battle took place from July 1st to November 18th 1916?*The Battle of the Somme
Which Greek island is also a variety of lettuce*cos
Which group had the hit album 'White on Blonde'*texas
Which hero of tv and cinema fights an unending battle for 'truth, justice, and the American way*superman
Which Hollywood heart throbs real name was Roy Scherer*rock hudson
Which houses fought the war of the roses*lancaster and york
Which is Britain's largest native carnivore*badger
Which is considered the most powerful piece on the chess board*queen
Which is the highest capital city in Europe*madrid
Which is the largest cathedral*st peter's
Which is the largest of the Canadian Provinces and Territories*northwest territories
Which is the largest river forming part of the u.s-mexico border*rio grande
Which is the last of the four Grand Slam tennis tournaments to be played in the year*us open
Which is the only bird that can see the colour blue*owl
Which is the only English word to both begin and end with the letters U-N-D ?*Underground
Which is the world's second largest monolith*ayers rock
Which is the worlds tallest grass*bamboo
Which jazz cornettist composed and recorded "Davenport Blues" in 1925*b1x beiderbecke
Which jazz pianist composed the jungle music "Black and Tan Fantasy" with Bubber Miley, and recorded it with his band in 1927*duke ellington
Which jockey rode a Derby winner called Pinza*gordon richards
Which kellogg's cereal was advertised by tusk tusk the elephant*coco
Which law did sir isaac newton discover when he was only twenty three years old*law of universal gravitation
Which lawyer broke the law by refusing to be finger-printed in the Transvaal during 1907*Gandhi
Which liqueur gives the cocktail "Tequila Sunrise" its red glow*grenadine
Which London MP is more famous as an actress*glenda jackson
Which london station handles trains directly to the continent, through the channel tunnel*waterloo
Which London's church's other name is the Collegiate Church of St. Peter*westminster abbey
Which major river flows through gloucester*severn
Which Mammal has the highest blood pressure*giraffe
Which modem country was formerly Nyasaland*malawi
Which modern author wrote The Regeneration trilogy*pat barker
Which mountain peak is the highest in the Western Hemisphere*aconcagua
Which mountain range forms a geographical boundary between Europe and Asia*urals
Which movie about a TV news show won Peter Finch a posthumous Oscar*network
Which museum now occupies the site of the old Bedlam Hospital in London*imperial war museum
Which musical includes the Barbara Dickson/Elaine Page song I Know Him So Well?*Chess
Which musical was based on the play The Matchmaker?*Hello Dolly
Which nazi leader had his 6 children poisoned prior to his own death*goebbels
Which nineteenth century author is buried in Samoa*robert louis stevenson
Which Nobel Prize winner wrote 'The Old Man and the Sea*ernest hemingway
Which object was known as a Churchwarden*long clay pipe
Which opera singer was known as "La Stupenda"*joan sutherland
Which opera, composed by Saint-Saens, and first performed in 1877, is set in Palestine*samson and delilah
Which ovine expression is used for a disreputable member of a family or group*black sheep
Which people slide down a pole to help them to get to work quickly*firemen
Which planet circles the sun every 84 years*uranus
Which planet did John Couch Adams and Urbain Leverrier work out the existence and position of before it could actually be seen*neptune
Which planet is orbited by the moon Charon*pluto
Which planet was discovered in 1846*neptune
Which planet was the 'Planet of the Apes'*earth
Which popular singer of the 80's has the real name Christopher Davidson*chris de burgh
Which port on the River Douro is the second largest city in Portugal*oporto
Which Prime Minister introduced Income Tax*pitt the younger
Which religion believes in the Four Noble Truths*buddhism
Which rock musician committed suicide in Scattle on 5th April 1994*kurt cobain
Which Russian word means openness*glasnost
Which saint founded a monastery at Iona in the sixth century*saint columba
Which Saint translated the Vulgate bible*jerome
Which sci-fi writer adapted his own book for the movie Pet Sematary*stephen king
Which Scottish Quarter day is on August 1st*lammas
Which sea route connects the North Atlantic with the Beaufort Sea and the Pacific Ocean*the northwest passage
Which Shakespeare character described himself as having "Loved not wisely but too well"*othello
Which singing King died in 1965*nat king cole
Which south east Asian city was formerly called Krung Threp*bangkok
Which sport is featured in the book and film "This Sporting Life"*rugby league
Which star of films such as 'Ryan's Daughter' died in 1997*robert mitchum
Which state became the 14th state of the u.s*vermont
Which state forms an enclave at the heart of the city of Rome*vatican city
Which strait separates Russian and Alaska*bering strait
Which substance, occurring naturally in fruit, causes jams and preserves to set*pectin
Which tennesee williams play is about a sicilian-american woman*rose tattoo
Which tennis star wore denim shorts during matches*andre agassi
Which town in the US had Clint Eastwood as its mayor*carmel
Which tube line goes to Brixton*victoria
Which two fighting ships other than the 'arizona' were sunk at pearl harbor*oklahoma and utah
Which two fruits are an anagram of each other*lemon and melon
Which two male fish give birth*sea horse and pipe fish
Which types of wood are most often used for firewood in the home*hardwood
Which U S president was fatally shot in 1881*garfield
Which U.S. president gave the 'four freedoms of democracy' speech- ie freedom from want; freedom from fear; freedom of worship and freedom of speech*franklin d roosevelt
Which UK city, other than London, has a station called Charing Cross*glasgow
Which US golfer was killed when his plane crashed in 1999*payne stewart
Which US writer wrote The Naked and the Dead*norman mailer
Which vegetable is used if a dish is described as 'a la Bretonne'*haricot beans
Which was the first 'spaghetti western' starring Clint Eastwood*a fistful of dollars
Which was the first apostle to be stoned to death*stephen
Which wedding anniversary is coral*thirty fifth
Which word comes from the Roman "where three roads meet" as a place where messages were left*trivia
Which word follows Juliet, Kilo and Lima*mike
Which word, taken from the French, translates literally as 'rotten pot'*potpourri
Which World Champion heavyweight boxer held the title for the longest*joe louis
Which writer's latest work, Birds of Prey , features the Courtneys - the family that appeared in his first, When the Lion Feeds , published in 1964*wilbur smith
Which year did Jemima Goldsmith marry Imram Kahn*1995
Who advocated the planting peanuts and sweet potatoes to replace cotton and tobacco (i.e. crop rotation)?*George Washington Carver
Who appeared in 'st. elmo's fire', 'the scarlett letter' and 'striptease'*demi moore
Who are santa's reindeer, in alphabetical order*blitzen, comet, dancer,
Who ate chicken little*foxy loxy
who ate chicken little?*foxy loxy
Who bought manattan island for the equivalent of 24 dollars*peter minuit
Who built the 'cherokee' and 'commanche' aircraft*piper
Who claimed that, in the Garden of Eden, God spoke Swedish, Adam spoke Danish, & the serpent spoke French*swedish philologist
Who composed "Invitation to the Dance " in 1819*weber
Who composed "Messiah"*handel
Who composed the music for the ballet 'l'apres-midi d'un faune'*claude
Who composed the musical piece Carmina Burana*carl orff
Who composed the opera, "The Queen of Spades"*tchaikovsky
Who conquered the matterhorn in 1865*edward whymper
Who controls more than 80% of the world's rough diamond supply*de beers
Who crashed out of the 1995 Tour de France after just 92 seconds*chris boardman
Who created john blackthorne*james clavell
Who created the 'grinch'*dr seuss
Who created the comic strip 'Doonesbury'*garry trudeau
Who created WinnieThe Pooh*a a milne
Who cremated on the banks of the ganges river on january 31, 1948*mahatma
Who did orson welles play in the film 'the third man'*harry lime
Who did pat sajak play on the soapie 'days of our lives'*kevin hathaway
Who did Roger Bannister beat at the Commonwealth Games of 1954*john landy
Who did Spain fight in the 1808-1814 Peninsular War?*Portugal
Who did Vivian Vance play on 'the lucy show'*Vivian Bagley
Who died three days before groucho marx*elvis presley
Who directed 'the breakfast club'*john hughes
who directed the monochrome (sepia) sequences at the beginning and end of "the wizard of oz" (1939)?*king vidor
Who discovered gold on the witwatersrand*george harrison
Who discovered oxygen*joseph priestley
Who founded the Church of Scientology*l. ron hubbard
Who gave excalibur to king arthur*lady of the lake
Who had a hit with 'Stand By Your Man'*tammy wynette
Who had a number one hit in 1969 with Something in the Air*thunderclap newman
Who has daughters named Jade, Elizabeth, Scarlett and Georgia*mick jagger
Who has played in the most consecutive baseball games*cal ripken jr
Who has the highest per capital consumption of cheese*france
Who hated mozart with a deadly passion*salieri
Who holds the nhl record for the most goals scored during a regular season*wayne gretzky
Who hosts the monza grand prix*italy
Who invented the cash register in 1879*james ritty
Who invented the difference engine*charles babbage
Who invented the first practical steam engine*thomas newcomen
Who invented the geodesic dome*buckminster fuller
Who invented the pneumatic tyre from a section of garden hose*john dunlop
Who invented the Windows o s*bill gates
Who is Dick Tracy's sweetheart*tess trueheart
Who is famously buried in the churchyard at Bamburgh, Northumberland*grace darling
Who is known as a collector of trivia*spermologer
Who is known as the "George Washington" of South America*simon bolivar
Who is known as the father of genetics*gregor mendel
Who is nick and nora charles' dog*asta
Who is Prime Minister of Australia*john howard
Who is Private Eyes "First Lady of Fleet St"*glenda slagg
Who is robert van winkle*vanilla ice
Who is schroeder's favourite composer*beethoven
Who is the babylonian goddess of love and fertility*ishtar
Who is the Barber of Seville*Figaro
Who is the central figure in Peter C Newmans 'The Establishment Man'*conrad black
Who is the current Secretary of State for Social Security*alastair darling
Who is the Greek messenger god*hermes
Who is the lead singer of 'the doors'*jim morrison
Who is the only singer to have no. 1 hits in the 50's, 60's, 70's, 80's and 90's*cliff richard
Who is the patron saint of foreign missions*st francis
Who is the patron saint of housewives*st anne
Who is the patron saint of lace makers*our lady of loretto
Who is the patron saint of mariners*star of the sea
Who is the patron saint of mathematicians*st hubert
Who is the patron saint of monks*st benedict
Who is the patron saint of nurses*st raphael
Who is the patron saint of organ makers*st genesius
Who is the patron saint of peasants*st lucy
Who is the patron saint of sick poor*st martin de porres
Who is the patron saint of stone masons*st sebastian
Who is the patron saint of surgeons*sts. cosmas & damian
Who is the patron saint of throat*st cecile
Who is the patron saint of women in labor*st anne
Who is the patron saint of writers*st paul
Who is the Prime Minister of France*lionel jospin
Who is the roman counterpart of hermes*mercury
Who is the roman god of light and sky*jupiter
Who kept searching for his long lost salt shaker*jimmy buffet
Who led the mormons to the great salt lake*brigham young
Who lost her sheep*little bo-peep
Who married prince albert*queen albert
Who married the Owl and The Pussycat*the turkey
Who or what was Rosanna Arquette seeking in 1985*susan
Who owned jerusalem before israel*jordan
Who owns: Right Guard deodorant*gillette
Who painted "The Naked Maja"*goya
Who patrols gotham city*batman and robin
Who played 'Banacek' in the 1970's TV series of the same name*george peppard
Who played bonnie to warren beatty's clyde*faye dunaway
Who played detective, Frank Cannon, in the TV series 'Cannon'*william conrad
Who played lestat in 'interview with the vampire'*tom cruise
Who played nick nack & came rolling home*this old man
Who played queen amidala in the latest 'star wars' film*natalie portman
Who played saxophone on "The Girl From Ipanema"*stan getz
Who played the female lead in the Alien films*sigourney weaver
Who played the lead role in the first Tarzan movie*elmo lincoln
who played the male lead in the 1965 film entitled the war lord?*charlton heston
Who played the named character in the following films: Darby's Rangers; Mister Buddwing; and Marlowe*james garner
Who played the respectable hooker in "From here to Eternity"*Donna Reed
Who played the telephone operator on laugh-in*lily tomlin
Who played the title role in the 1921 film 'The Sheik'*rudolf valentino
Who playes Captain Picard in Star Trek: the next generation*patrick stewart
Who plays the part of Inspector Gadget in the film 'Gadget'*matthew broderick
Who presents the radio programme "In the Psychiatrist's Chair"*anthony clare
Who ran unsuccesfully against Regan in 1984?*Walter Mondale
Who recorded "burning bridges" in 1960*jack scott
Who recorded 'a boy named sue'*johnny cash
Who recorded the 1996 alburn, "Older"*george michael
Who recorded the 1997 album "Flaming Pie"*paul mccartney
Who recorded the album "Get Lucky" in 1982*loverboy
Who recorded the album "wish you were here" in 1975*pink floyd
Who released 'time, love and tenderness' in 1981*michael bolton
Who released the No.1 hit single 'Barbie Girl' in October 1997*aqua
Who ruled the seas in Greek mythology*poseidon
Who said 'et tu brute'*julius caesar
Who said 'ronald reagan doesn't dye his hair; he bleaches his face'*johnny
Who said that all matter comes from fire, water, earth & air*aristotle
Who said, ich bin ein Berliner*john f kennedy
Who sailed to the Antarctic in the ship Discovery*scott amundsen
Who sailed to the new world in 'the mayflower'*pilgrims
Who sang 'another one bites the dust'*queen
Who sang 'foolish games'*jewel
Who sang 'friends in low places' and 'thunder rolls'*garth brooks
Who sang 'think' in the original 'blues brothers' film*aretha franklin
Who sang about desmond and molly jones*beatles
Who sang about the fall of man in 'the tall oak tree'*dorsey burnette
Who sang the song from the Disney movie 'Can You Feel The Love Tonight*elton john
Who sang the title theme of the James Bond film A View to a Kill*duran duran
Who sent the brief message "i came, i saw, i conquered"*julius caesar
Who signed the Emancipation Proclamation?*Abraham Lincoln
Who signed the USA for Africa poster with his thumbprint*stevie wonder
Who sought to create the great society*lyndon johnson
Who starred in the film 'the man with two brains'*steve martin
Who started the dragonlance series*margaret weiss and tracy hickman
Who stood at the top with "stand by your man*tammy wynette
Who taught alexander the great*aristotle
who was "bonnie prince charles"?*charles edward stuart
Who was 'too sexy for his shirt'*right said fred
Who was a member of 'crosby, stills and nash' and 'the hollies'*graham nash
Who was Alexander the Great's father*phillip ii
Who was assassinated on november 22, 1963 in dallas*president john f kennedy
Who was born on krypton*superman
Who was British Prime Minister at the outbreak of WWI*herbert asquith
Who was Canadian parliaments first Inuk member*peter ittinuar
Who was captain of the Titanic*edward smith
Who was defendant in the so called "monkey trial"*john t scopes
Who was dictator of Spain from 1937 to 1975*francisco franco
Who was Jack the Ripper's first victim*mary ann nichols
Who was john reid*lone ranger
Who was king arthur's father*uther pendragon
Who was king of macedonia from 336 to 323 b.c*alexander the great
Who was known as the maid of orleans*joan of arc
Who was married to Francis II, Lord Darnley and The Earl of Bosworth*mary, queen of scots
Who was marshall james butler "wild bill" hickock's sidekick*jingles
Who was Michelle's first boyfriend on Full House?*Howie
Who was Mr. Wizard*Don Herbert
Who was named Chairman of the U.S. Federal Reserve Board by Ronald Reagan in 1987, a post he still (February '99) holds*alan greenspan
Who was Pope during World War II*pius xii
Who was responsible for driving the english out of scotland in 1297*william
Who was responsible for the infamous assination attempt on then President Reagan?*John Hinkley Jr.
Who was ronald reagan's first wife*jane wyman
Who was shot as he left the Washington Hilton in 1981*ronald reagan
Who was Tasmania's famous swashbuckler*errol flynn
Who was the 10th president of the U S*john tyler
Who was the 16th president of the united states*abraham lincoln
who was the 16th president of the united states?*abraham lincoln
Who was the 26th president of the U S*theodore roosevelt
Who was the alter ego of 'the incredible hulk'*dr david banner
Who was the defeated Socialist Prime Minister in the Spanish General Election of March 1996*felipe gonzalez
Who was the Egyptian god of the Nile, depicted in human form with a beard, large belly, & a crown of aquatic*hapi
Who was the famous individual who originated the catch phrase "Just Say No"*Nancy Reagan
Who was the first black actress to win an oscar*hattie macdaniel
Who was the first British Prime Minister, although he did not use the title*sir robert walpole
Who was the first Briton to win the Nobel Prize for Literature*rudyard kipling
Who was the first driver to wear a helmet in the indy 500*eddie rickenbacker
Who was the first female prime minister of india*indira gandhi
Who was the first lady to have made the "old blue dress" she wore to an inauguration*rosalynn carter
Who was the first person to swim the english channel*captain matthew webb
Who was the first president born in a hospital*jimmy carter
Who was the first president of the Royal Academy*sir joshua reynolds
Who was the first woman to lead a British trade union*Brenda Dean
Who was the founder of microsoft*bill gates
Who was the French sculptor of the Statue of Liberty*frederic bartholdi
Who was the greek goddess of spring*persephone
Who was the Greek philosopher who decided he'd rather drink hemlock than deny his beliefs*Socrates
Who was the king of Judah (800-783 bc)?*Amaziah
Who was the last king of Troy*priam
Who was the last president of the U S, as of 1998, to die in office*john kennedy
Who was the leader of the bad guys on Hulk Hogan's Rock N Wrestling that annoyed Hulk Hogan and his freinds?*Rowdy Roddy Piper
Who was the leader of the good Transformers?*Optimus Prime
Who was the lone ranger's indian companion*tonto
Who was the longest serving president in French history*francois mitterand
Who was the losing Republican candidate in the 1964 U.S. Presidential Election*barry goldwater
Who was the male star of the film Fatal Attraction*michael douglas
who was the norse goddess of lust and fertility?*freya
Who was the only astronaut to lose his spacecraft*gus grissom
Who was the only pope born in England*adrian iv
Who was the only president born in Illinois, the land of lincoln*ronald reagan
Who was the only President of the Confederate States of America*jefferson davies
Who was the second king of israel*david
Who was the shortest ever mature human?*Gul Mohammed
Who was the sun king*louis xiv
Who was the villain in 'star wars'*darth vader
Who was ulysses' son, who grew to manhood in his absence*telemachus
Who was Vice President to Jimmy Carter, and the Democratic nomination for the presidency in 1984*walter mondale
Who was William Claude Dukenfield better known as*W C Fields
Who was with patricia hearst the night she was kidnaped*steven weed
Who was world champion in boxing from 1952-1962*archie moore
Who went on to become an Eastern Communist leader after working as a pastry chef at London's Carlton Hotel*Ho Chi Minh
Who were the guests on Johnny Carson's final tonite show*bette midler and
Who were the legendry founders of Rome*romulus and remus
Who won the 1995 rugby world cup*south africa
Who won the Oscar for Best Director for the 1988 film "Rainman"*barry levinson
Who wore a cabbage leaf under his cap*babe ruth
Who worked for dr zorba*ben casey
Who wrote "Death of a Salesman" in 1949*arthur miller
Who wrote "Farwell to Arms"*ernest hemingway
Who wrote "The Corn is Green"*emlyn williams
Who wrote "The Outcast of the Islands"*joseph conrad
who wrote "titus groan"?*mervyn peake
who wrote 'a clockwork orange'?*anthony burgess
Who wrote 'alice in wonderland'*lewis carroll
Who wrote 'born free', 'living free' and 'forever free'*joy adamson
Who wrote 'la traviata'*guiseppe verdi
Who wrote 'the female eunuch'*germaine greer
Who wrote 'valley of the dolls'*jacqueline susann
Who wrote Auld Lang Syne*robert burns
Who wrote David Copperfield*charles dickens
Who wrote Moll Flanders*daniel defoe
Who wrote most of the new testament books*paul
Who wrote MASH*richard hooker
Who wrote the 'Noddy' books*enid blyton
Who wrote the 'noddy' books*enid blyton
Who wrote the book on which the Oscar winning film 'The Godfather' was based*mario puzo
Who wrote the children's story Badjelly the Witch*spike milligan
Who wrote the classic thriller 'The Birds'*Alfred Hitchcock
Who wrote the Father Brown crime stories*gk chesterton
Who wrote the hit musical West Side Story*leonard bernstein
Who wrote The Ipcress File*len deighton
Who wrote the novel "Slaughterhouse Five"*kurt vonnegut jr
Who wrote the novel ' Anna of the Five Towns'*arnold bennett
Who wrote the novel Enigma in 1995, about the wartime German coding machines*robert harris
Who wrote the opera "the trojans"*hector berlioz
Who wrote the opera 'i pagliacci'*ruggiero leoncavallo
Who wrote the Robocomm computer program*dan parsons
Who wrote the song 'Anything Goes'*cole porter
Who wrote the song 'do they know it's christmas' with midge ure*bob geldof
Who wrote the song of songs*solomon
Who wrote the story of "the nutcracker"*eta hoffmann
Who wrote the supernatural tale The Turn of the Screw*henry james
Who wrote To Kill A Mockingbird*harper lee
Who wrote Vanity Fair*william thackeray
Who's Best of Album is called Paint the Sky With Stars*enya
Who's the leading rebounder in NBA playoff history*bill russell
Who, in 1655, discovered Saturn's rings*christiaan huygens
Who, in 1874, painted the picture called La Loge*auguste renoir
Who, in egyptian mythology, is the god of the dead*aker
Whose 31st and 38th Symphonies are the Paris and the Prague*mozart
Whose car, when found in Dallas in 1963, contained brass knuckles, a pistol holder, and a newspaper detailing JFK's motorcade route?*Jack Ruby
Whose girl friend was Virginia Hill*bugsy siegel
Whose grandson got the first phone call from a commercial cellular system, in 1983*alexander graham bell's
Whose hamburger patties weigh 1.6 oz*mcdonald's
Whose last words were reportedly, 'I shall hear in heaven!'*beethoven
Whose life story is titled 'fly me, i'm freddie!'*freddie laker
Whose name did God change to Israel*jacob
Whose novels include 'The Cement Garden' and 'Comfort of Stangers'*ian mcewan
Whose novels include 'The Ice-Cream Wars' and 'Brazzaville Beach'*william boyd
Whose only loss in 1983 was to kathy horvath*martina navratilova
Whose patron is Holy Spirit*understanding
Whose patron is St Barbara*artillery
Whose patron is St Christopher*truck Drivers
Whose patron is St Dymphna*runaways
Whose patron is St Francis de Sales*authors
Whose patron is St Francis de Sales*teachers
Whose patron is St Matthew*stockbrokers
Whose patron is St Nicholas*sicily
Whose patron is St Paul*authors
Whose patron is St Peter*stationers
Whose patron is St Rose of Lima*vanity
Whose patron is St Stephen*austria
Whose patron is St William*adopted children
Whose recent books include 'Crisis Four' and 'Firewall'*andy mcnab
Whose rule is used to solve simultaneous linear equations by using determinants*cramer
Whose single season strikeout record did Nolan Ryan beat by one*sandy koufax
Wide muscular partition separating the thoracic, or chest cavity, from the abdominal cavity*diaphragm
Wild Australian dog*dingo
William Golding won the Nobel Prize for literature in which year*1983
Winston churchill resigned from office in 1954, 1955 or 1956*1955
With what branch of medicine is mesmer associated*hypnotism
With what country is prince rainier iii associated*monaco
With what is 'Grand Marnier' flavoured*orange
With which island is the puffin associated*lundy island
With which musical instrument is Dizzy Gillespie chiefly associated*trumpet
With which sport is Chris Evert identified*tennis
With which sport is Willie Mays associated*baseball
Woollen covering for head and neck*balaclava helmet
Words containing for: many trees*forest
Words containing pot or pan: Tyrant*despot
Workshop for casting metal*foundry
Wreath of flowers used as a decoration*garland
Wwhat is the name of Mulder and Scully's supervisor on the X-files?*Walter Skinner
X-Men Comics: Gahck battled wolverine here*savage land
X-Men Comics: Wolverine and SpiderMan discovered the identity of HobGoblin Here*berlin
Year in which the Battle of Balaklava took place*1854
You have to run 360 feet if you hit a ----------*home run
Young man paid by older woman to be escort or lover*gigolo
----------, the story of prize fighter Jake Lamotta, packs a real punch*Raging Bull
A bell tower, usually not actually attached to a church*campanile
A continuous aisle in a building, especially around the apse in a church*ambulatory
A curved structure used to span an opening*arch
A curved triangle at the corners of a square or polygonal room, used at the opening of a dome*pendentive
A cylindrical vertical support usually consisting of a base, shaft and capital*column
A flattened, shallow column or pier projecting from a wall. It usually has a base, shaft, and capital but is decorative rather than structural*pilaster
A horizontal projection, such as a balcony or beam, supported at one end only*cantilever
A method of construction in which vertical beams are used to support a horizontal beam*post and lintel
A movement that developed in the 1920s, characterized by a regularized surface, a lightening of mass, and  often large expanses of glass*international style
A multistoried building, typically Asian, forming a tower with upward curving roofs over the individual stories*pagoda
A part of a church or a separate building, often octagonal or round, in which baptisms take place*baptistery
A passageway of a Christian Church or a Roman basilica running paralell to the nave, separated from it by an arcade or colonnade*aisle
A projecting support built into or against the external wall of a building, typically used in Gothic buildings. A flying -------- is an arch that transfers the thrust of a vault to a lower support*buttress
A roofed gallery with an open arcade or colonnade on at least one side*loggia
A row of columns, usually equidistant, supporting a beam or entablature*colonnade
A row of windows in the upper part of a wall, especially in a church, to admit light below*clerestory
A semicircular area at the end of a church; in most churches it contains the altar*apse
A series of arches supported by columns or piers, or a passageway formed by these arches*arcade
A slender, lofty tower with balconies, attached to a Muslim mosque*minaret
A small structure on top of a dome, tower or roof often open to admit light below*lantern
A spiral scroll used on Ionic and Corinthian capitals*volute
A spout placed on the roof gutter of a Gothic building to carry away rainwater, usually carved in the shapes of fanciful animals and grotesque beasts*gargoyle
A square of rectangular area in a church between the apse and the crossing*choir
A stone slab at the top of a classical column aiding the support of the architecture*abacus
A structure that forms the arms of a cross-shaped church*transept
A structure usually attached to a building, such as a porch, consisting of a roof supported by piers or columns*portico
A style of English architecture prevalent from 1485-1558 transitional between Gothic and Palladian, with emphasis on country manors*Tudor
A style that emerged in the 1970s characterized by references to and evocations of past architectural styles, particularly the classical tradition. It is frequently colorful and wittily ornamentive*postmodernism
A style that flourished in the seventeenth and early eighteenth centuries, characterized by exuberant decoration, curvaceous forms, and a grand scale generating a sense of movement; later developments within the movement show more restraint*baroque
A tall, tapering, four-sided stone shaft with a pyramidal top*obelisk
A tall, tapering, pointed roof on a tower, as in the top of a steeple*spire
A vaulted roof of circular or polygonal shape*dome
An arched brick or stone ceiling or roof*vault
An ornamented canopy over an altar, tomb or throne*baldachin
An upright masonry support*pier
Any important face of a building, usually the principal front with the main entrance*facade
In German Romanesque, a monumental entrance to a church consisting of towers, with a chapel above*westwork
In a church, the area where the transept and the nave intersect, usually emphasized by a dome or a tower*crossing
In a classical building, the triangular gable between the horizontal entablature and the sloping roof; in general, and architechtural feature over a door or window*pediment
In 1958 the Nobel prize in literature was given to Boris Leonidovich Paternak for his important achievement both in contemporary lyrical poetry and in the field of the great ... ?*Russian epic tradition
In a roman basilika, the central aisle. In a church, the main section extending from the entrance to the crossing*nave
In an ancient Roman house, a central room open to the sky, usually having a pool for the collection of rainwater. In churches, a front courtyard*atrium
In ancient Assyria and Babylonia, a tower in the shape of a stepped pyramide. It formed the base of a temple*ziggurat
In ancient Roman architecture, a large oblong building, generally with double columns and a semicircular apse at one end. In Christian architecture, a church with a nave, apse, and aisles*basilica
In religious institutions, a courtyard with covered walks*cloister
Moldings and ornamentation projecting from the surface of a wall*relief
Ornament of ribs, bars, etc. in panels or screens, as in the upper part of a Gothic window*tracery
Stones hewn, squared, and smoothed for use in building, as distinguished from rough building stones*ashlar
Sun-dried brick used in places with warm, dry climates, such as Egypt and Mexico; also, the structures built out of these bricks*adobe
The elevated stronghold in ancient Greek cities*acropolis
The lowest part of an entablature resting on the capital of a column*architrave
The measurement by which parts of a building are related to one another, for example, the diameter of a column*module
The middle part of an entablature, often decorated with sculpture*frieze
The pointed arch used in Gothic architecture*ogive
The story above the cornice of a building*attic
The style of this school, founded in Germany by Walter Gropius in 1919, emphasizing simplicity, functionalism and craftsmanship*Bauhaus
The transverse entrance hall of a church*narthex
The triangular area between the two sides of two adjacent arches*spandrel
The upper horizontal part of a classical order, between a capital and the roof; it consists of the architrave, frieze, and cornice*entablature
The upper part of an entablature, extending beyond the frieze; also, ornamental molding projecting along the top of a building or wall*cornice
Friendship is a single soul, dwelling in ... ?*two bodies
A European movement beginning in France. Gothic sculpture emerged c. 1200, Gothic painting later in the thirteenth century. The artworks are characterized by a linear, graceful, elegant style more naturalistic than that which had existed previously in Europe*gothic
A European movement of the late eighteenth to mid-nineteenth century. In reaction to neoclassicism, it focused on emotion over reason, and on spontaneous expression*romanticism
A European style developed in France in the late eleventh century. Its sculpture is ornamental, stylized and complex*romanesque
A European style of the late eighteenth and early nineteenth centuries. Its elegant, balanced works revived the order and harmony of ancient Greek and Roman art*neoclassicism
A Russian abstract movement originated by Malevich c. 1913. It was characterized by flat geometric shapes on plain backgrounds and emphasized the spiritual qualities of pure form*suprematism
A decorative art movement that emerged in the late nineteenth century.  Characterized by dense assymmetrical ornamentation in sinuos forms, it is often symbolic and of an erotic nature*art noveau
A figurative movement that emerged in the United States and Britain in the late 1960s and 1970s. The subject matter, usually everyday scenes, is portrayed in an extremely detailed, exacting style. It is also called superrealism, especially when referring to sculpture*photorealism
A group of American painters who united out of opposition to academic standards in the early twentieth century*The eight
A group of English painters formed in 1848. These artists attempted to recapture the style of painting preceding Raphael. They rejected industrialized England and focused on painting from nature, producing detailed, colorful works*Pre-Raphaelite Brotherhood
A late-nineteenth-century French school of painting. It focused on transitory visual impressions, often painted directly from nature, with an emphasis on the changing effects of light and color*impressionism
A method of painting developed by Seurat and Signac in the 1880s. It used dabs of pure color that were intended to mix in the eyes of viewers rather than on the canvas. It is also called divisionism or neoimpressionism*pointillism
A movement in American painting and sculpture that originated in the late 1950s. It emphasized pure, reduced forms and strict, systematic compositions*minimalism
A movement in European painting in the seventeenth and early eighteenth centuries, characterized by violent movement, strong emotion, and dramatic lighting and coloring*baroque
A movement of the 1920s and 1930s that began in France. It explored the unconscious, often using images from dreams. It used spontaneous techniques and featured unexpected juxtapositions of objects*surrealism
A movement of the 1960s and 1970s that emphasized the artistic idea over the art object. It attempted to free art from the confines of the gallery and the pedestal*conceptual art
A movement that began in Britain and the United States in the 1950s. It used the images and techniques of mass media, advertising, and popular culture, often in an ironic way*pop art
A movement, c. 1915-23, that rejected accepted aesthetic standards. It aimed to create antiart and nonart, often employing a sense of the absurd*dadaism
A painting movement that flourished in France in the 1880s and 1980s in which subject matter was suggested rather than directly presented. It featured decorative, stylized, and evocative images*symbolism
A russian abstract movement begun in the early twentieth century. It employs an analytic vision based on fragmentation and multiple viewpoints*cubism
A style, c. 1520-1600, that arose in reaction to the harmony and proportion of the High Renaissance. It featured elongated, contorted poses, crowded canvases, and harsh lighting and coloring*mannerism
A technique in abstract painting developed in the 1950s. It focuses on the lyrical effects of large areas of color, often poured or stained onto the canvas*color field painting
A termed coined by British art critic Roger Fry to refer to a group of nine-teenth century painters, who were dissatisfied with the limitations of impressionism. It has since been used to refer to various reactions against impressionism, such as fauvism and expressionism*postimpressionism
An abstract movement in Europe and the United States, begun in the mid-1950s, based on the effect of optical patterns*op art
An eighteenth-century European style, originating in France. In reaction to the grandeur and massiveness of the baroque, it employed refined, elegant, highly decorative forms*rococco
An italian movement c.1909-1919. It attempted to integrate the dynamism of the machine age into art*futurism
Artwork, usually paintings, characterized by a simplified style, nonscientific perspective, and bold colors. The artists are generally not professionally trained*naive art
Design style prevalent during the 1920s and 1930s, characterized by a sleek use of straight lines and slender forms*art deco
From the Hebrew word for 'prophet'. A group of French painters active in the 1890s who worked in a subjective, sometimes mystical style, stressing flat areas of color and pattern*Nabis
From the french word 'fauve', meaning 'wild beast'. A style adopted by artists associated with Matisse, c. 1905-1908. They painted in a spontaneous manner, using bold colors*fauvism
Group of American artists from 1908 to 1918. Their work featured scenes of urban realism*Ash Can School
In a general sense, refers to objective representation. More specifically, a nineteenth century movement, especially in France, that rejected idealized academic styles in favor of everyday subjects*realism
Meaning 'rebirth' in french. Refers to Europe c. 1400-1600. The style began in Italy and stressed the forms of classical antiquity, a realistic representation of space based on scientific perspective, and secular subjects*Renaissance
Movement in painting, originating in New York City in the 1940s. It emphasized spontaneous personal expression, freedom from accepted artistic values, surface quallities of paint, and the act of painting itself*abstract expressionism
Referring to the principles of Greek and Roman art of antiquity with its emphasis on harmony, proportion, balance, and simplicity. In a general sense, it refers to art based on accepted standards of beauty*classicism
Refers to art that uses emphasis and distortion to communicate emotion. More specifically, it refers to early twentieth-century northern European art, especially in Germany c. 1905-23*expressionism
Works of a culturally homogenous people without formal training, generally according to regional traditions and involving crafts*folk art
A band of painted or sculpted decoration, often at the top of a wall*frieze
A composition made of cut and pasted pieces of materials, sometimes with images added by the artist*collage
A flat board used by a painter to mix and hold colors, traditionally oblong, with a hole for the thumb; also, a range of colors used by a particular painter*palette
A large painting or decoration done on a wall*mural
A method of producing images or letters from sheets of cardboard, metal, or other materials from which forms have been cut away*stenciling
A method of watercolor painting, but prepared with a more gluey base, producing a less transparent effect*gouache
A painting or drawing executed in a single color*monochrome
A painting technique using pigments mixed with egg yolk and water. It produces clear, pure colors*tempera
A print made by carving on a wood block, which is then inked and printed*woodcut
A printing process in which ink impressions are taken from a flat stone or metal plate prepared with a greasy substance, such as an oily crayon*litography
A realistic style of painting in which everyday life forms the subject matter, as distinguished from religious or historical painting*genre painting
A representation of a human or an animal form*figure
A single print made from a metal or glass plate on which an image has been represented in paint, ink, etc*monotype
A soft, subdued color; a drawing stick made of ground pigments, chalk, and gum water*pastel
A technique of engraving, using a sharp-pointed needle, that produces a furrowed edge resulting in a print with soft, velvety lines*drypoint
An artwork humoously excaggerating the qualities, defects, or pecularities of a person or idea*caricature
An etching tecnique in which a solution of asphalt or resin is used on the plate. It produces prints with rich, gray tones*aquatint
Ground chalk or plaster mixed with glue, used as a base coat for tempera and oil painting*gesso
In painting, a thin layer of translucent color*wash
In painting, a work made of several panels or scenes joined together. A diptych has two panels; a triptych, three*polyptych
In painting, the degree of lightness or darkness in a color*values
In sculpting, the cutting of a form from a solid, hard material such as stone or wood, in contrast to the technique of modeling*carving
In sculpture, the building up of form using a soft medium such as clay or wax, as distinguished from carving. In painting and drawing, using color and lighting variations to produce a three-dimensional effect*modeling
In sculpture, the projection of an image or form from its background. Sculpture formed in this manner is described as high relief or low relief (bas-relief), depending on the degree of projection. In painting or drawing, the apparent projection of parts conveying the illusion of three dimensions*relief
Meaning 'fool the eye' in french. In painting, the fine, detailed rendering of objects to convey the illusion that the painted forms are real and three-dimensional*trompe l'oeil
Meaning 'fresh' in italian. The technique of painting on moist lime plaster with colors ground in water*fresco
On a represented form, a point of most intense light*highlight
Paint applied very thickly. It often projects from the picture surface*impasto
Painting in which natural scenery is the subject*landscape
Reducing or distorting in order to represent three-dimensional space as perceived by the eye, according to the rules of perspective*foreshortening
The effect of the harmony of color and values in a work*tone
The rendering of light and shade in painting; the subtle graduations and marked variations of light and shade for dramatic effect*chiaroscuro
The representation of inanimate objects in painting, drawing or photography*still life
The technique of producing printed designs through various methods of incising on wood or metal blocks, which are then inked and printed*etching
The visual and tactile quality of a work based on the particular way the materials are handled; also, the distribution of tones or shades of a single color*texture
Water-soluble paint made from pigments and a plastic binder?*acrylic
Deep down I am fairly ... ?*superficial
Never confuse motion with ... ?*action
I do not sing, I do not dance, and I don't say ... ?*sir
I am not an actor. I am a ... ?*phenomenon
Americans usually get it right, after trying ... ?*everything else first
The first half of our lives is ruined by our parents, and the second half ... ?*by our children
A Bohemian folk dance in duple time with a hop on the fouth beat. It became a popular ballroom dance in the mid-nineteenth century*polka
A Japanese dance drama featuring stylized narrative choreographic movements*kabuki
A Sevillian gypsy dance, possibly originating in India, also with Moorish and Arabian influences, originally accompanied by songs and clapping and later by the guitar, and characterized by its heelwork*flamenco
A Spanish dance in 3/4 time or 3/8 time with castanets*cachucha
A ballet bow or curtsy in which one foot is pointed in front and the body leans forward*reverence
A ballet in which the women wear white tutus, such as the second and fourth acts of Swan Lake*ballet blanc
A ballet movement in which the dancer repeatedly crosses his or her legs in the air*entrechat
A ballet with a plot, usually tragic, to bring dramatic coherence to the performance of ballet*ballet d'action
A basic movement in the technique of Martha Graham, based on breath inhalation and exhalation*contraction
A beating movement of the legs*battement
A bending of the knees in any of the five positions*plie
A dance for two, usually a woman and a man. In its traditional form, it begins with an entree and adagio, followed by solo variations for each dancer, and a coda*pas de deux
A dance with a fast or moderate tempo*allegro
A grave, processional court dance popular in the sixteenth and seventeenth centuries*pavane
A jump in which the legs open in second position in the air, resembling a scissors*ciseaux
A leap from one leg to the other in which one leg is thrown to the side, front or back*jete
A lively Spanish dance in triple time performed with castanets or tambourines*fandango
A lively social dance popular during the 1930s; it originated at the Savoy Ballroom in Harlem in 1928, where it was known as the Lindy*jitterbug
A male dancer who performs the 'princely' roles of the classical ballet, such as the Prince in Swan Lake*danseur noble
A plotless work composed of pure dance movements, although the composition may suggest a mood or subject*abstract dance
A polish national dance in triple time with an accent on the second beat, characterized by proud bearing, clicking of heels, and holubria, a special turning step*mazurka
A position on the tip of the toes*point
A series of small, fast steps executed with the feet very close together*pas de bourree
A sliding step in which one foot 'chases' and displaces the other*chasse
A slow and graceful dance, the most popular dance of the eighteenth century, characterized by symmetrical figures and elaborate curtsys and bows*minuet
A social dance in 3/4 time that became widely popular in the nineteenth century. It developed from the Landler, a German-Austrian turning dance*waltz
A social dance in 2/4 time, which after originating in Spain, developed in Argentina, where it was influenced by black dance style and rhytm*tango
A social dance of American origin in duple time*fox-trot
A social dance popular in the nineteenth century. It was a square dance in five sections, each in a different time*quadrille
A solemn court dance usually in duple time, popular in the fifteenth and sixteenth centuries*basse danse
A spectacular movement in which the dancer propels himself or herself around a supporting leg with rapid movements of the other leg while remaining in a fixed spot*fouette en tournant
A step that rocks from one foot to the other, usually in � time*balance
A turn on one leg, with the toe of the other leg touching the knee of the turning leg*pirouette
A turn while jumping straight up in the air*tour en l'air
An American folk dance with an even number of couples forming a square, two lines, or a circle. The dance is comprised of figures announced by a caller*square dance
An English folk dance that appeared in the fifteenth century, in which dancers wore bells on their legs and characters included a fool, a boy on a hobby horse, and a main in blackface*morris dance
An african-american dance in which couples strut and compete with high kicks and fast steps*cakewalk
An unfolding of the leg in the air*developpe
Any dance to slow music; also, part of the classical pas-de-deux in ballet*adagio
Any solo performance in a ballet*variation
Catlike leap in which one foot follows the other into the air, knees bent; the landing is in the fifth position*pas de chat
In ballet, a bend from the waist to the side or to the back*cambre
In ballet, a closed position of the feet*ferme
In ballet, a gliding step which usually connects two steps*glissade
In ballet, a jump from one to both feet, usually landing in fifth position*assemble
In ballet, a jump off one foot that is 'broken' by a beating of the legs in the air*brise
In ballet, a leap in which the lower leg beats against the upper one at an angle, before the dancer lands again on the lower leg*cabriole
In ballet, a lowering of the body by bending the knee*fondu
In ballet, a pose in which one leg is raised in back or in front with knee bent, usually with one arm raised*attitude
In ballet, a position of the arms above the head*en haut
In ballet, a position of the body at an oblique angle and partly hidden*efface
In ballet, a position with one leg extended at an oblique angle while the body is also at an oblique angle*ecarte
In ballet, a position with the body at an oblique angle and the working leg crossing the line of the body*croisee
In ballet, a rising with a spring movement to point or demi-point*releve
In ballet, a slow turn of the body on the whole foot*promenade
In ballet, a step done off the ground*en l'air
In ballet, an elongated line; in particular, the horizontal line of an arasbesque with one arm stretched front and the other back*allonge
In ballet, an open position of the feet*ouvert
In ballet, leaning forward*penche
In ballet, low, as in placement of arms*en bas
In ballet, shifting weight from one foot to the other*degage
In ballet, the ability of a dancer to remain suspended in air during a jump; elasticity in jumping*ballon
In ballet, the position of the arms*port de bras
In ballet, the position of the torso from the waist up*epaulement
In ballet, the third and final part of the classical pas de deux*coda
In ballroom dance, a characteristic figure that remains constant*basic movement
Originating around 1830 as a social dance, by 1844 it had become a raucous dance performed in French music halls*cancan
Popular social dance during the eighteenth century; done in rows or circles, it may have derived from English country dancing*contredanse
Principial male dancer*premieur danseur
Social dances usually performed by couples, including the fox-trot, waltz, tango, rumba and cha cha*ballroom dances
Spectacles for entertainment, usually with allegorical or mythological themes, performed by the aristocracy in the sixteenth and seventeenth centuries, combining music, recitatives and mime*court ballet
Standard Italian dances and their music of the fifteenth and sixteenth centuries*ballo
Stepping directly onto the point of a foot*pique
Steps performed on the floor. It is the opposite of en l'air*par terre
The members of a ballet company who do not perform solo*corps de ballet
Traditional English dance in which dancers form two facing lines*country dance
What are the two basic concepts in economics ?*Wealth and Welfare
'Pareto efficiency' (after Vilfredo Pareto) is a situation in which it is not possible to make someone better off without making someone else ... ?*worse off
A decision to satisfy one set of wants necessarily means sacrificing some other set: this sacrifice is called by economics the ... ?*Opportunity cost
A firm is a decision-making production unit which transforms resources into goods and services which are ultimately bought by consumers, the government and ... ?*other firms
A formation resulting from mergers or take-overs involving firms whose activities are not directly related could be called a ... ?*conglomerate
A pure public good is a good or a service, such as defense, the consumption of which by one person does not reduce its benefit to ... ?*others
According to one argument economists are able to give advice on issues related to economic efficiency, but equity (fairness) considerations are outside the purview of economics and should be left to ... ? (3 other groups)*Philosophers, politicians and social reformers
Although economists tend to concentrate on the relationship between demand and price, it is sometimes useful to consider the relationship between demand and income, ceteris paribus. Represented graphically, such a relationship, named after the economist Ernst Engel is called an ... ?*Engel curve
An individual's demand for a commodity may be defined as the quantity of that commodity that the individual is willing and able to buy during a given ... ?*time period
As the price of one of two comparable commodities falls, the price of the other becomes relatively more expensive. The consumer is therefore induced to buy (choose) the first. This is called the ... of the price change*substitution effect
Besides land, labour and capital, there is a fourth factor sometimes added to the main factors of production. This fourth factor is ... ?*enterprise
Differentiated products are products which are very similar, but not ... ?*identical
Economists analyse the relationship between a consumer's demand for a specific good and the price of a specific good by assuming that all other influencing factors remain unchanged. This is the important assumption called ... ?*ceteris paribus
Economists have managed to separate the problems of an efiicient allocation of resources from the controversial question of the distribution of income and wealth. The latter is concerned with ... ?*value judgements
Economists who believe that utility (the satisfaction deriving from consuming a certain good) can be measured in units, as if it were a physical commodity, are known as ... ?*cardinalists
Economists who oppose the view that utility (the satisfaction deriving from consuming a certain good) can be measured cardinally have become known as ... ?*ordinalists
In economic theory, it is often found convenient to assume that there are 'constant returns to scale' in production. What this means is that when a producer employs more labour and more capital, his output increases ... ?*proportionally
In the 1880s the German economist Adolph Wagner advanced his law of ever rising ... ?*public expenditures
In the long-run, all factors of production are variable. Firms wishing to maximise their profits, therefore, will attempt to produce their chosen output by employing combinations of capital, labour and land which minimise their ... ?*production costs
Macroeconomics concerns itself with large aggregates, particularly for the economy as a whole. It deals with the factors which determine national output and employment, the general price level, toal spending and saving, total imports and exports, and the demand and supply of ... ?*money
Most economics stress the idea of a 'state of rest', in which no economic forces are being generated to change the situation. This is also called ... ?*Equilibrium
One of the three major categories of inputs into productions processes is capital. Capital consists of goods which are not for current consumption, buit which will assist consumer goods to to be produced in the ... ?*future
One of the three major categories of inputs into productions processes is labour. Labour includes all the ... which are used in production*human attributes
One of the three major categories of inputs into productions processes is land. Land includes all the ... which are used in production*natural resources
Over the past twenty-five years or so, economists have divided their subject matter into two main branches. Which ?*Microeconomics and macroeconomics
Saving is that part of disposable income which is not spent in the current period. It follows that disposable income minus saving equals ... ?*consumption
The 'elasticity of demand' is a measure of the extent to which the quantity demanded of a good responds to changes in one of the influencing factors. The main measures are the price, income and ... elasticity of demand*cross
The 'law of demand' states that a rise in the price of a good leads to a fall in the total quantity demanded. A fall in the price of a good leads to a rise in the total ... demanded*quantity
The 'law of diminishing returns' states that as additional units of a variable factor are added to a given quantity of fixed factors, with a given state of technology, the average and marginal products of the variable factor will eventually ... ?*decline
The 'theory of revealed preference' is based on the reasonable proposition that a consumer will actually choose to consume that collection of goods that he ... ?*prefers
The basic economic problem is that of allocating scarce resources among the competing and virtually limitless wants of ... ?*consumers in society
The combining of firms that produce at a similar stage of an industry's production is called a ... ?*horizontal integration
The economist K.Lancaster has argued that it is the characteristics or ... of goods which yield utility to the consumer, rather than the goods themselves*attributes
The equilibrium (state of rest) behaviour of consumers and producers, whether in a single market or in the economy as a whole, is characterised by the fact that there exists no feeling of urgency on the part of buyers and sellers to ... ?*change their behaviour
The hypothesis of diminishing marginal utility states that as the quantity of a good consumed by an individual increases, the marginal utility of the good will eventually ... ?*decrease
The inverse relationship between the price of a commodity and the quantity demanded in the market is summed up in the so-called  ... ?*law of demand
The logical progession from a one-man business is to a ... ?*partnership
The short-run is that period of time over which the input of at least one factor of production cannot be varied. Those factors which can be varied in the short-run (typically labour, raw materials and fuel) are called variable factors; those which cannot be varied (typically capital and land) are called ... ?*fixed factors
The term 'methodology' refers to the way in which economists go about the study of their subject matter. Broadly, they have followed two main lines of approach. Which ?*positive economics and normative economics
There are many different inputs into most production processes. For the purpose of analysis economists typically place each of the many different factor inputs into one of three categories ... ?*Land, labour and capital
There are two major exceptions to the law of demand. These exceptions (inferior goods and luxury items) are called ... ?*Giffen goods and Veblen goods
Traditional economic theory has assumed that the typical firm has a single objective, namely to ... ?*maximise its profits
When a good is consumed, the consumer presumably derives some benefit or satisfaction from the activity. Economists have called this benefit or satisfaction ... ?*utility
When two or more firms in the same industry, but at different stages in the production process, join together, this is an example of ... ?*vertical integration
In a market of equilibrium (state of rest) the price and quantity of a commodity match both consumers and producers expectations and thus there is no discrepancy (conflict) betweeen the actual and desired ...?*prices and quantities
The famous italian economist Vilfredo Pareto (1848-1923) concentrated on the efficiency aspect of welfare because he believed that ... ?*welfare was a highly subjective concept
The word equity means ?*fairness or justice
What are the 3 central economic questions facing all nations ? (concerning what, how and whom)*What goods and services to produce, how to produce them and for whom
Microeconomics is concerned with the behaviour of ... ? (3 groups)*Individual firms, industries and consumers
The most common reason for divorces is ... ?*Men and women
Homosexuality is a sickness, just as are baby-rape or wanting to become head of ...?*General Motors
I hate the giving of the hand unless the whole man ... ?*accompanies it
Within, I do not find wrinkles and used heart, but unspent ... ?*youth
General and abstract ideas are the source of  the greatest ... ?*errors of mankind
Democracy substitutes election by the incompetent many for appointment by the ... ?*corrupt few
He knows nothing and he thinks he knows everything. That points clearly to a ... ?*political career
From the moment I picked up your book until I laid it down, I was convulsed with laughter. Some day I ... ?*intend to read it
I find television very educating. Every time somebody turns on the set ... ?*I go into the other room and read a book
I knew Doris Day before she became a ... ?*virgin
It is even harder for the average ape to believe that he has ... ?*descended from man
Marriage is a wonderful institution, but who would want to live ... ?*in an institution
Thank god that I am still an ... ?*atheist
Allen strikes gold as he examines some typically and neurotic New Yorkers whose lives intertwine*Hannah and her sisters
Angie gets a workout in three separate stories about the effects of that proverbial green-eyed monster*Jealousy
Appealing and intelligent comedy about a highly charged, neurotic woman who's a succesful TV news producer, and her attraction to a pretty-boy anchorman who joins her network - and represents everything she hates about TV news*Broadcast News
Army officer who survies an atomic explosion starts growing; at sixty feet he attacks Las Vegas*Amazing colossal man
Bishops turn into skeletons and the cow wanders into the bedroom in Bunuels first feature, a surrealistic masterpiece coscripted by Salvadore Dali*L'age d'or
Bizarre, sexually-oriented parasites run rampant through dwellers in high-rise apartment building with plenty of gory violence quick to ensue. First major film by cult favorite Cronenberg sets the disgusting pattern for most of his subsequent pictures*They came from within
Bland adaption of F.Scott Fitzgerald's jazz-age novel about a golden boy in Long island society; faithful to the book, and visually opulent, but lacks substance and power*The Great Gatsby
Blockbuster biography of enigmatic adventurer T.E.Lawrence is that rarity, an epic film that is also literate*Lawrence of Arabia
Chaplin attacks the machine age in inimitable fashion, with sharp pokes at other social ills and the struggle of modern-day survival*Modern Times
Director Losey ate watermelon, pickles, and ice cream, went to sleep, woke up, and made this adaption of the comic strip about a sexy female spy*Modesty Blaise
Dull, cheap version of the D.H.Lawrence classsic, with Kristel as the lady and Clay as the lover. Kristel is beautiful but still cannot act*Lady Chatterley's lover
Excellent adaption of Albert Camus' existential novel about a man who feels completely isolated from society*The Stranger
Exciting Raymond Chandler melodrama has Ladd returning from military service to find wife unfaithful. She's murdered , he's suspected in well-turned film*The Blue Dahlia
Family encounters a bigfoot-type monster in the woods and takes it home, thinking it's dead*Harry and the Hendersons
Flat naval comedy set in WW2 with Cooper commanding a dumb crew on the U.S.S. Teakettle. Film debuts for Marvin and Charles Bronson*You're in the Navy now
Formula filmmaking that even bored its intended audience. Cop Eastwood chews stogies for breakfast, while new partner Sheen is a rich kid, who apparently enjoys collecting facial contusions*The Rookie
Gangster Steiger hires 'sensitive' hitman Palance to killl Svenson, but there are complications: Svensson is a pal who once saved Palance's life, and both of them are in love with Turkel*Portrait of a hitman
Genteel, entertaining adaption of Alfred Uhry's stage play about a simple black man who's hired as chauffeur for a cantankerous old Southern woman, and winds up being her most fatihful companion*Driving Miss Daisy
Good-buddy truckdrivers Fonda and Reed battle a rival kingpin's goon who want to force them off the road for good*High-ballin'
Jerry is the poor stepson turned into a handsome prince for a night by fairy godfather*Cinderfella
John and Mary meet, make love, but don't know if the relationship should end right there. Innocuous, uncompelling trifle. Hoffman seems to be sleepwalking; audience may join him*John and Mary
Ladd in his element as a soft-spoken but iron-willed railroad agent whose hot-headed best friend becomes involved in shady dealings*Whispering Smith
Lee dies midway thorugh the production of this karate thriller*Game of death
Lee is suitably cast as sinister Count Drago, who mummifies visitors to his gothic castle. Unexceptional horror fare, notably mainly as Sutherland's film debut in two roles - one as an old lady*Castle of the living dead
Lithgow plays a meek butcher who wrongly believes he has killed his partner after discovering him frozen to death in the freezer. A wonderfully adept cast tries to pull off this black comedy, but the script knocks their efforts out cold*Out Cold
Loud, boring, interminable tale of skier 'John' who romances skier 'Suzy' on the slopes*Fire and ice
Lowbudget garbage about families developing from people who remained on island after the Mutiny on the Bounty*The Women of Pitcairn Island
Modest little black comedy about a hit-man whose current job is mucked up by an intrusive stranger*Buddy Buddy
Moody version of Herman Melville sea classic, with Peck lending a deranged dignity to the role of Captain Ahab*Moby Dick
Occult expert called in by San Francisco police in connection with series of weird murders. Intricate plot and and exceptional time period blending makes this a one-of-a-kind movie*Dark Intruder
Our candidate for the best Hollywood movie of all time*Casablanca
Paddy Chayefsky's outrageous satire on television looks less and less like fantasy as the years pass*Network
Phobic patient Murray pursues pompous psychiatrist Dreyfuss to his vacation retreat, where he ingratiates himself with the shrink's family - and drives the doctor crazy*What about Bob?
Pleasant if pointless fable about a stressed-out nerd who learns he has six months to live, and accepts a millionaire's offer to enable him to live like a king, so long as he jumps into a volcano at the end of his vacation. Unfortunately the story also takes a dive*Joe versus the Volcano
Reed is a grotesquely ugly podiatrist who drinks a poison to commit suicide, instead turns into a handsome murderer*Dr. Heckyl and Mr. Hype
Repressed homosexual doctor jeopardizes 8-year-marriage by coming out of the closet with sexually carefree novelist*Making love
Saucy, sexy single mother of two is a source of constant embarassment to her teenage daughter, who's trying to deal with her own sexual awakening - and not having an easy time of it*Mermaids
Sexually precocious Lyon becomes involved with stolid professor Mason, and bizarre Sellers provides peculiar romance leading to murder and lust*Lolita
Slow-moving mystery with Modine playing two men - a weakling auto mechanic and an underworld tough-guy - who live in the same city and lead parallel lives*Equinox
Smart, sophisticated comedy about husband and wife lawyers on opposing sides of the same murder case*Adam's rib
Standard Murphy western with every horseopera cliche intact - plot centers around missing shipment of rifles*40 Guns to Apache Pass
Standard film of couple quarreling over adopting war orphan. Nice locations in Switzerland*High Fury
Tender film of faithful horse who returns to mistress after parents sell it to racing stable. Remake of 'Lassie Come Home'*Gypsy Colt
The sword-wielding warrior seeks vengeance on the cult leader who enslaved him and massacred his villagein this fullblooded adventure epic based on Robert E. Howards pulp tales*Conan the Barbarian
Truly perverse black-comedic suspense film about runaway teenage girl staying at her aunt's strange hotel where occupants are extremely weird*Private Parts
Two young men kill prep-school pal, just for the thrill of it, and challenge themselves by inviting friends and family to their apartment afterwards*Rope
Unusual hybrid of comedy, farce and screwball romance that doesn't know what it's trying to be and thus never goes anywhere. Martin plays a New England architect who meets kooky nonconformist Hawn*HouseSitter
Uptight american businessman goes to Naples and finds some unfinished personal business left over from his last visit - when he was an amorous soldier during WW2. Two of the worlds most endearing actors try to keep this souffle from falling, and almost succeed*Macaroni
Wimpy college professor becomes embroiled with pimps, prostitutes and underworld intrigue*Doctor Detroit
Window dresser Locke stumbles across a half-man, half-rat - then tries to parlay him into showbiz by becoming his manager*Ratboy
It is often the case that a man who can't tell a lie thinks that he is the best ... ?*judge of one
Man is the only animal that blushes, or ... ?*needs to
Obstinacy and dogmatism are the surest signs of stupidity. Is there anything more confident, resolute, disdainful, grave and serious than an ... ?*ass
There is no place in a fanatics head, where ... ?*reason can enter
In 1951 the Nobel prize in chemistry was awarded jointly to Edwin Mattison McMillan and Glenn Theodore Seaborg for their discoveries in the chemistry of the transuranium ... ?*elements
In 1951 the Nobel prize in literature was given to Par Fabian Lagerkvist for the artistic vigour and true independence of mind with which he endeavours in his poetry to find answers to the eternal questions ... ?*confronting mankind
In 1951 the Nobel prize in physics was awarded jointly to Sir John Douglas Cockcroft and Ernest Thomas Sinton Walton for their pioneer work on the transmutation of atomic nuclei by artificially accelerated ... ?*atomic particles
In 1951 the Nobel prize in physiology or medicine was given to Max Theiler for his discoveries concerning yellow fever and how to ... ?*combat it
In 1952 the Nobel prize in chemistry was given to awarded jointly to Archer John Porter Martin and Richard  Laurence Millington Synge for their invention of partition ... ?*chromatography
In 1952 the Nobel prize in literature was given to Francois Mauriac for the deep spiritual insight and the artistic intensity with which he has in his novels penetrated the ... ?*drama of human life
In 1952 the Nobel prize in physics was jointly to Felix Bloch and Edward Mills Purcell for their development of new methods for nuclear ... measurements and discoveries in connection therewith*magnetic precision
In 1952 the Nobel prize in physiology or medicine was given to Selman Abraham Waksman for his discovery of streptomycin, the first antibiotic effective against ... ?*tuberculosis
In 1953 the Nobel prize in chemistry was given to Hermann Staudinger for his discoveries in the field of ... ?*macromolecular chemistry
In 1953 the Nobel prize in literature was given to Sir Winston Leonard Spencer Churchill for his mastery of historical and biographical description as well as for brilliant oratory in defending ... ?*exalted human values
In 1953 the Nobel prize in physics was given to Frits (Frederik) Zernike for his demonstration of the phase contrast method, especially for his invention of the phase contrast ... ?*microscope
In 1953 the Nobel prize in physiology or medicine was given to Sir Hans Adolf Krebs for his discovery of the citric acid cycle and the other half to Fritz Albert Lipmann for his discovery of co-enzyme A and its importance for intermediary ... ?*metabolism
In 1954 the Nobel prize in chemistry was given to Linus Carl Pauling for his research into the nature of the chemical bond and its application to the elucidation of the structure of complex ... ?*substances
In 1954 the Nobel prize in literature was given to Ernest Miller Hemingway for his mastery of the art of narrative, most recently demonstrated in 'The Old Man and the Sea' ,and for the influence that he has exerted on ... ?*contemporary style
In 1954 the Nobel prize in physics was divided equally between Max Born for his fundamental research in quantum mechanics, especially for his statistical interpretation of the wavefunction, and Walther Bothe for the ... and his discoveries made therewith*coincidence method
In 1954 the Nobel prize in physiology or medicine was given to John Franklin Enders, Thomas Huckle Weller and Frederick Chapman Robbins for their discovery of the ability of poliomyelitis viruses to grow in cultures of various types of ... ?*tissue
In 1955 the Nobel prize in chemistry was given to Vincent du Vigneaud for his work on biochemically important sulphur compounds, especially for the first synthesis of a polypeptide ... ?*hormone
In 1955 the Nobel prize in literature was given to Halldor Kiljan Laxness, for his vivid epic power which has renewed the great narrative art of ... ?*Iceland
In 1955 the Nobel prize in physics was divided equally between: Willis Eugene Lamb for his discoveries concerning the fine structure of the hydrogen spectrum Polykarp Kusch for his precision determination of the magnetic moment of the ... ?*electron
In 1955 the Nobel prize in physiology or medicine was given to Axel Hugo Theodor Theorell for his discoveries concerning the nature and mode of action of oxidation ... ?*enzymes
In 1956 the Nobel prize in chemistry was given to awarded jointly to Sir Cyril Norman Hinshelwood and Nikolay Nikolaevich Semenov for their researches into the mechanism of ... ?*chemical reactions
In 1956 the Nobel prize in literature was given to Juan Ramon Jimenez for his lyrical poetry, which in Spanish language constitutes an example of high spirit and ... ?*artistical purity
In 1956 the Nobel prize in physics was awarded jointly, one third each, to William Shockley, John Barden and Walter Houser Brattain for their researches on semiconductors and their discovery of the ... ?*transistor effect
In 1956 the Nobel prize in physiology or medicine was given to Andre Frederic Cournand, Werner Forssmann and Dickinson W. Richards for their discoveries concerning heart catherization and pathological changes in the ... ?*circulatory system
In 1957 the Nobel prize in chemistry was given to Lord Alexander R. Todd for his work on nucleotides and nucleotide ... ?*co-enzymes
In 1957 the Nobel prize in literature was given to Albert Camus for his important literary production, which with clear-sighted earnestness illuminates the problems of the ... ?*human conscience in our times
In 1957 the Nobel prize in physics was given jointly Chen Ning Yang and Tsung-Dao Lee for their penetratinginvestigation of the so-called ... which has led to important discoveries regarding the elementary particles*parity laws
In 1957 the Nobel prize in physiology or medicine was given to Daniel Bovet for his discoveries relating to synthetic compounds that inhibit the action of certain body substances, and especially their action on the vascular system and the skeletal ... ?*muscles
In 1958 the Nobel prize in chemistry was given to Frederick Sanger for his work on the structure of proteins, especially that of ... ?*insulin
In 1958 the Nobel prize in physics was awarded jointly to Pavel Alekseyevich Cherenkov, Il'ja Mikhailovich Frank and Igor Yevgenyevich Tamm for the discovery and the interpretation of the ... ?*Cherenkov effect
In 1958 the Nobel prize in physiology or medicine was given to George Wells Beadle and Edward Lawrie Tatum for their discovery that genes act by regulating definite chemical events and the other half to Joshua Lederberg for his discoveries concerning genetic recombination and the organization of the genetic material of ... ?*bacteria
In 1959 the Nobel prize in chemistry was given to Jaroslav Heyrovsky for his discovery and development of the polarographic methods of ... ?*analysis
In 1959 the Nobel prize in literature was given to Salvatore Quasimodo for his lyrical poetry, which with classical fire expresses the tragic experience of ... ?*life in our own times
In 1959 the Nobel prize in physics was awarded jointly to Emilio Gino Segre and Owen Chamberlain for their discovery of the ... ?*antiproton
In 1959 the Nobel prize in physiology or medicine was given to Severo Ochoa and Arthur Kornberg for their discovery of the mechanisms in the biological synthesis of ribonucleic acid and deoxiribonucleic ... ?*acid
In 1960 the Nobel prize in chemistry was given to Willard Frank Libby for his method to use carbon-14 for age determination in archaeology, geology, geophysics, and other branches of ... ?*science
In 1960 the Nobel prize in literature was given to Saint-John Perse  (pen-name of Alexis Leger), for the soaring flight and the evocative imagery of his poetry which in a visionary fashion reflects the ... ?*conditions of our time
In 1960 the Nobel prize in physics was given to Donald A. Glaser for the invention of the bubble ... ?*chamber
In 1960 the Nobel prize in physiology or medicine was given to Sir Frank Macfarlane Burnet and Sir Peter Brian Medawar for discovery of acquired immunological ... ?*tolerance
In 1961 the Nobel prize in chemistry was given to Melvin Calvin for his research on the carbon dioxide assimilation in ... ?*plants
In 1961 the Nobel prize in literature was given to Ivo Andri'c for the epic force with which he has traced themes and depicted human destinies drawn from the history of ... ?*his country
In 1961 the Nobel prize in physics was divided equally between Robert Hofstadter for his pioneering studies of electron scattering in atomic nuclei and for his thereby achieved discoveries concerning the stucture of the nucleons Rudolf Ludwig Mossbauer for his researches concerning the resonance absorption of gamma radiation and his discovery in this connection of the effect which ... ?*bears his name
In 1961 the Nobel prize in physiology or medicine was given to Georg von Bekesy for his discoveries of the physical mechanism of stimulation within the ... ?*cochlea
In 1962 the Nobel prize in chemistry was divided equally between Max Ferdinand Perutz and Sir John Cowdery Kendrew for their studies of the structures of globular ... ?*proteins
In 1962 the Nobel prize in literature was given to John Steinbeck for his realistic and imaginative writings, combining as they do sympathetic humour and keen ... ?*social perception
In 1962 the Nobel prize in physics was given to Lev Davidovich Landau for his pioneering theories for condensed matter, especially liquid ... ?*helium
In 1962 the Nobel prize in physiology or medicine was given to Francis Harry Compton Crick, James Dewey Watson and Maurice Hugh Frederick Wilkins for their discoveries concerning the molecular structure of nuclear acids and its significance for information transfer in ... ?*living material
In 1963 the Nobel prize in chemistry was divided equally between Karl Ziegler and Giulio Natta for their discoveries in the field of the chemistry and technology of high ... ?*polymers
In 1963 the Nobel prize in literature was given to  Giorgos Seferis (alias Seferiadis) for his eminent lyrical writing, inspired by a deep feeling for the Hellenic ... ?*world of culture
In 1963 the Nobel prize in physics was divided, one half being awarded to  Eugene P. Wigner for his contributions to the theory of the atomic nucleus and the elementary particles, particularly through the discovery and application of fundamental symmetry principles and the other half jointly to Maria Goeppert-Mayer and J.Hans D. Jensen for their discoveries concerning nuclear ... ?*shell structure
In 1963 the Nobel prize in physiology or medicine was given to Sir John Carew Eccles, Sir Alan Lloyd Hodgkin and Sir Andrew Fielding Huxley for their discoveries concerning the ionic mechanisms involved in excitation and inhibition in the peripheral and central portions of the nerve cell ... ?*membrane
In 1964 the Nobel prize in chemistry was given to Dorothy Crowfoot Hodgkin for her determinations by X-ray techniques of the structures of important biochemical ... ?*substances
In 1964 the Nobel prize in literature was given to Jean-Paul Satre (who declined the prize) for his work which, rich in ideas and filled with the spirit of freedom and the quest for truth, has exerted a farreaching influence ... ?*on our age
In 1964 the Nobel prize in physics was divided, one half being awarded to Charles H. Townes, the other half jointly to Nicolay Gennadiyevich Basov and Aleksandr Mikhailovich Prokhorov for fundamental work in the field of quantum electronics, which has led to the construction of oscillators and amplifiers based on the maser-laser ... ?*principle
In 1964 the Nobel prize in physiology or medicine was given to Konrad Bloch and Feodor Lynen for their discoveries concerning the mechanism and regulation of the cholesterol and ... ?*fatty acid metabolism
In 1965 the Nobel prize in chemistry was given to Robert Burns Woodward for his outstanding achievements in the art of ... ?*organic synthesis
In 1965 the Nobel prize in literature was given to Michail Aleksandrovich Sholokhov for the artistic power and integrity with which, in his epic of the Don, he has given expression to a historic phase in the life of ... ?*the Russian people
In 1965 the Nobel prize in physics was awarded jointly to Sin-itiro Tomonaga, Julian Schwinger and Richard P. Feynman for their fundamental work in quantum electrodynamics, with deep-ploughing consequences for the physics of elementary ... ?*particles
In 1965 the Nobel prize in physiology or medicine was given to Francois Jacob, Andre Lwoff and Jacoues Monod for their discoveries concerning genetic control of enzyme and ... ?*virus synthesis
In 1966 the Nobel prize in chemistry was given to Robert S. Mulliken for his fundamental work concerning chemical bonds and the electronic structure of molecules by the molecular orbital ... ?*method
In 1966 the Nobel prize in literature was divided equally between Shmuel Yosef Agnon for his profoundly characteristic narrative art with motifs from the life of the Jewish people and Nelly Sachs for her outstanding lyrical and dramatic writing, which interprets Israel's destiny with ... ?*touching strength
In 1966 the Nobel prize in physics was given to Alfred Kastler for the discovery and development of optical methods for studying hertzian resonances in ... ?*atoms
In 1966 the Nobel prize in physiology or medicine was given to Peyton Rous for his discovery of tumorinducing viruses and the other half to  Charles Brenton Huggins for his discoveries concerning hormonal treatment of*prostatic cancer
In 1967 the Nobel prize in chemistry was divided, one half being awarded to  Manfred Eigen and the other half jointly to Ronald George Wreyford Norrish and Lord George Porter for their studies of extremely fast chemical reactions, effected by disturbing the equlibrium by means of very short ... ?*pulses of energy
In 1967 the Nobel prize in literature was given to Miguel Angel for his vivid literary achievement, deep-rooted in the national traits and traditions of ... ?*Indian peoples of Latin America
In 1967 the Nobel prize in physics was given to Hans Albrecht Bethe for his contributions to the theory of nuclear reactions, especially his discoveries concerning the energy production in ... ?*stars
In 1967 the Nobel prize in physiology or medicine was given to Ragnar Grant, Haldan Keffer Hartline and George Wald for their discoveries concerning the primary physiological and chemical visual processes in the ... ?*eye
In 1968 the Nobel prize in chemistry was given to Lars Onsager for the discovery of the reciprocal relations bearing his name, which are fundamental for the thermodynamics of irreversible ... ?*processes
In 1968 the Nobel prize in literature was given to Yasunari Kawabata for his narrative mastery, which with great sensibility expresses the essence of the ... ?*Japanese mind
In 1968 the Nobel prize in physics was given to Luis W. Alvarez for his decisive contributions to elementary particle physics, in particular the discovery of a large number of resonance states, made possible through his development of the technique of using hydrogen bubble chamber and ... ?*data analysis
In 1968 the Nobel prize in physiology or medicine was given to Robert W. Holley, Har Gobind Khorana and Marshall W. Nirenberg for their interpretation of the genetic code and its function in protein ... ?*synthesis
In 1969 the Nobel prize in chemistry was divided equally between Sir Derek H.R. Baron and Odd Hassel for their contributions to the development of the concept of conformation and its application in ... ?*chemistry
In 1969 the Nobel prize in literature was given to Samuel Beckett for his writing, which - in new forms for the novel and drama - in the destitution of modern man acquires its ... ?*elevation
In 1969 the Nobel prize in physics was given to Murray Gell-Mann for his contributions and discoveries concerning the classification of elementary particles and their ... ?*interactions
In 1969 the Nobel prize in physiology or medicine was given to Max Delbruck, Alfred D. Hershey and Salvador E. Luria for their discoveries concerning the replication mechanism and the gentic structure of ... ?*viruses
In 1970 the Nobel prize in chemistry was given to Luis F. Leloir for his discovery of sugar nucleotides and their role in the biosynthesis of ... ?*carbohydrates
In 1970 the Nobel prize in literature was given to Aleksandr Isaevich Solzhenitsyn for the ethical force with which he has pursued the indispensable traditions of ... ?*Russian literature
In 1970 the Nobel prize in physics was divided equally between Hannes Alfven for fundamental work and discoveries in magneto-hydrodynamics with fruitful applications in different parts of plasma physics and Louis Neel for fundamental work and discoveries concerning antiferromagnetism and ferrimagnetism which have led to important applications in ... ?*solid state physics
In 1970 the Nobel prize in physiology or medicine was given to Sir Bernard Katz, Ulf von Euler and Julius Axelrod for their discoveries concerning the humoral transmittors in the nerve terminals and the mechanism for their storage, release and ... ?*inactivation
In 1971 the Nobel prize in chemistry was given to Gerhard Herzberg for his contributions to the knowledge of electronic stucture and geometry of molecules, particularly ... ?*free radicals
In 1971 the Nobel prize in literature was given to Pablo Neruda for a poetry that with the action of an elemental force brings alive a continent's destiny and ... ?*dreams
In 1971 the Nobel prize in physics was given to Dennis Gabor for his invention and development of the holographic ... ?*method
In 1971 the Nobel prize in physiology or medicine was given to Earl W.Jr. Sutherland for his discoveries concerning the mechanisms of the action of ... ?*hormones
In 1972 the Nobel prize in chemistry was divided, one half being awarded to Christian B. Anfinsen for his work on ribonuclease, especially concerning the connection between the amino acid sequence and the biologically active confirmation and the other half jointly to Stanford Moore and William H. Stein for their contribution to the understanding of the connection between chemical structure and catalytic activity of the active centre of the ribonuclease ... ?*molecule
In 1972 the Nobel prize in literature was given to Heinrich Boll for his writing which through its combination of a broad perspective on his time and a sensitive skill in characterization has contributed to a renewal of ... ?*German literature
In 1972 the Nobel prize in physics was awarded jointly to: John Bardeen, Leon N. Cooper and J. Robert Schrieffer for their jointly developed theory of superconductivity, usually called the ... ?*BCS-theory
In 1972 the Nobel prize in physiology or medicine was given to Gerald M. Edelman and Rodney R. Porter for their discoveries concerning the chemical structure of ... ?*antibodies
In 1973 the Nobel prize in chemistry was divided equally between Ernst Ott Fischer and Sir Geoffrey Wilkinson for their pioneering work, performed independently, on the chemistry of the organometallic, so called sandwich ... ?*compounds
In 1973 the Nobel prize in literature was given to Patrick White for an epic and psychological narrative art which has introduced a new continent ... ?*into literature
In 1973 the Nobel prize in physics was divided, one half being equally shared between Leo Esaki and Ivar Glaever for their experimental discoveries regarding tunneling phenomena in semiconductors and superconductors, respectively, and the other half to Brian D. Josephson for his theoretical predictions of the properties of a supercurrent through a tunnel barrier, in particular those phenomena which are generally known as the ... ?*Jospehson effects
In 1973 the Nobel prize in physiology or medicine was given to Karl von Frisch, Konrad Lorenz and Nikolaas Tinbergen for their discoveries concerning organization and elicitation of individual and social ... ?*behaviour patterns
In 1974 the Nobel prize in chemistry was given to Paul J. Flory for his fundamental achievements, both theoretical and experimental, in the physical chemistry of the ... ?*macromolecules
In 1974 the Nobel prize in literature was divided equally between Eyvind Johnson for a narrative art, farseeing in lands and ages, in the service of freedom and Harry Martinson for writings that catch the dewdrop and ... ?*reflect the cosmos
In 1974 the Nobel prize in physics was awarded jointly to Sir Martin Ryle and Antony Hewish for their pioneering research in radio astrophysics Ryle for his observations and inventions, in particular of the aperture synthesis technique, and Hewish for his decisive role in the discovery of ... ?*pulsars
In 1974 the Nobel prize in physiology or medicine was given to Albert Claude, Christian De Duve and George E. Palada for their discoveries concerning the structural and functional organization of the ... ?*cell
In 1975 the Nobel prize in chemistry was divided equally between Sir John Warcup Cornforth for his work on the stereochemistry of enzyme-catalyzed reactions and Vladimir Prelog for his research into the stereochemistry of organic molecules and ... ?*reactions
In 1975 the Nobel prize in literature was given to Eugenio Montale for his distinctive poetry which, with great artistic sensitivity, has interpreted human values under the sign of an outlook on life with ... ?*no illusions
In 1975 the Nobel prize in physics was awarded jointly to Aage Bohr, Ben Mottelson and James Rainwater for the discovery of the connection between collective motion and particle motion in atomic nuclei and the development of the theory of the structure of the atomic nucleus based on ... ?*this connection
In 1975 the Nobel prize in physiology or medicine was given to David Baltimore, Renato Dulbecco and Howard Martin Temin for their discoveries concerning the interaction between tumour viruses and the genetic material of ... ?*the cell
In 1976 the Nobel prize in chemistry was given to William N. Lipscomb for his studies on the structure of boranes illuminating problems of chemical ... ?*bonding
In 1976 the Nobel prize in literature was given to Saul Bellow for the human understanding and subtle analysis of contemporary culture that are ... ?*combined in his work
In 1976 the Nobel prize in physics was divided equally between Burton Richter and Samule C.C. Ting for their pioneering work in the discovery of a heavy elementary particle of a new ... ?*kind
In 1976 the Nobel prize in physiology or medicine was given to Baruch S. Blumberg and D. Carleton Gajduser for their discoveries concerning new mechanisms for the origin and dissemination of ... ?*infectious diseases
In 1977 the Nobel prize in chemistry was given to Ilya Prigogine for his contributions to non-equilibrium thermodynamics, particularly the theory of dissipative ... ?*structures
In 1977 the Nobel prize in literature was given to Vicente Aleixandre for a creative poetic writing which illuminates man's condition in the cosmos and in present-day society, at the same time representing the great renewal of the traditions of ... ?*Spanish poetry
In 1977 the Nobel prize in physics was divided equally between: Philip W. Anderson, Sir Nevill F. Mott and John H. Van Vleck for their fundamental theoretical investigations of the electronic structure of magnetic and disordered ... ?*systems
In 1977 the Nobel prize in physiology or medicine was divided equally, one half awarded jointly to Roger Guillemin and Andrew V. Schally for their discoveries concerning the peptide hormone production of the brain and the other half awarded to Rosalyn Yalow for the development of radioimmunoassays of ... ?*peptide hormones
In 1978 the Nobel prize in chemistry was given to Peter D. Mitchell for his contribution to the understanding of biological energy transfer through the formulation of the chemiosmotic ... ?*theory
In 1978 the Nobel prize in literature was given to for his impassioned narrative art which, with roots in a Polish-Jewish cultural tradition, brings ... ?*universal human conditions to life
In 1978 the Nobel prize in physics was divided, one half being awarded to Pyotr Leonidovich Kapitsa for his basic inventions and discoveries in the area of low-temperature physics and the other half divided equally between Arno A. Penzias and Robert W. Wilson for their discovery of cosmic microwave background ... ?*radiation
In 1978 the Nobel prize in physiology or medicine was given to Werner Arber, Daniel Nathans and Hamilton O. Smith for the discovery of restriction enzymes and their application to problems of ... ?*molecular genetics
In 1979 the Nobel prize in chemistry was divided equally between Herbert C. Brown and Georg Wittig for their development of the use of boron- and phosphorus-containing compounds, respectively, into important reagents in organic ... ?*synthesis
In 1979 the Nobel prize in literature was given to Odysseus Elitis (alias Alepoudhelis) for his poetry, which, against the background of Greek tradition, depicts with sensuous strength and intellectual clear-sightedness modern man's struggle for freedom and ... ?*creativeness
In 1979 the Nobel prize in physics was divided equally between Sheldon L. Glashow, Abdus Salam and Steven Weinberg for their contributions to the theory of the unified weak and electromagnetic interaction between elementary particles, including inter alia the prediction of the weak neutral ... ?*current
In 1979 the Nobel prize in physiology or medicine was given to Alan M. Cormack and Sir Godfrey N. Hounsfield for the development of computer assisted ... ?*tomography
In 1980 the Nobel prize in chemistry was divided, one half being awarded to Paul Berg for his fundamental studies of the biochemistry of nucleic acids, with particular regard to recombinant-DNA and the other half jointly to Walter Gilbert and Frederick Sanger for their contributions concerning the determination of base sequences in ... ?*nucleic acids
In 1980 the Nobel prize in literature was given to Czeslaw Milosz, who with uncompromising clear-sightedness voices man's exposed condition in a world of ... ?*severe conflicts
In 1980 the Nobel prize in physics was divided equally between James W. Cronin and Val L. Fitch for the discovery of violations of fundamental symmetry principles in the decay of neutral ... ?*K-mesons
In 1980 the Nobel prize in physiology or medicine was given to Baruj Benacerraf, Jean Dausset and George D. Snell for their discoveries concerning genetically determined structures on the cell surface that regulate*immunological reactions
In 1981 the Nobel prize in chemistry was given to Kenichi Fukui and Roald Hoffmann for their theories, developed independently, concerning the course of chemical ... ?*reactions
In 1981 the Nobel prize in literature was given to Elias Canetti for writings marked by a broad outlook, a wealth of ideas and ... ?*artistic power
In 1981 the Nobel prize in physics was awarded by one half jointly to Nicolaas Bloembergen and Arthur L. Schawlow for their contribution to the development of laser spectroscopy and the other half to Kai M. Siegbahn for his contribution to the development of high- resolution electron ... ?*spectroscopy
In 1981 the Nobel prize in physiology or medicine was divided between Roger W. Sperry for his discoveries concerning the functional specialization of the cerebral hemispheres. and the other half awarded jointly to David H. Hubel and Torstein N. Wiesel for their discoveries concerning information processing in the ... ?*visual system
In 1982 the Nobel prize in chemistry was given to Sir Aaron Klug for his development of crystallographic electron microscopy and his structural elucidation of biologically important nuclei acid-protein ... ?*complexes
In 1982 the Nobel prize in literature was given to Gabriel Garcia Marquez for his novels and short stories, in which the fantastic and the realistic are combined in a richly composed world of imagination, reflecting a continent's ... ?*life and conflicts
In 1982 the Nobel prize in physics was given to Kenneth G. Wilson for his theory for critical phenomena in connection with phase ... ?*transitions
In 1982 the Nobel prize in physiology or medicine was given to Sune K. Bergstrom, Bengt I. Samuelsson and Sir John R. Vane for their discoveries concerning prostaglandins and related biologically ... ?*active substances
In 1983 the Nobel prize in chemistry was given to Henry Taube for his work on the mechanisms of electron transfer reactions, especially in ... ?*metal complexes
In 1983 the Nobel prize in literature was given to Sir William Golding for his novels which, with the perspicuity of realistic narrative art and the diversity and universality of myth, illuminate the human condition in the ... ?*world of today
In 1983 the Nobel prize in physics was divided equally between  Subramanyan Chandrasekhar for his theoretical studies of the physical processes of importance to the structure and evolution of the stars and William A. Fowler for his theoretical and experimental studies of the nuclear reactions of importance in the formation of the chemical elements in the ... ?*universe
In 1983 the Nobel prize in physiology or medicine was given to Barbara Clintock for her discovery of mobile ... ?*genetic elements
In 1984 the Nobel prize in chemistry was given to Robert Bruce Merrifield for his development of methodology for chemical synthesis on a solid ... ?*matrix
In 1984 the Nobel prize in literature was given to Jaroslav Seifert, for his poetry which endowed with freshness, sensuality and rich inventiveness provides a liberating image of the indomitable spirit and versatility of man ... ?*versatility of man
In 1984 the Nobel prize in physics was awarded jointly to Carlo Rubbia and Simon van der Meer for their decisive contributions to the large project, which led to the discovery of the field particles W and Z, communicators of weak ... ?*interaction
In 1984 the Nobel prize in physiology or medicine was given to Niels K. Jerne, Georges J.F. Kohler and Cesar Milstein for theories concerning the specificity in development and control of the immune system and the discovery of the principle for production of monoclonal ... ?*antibodies
In 1985 the Nobel prize in chemistry was awarded jointly to Herbert A. Hauptman and Jerome Karle for their outstanding achievements in the development of direct methods for the determination of  ... ?*crystal structures
In 1985 the Nobel prize in literature was given to Claude Simon, who in his novel combines the poet's and the painter's creativeness with a deepened awareness of time in the depiction of the ... ?*human condition
In 1985 the Nobel prize in physics was given to Klaus von Klitzing for the discovery of the quantized ... ?*Hall effect
In 1985 the Nobel prize in physiology or medicine was given to Michael S. Brown and Joseph L. Goldstein for their discoveries concerning the regulation of cholesterol ... ?*metabolism
In 1986 the Nobel prize in chemistry was awarded jointly to Dudley R. Herschbach, Yuan T. Lee and John C. Polanyi for their contributions concerning the dynamics of chemical ... ?*elementary processes
In 1986 the Nobel prize in literature was given to Wole Soyinka,  who in a wide cultural perspective and with poetic overtones fashions the ... ?*drama of existence
In 1986 the Nobel prize in physics was awarded by one half to Ernst Ruska for his fundamental work in electron optics, and for the design of the first electron microscope. Gerd Binnig and Heinrich Rohrer for their design of the scanning tunneling ... ?*microscope
In 1986 the Nobel prize in physiology or medicine was given to Stanley Cohen and Rita Levi-Montalcini for their discoveries of ... ?*growth factors
In 1987 the Nobel prize in chemistry was given to Donald J. Cram, Jean-Marie Lehn, and Charles J. Pedersen for their development and use of molecules with structure-specific interactions of high ... ?*selectivity
In 1987 the Nobel prize in literature was given to Joseph Brodsky, for an all-embracing authorship, imbued with clarity of thought and ... ?*poetic intensity
In 1987 the Nobel prize in physics was awarded jointly to J. Georg Bednorz and K. Alexander Muller for their important breakthrough in the discovery of superconductivity in ... ?*ceramic materials
In 1987 the Nobel prize in physiology or medicine was given to Susumu Tonegawa for his discovery of the genetic principle for generation of ... ?*antibody diversity
In 1988 the Nobel prize in chemistry was awarded jointly to Johann Deisenhofer, Robert Huber and Hartmut Michel for the determination of the three-dimensional structure of a photosynthetic reaction ... ?*centre
In 1988 the Nobel prize in literature was given to Naguib Mahfouz,  who through works rich in nuance-now clearsightedly realistic, now evocatively ambigous-has formed an Arabian narrative art that applies to ... ?*all mankind
In 1988 the Nobel prize in physics was awarded jointly to Leon M Lederman, Melvin Schwarts and Jack Steinberger for the neutrino beam method and the demonstration of the doublet structure of the leptons through the discovery of the ... ?*muon neutrino
In 1988 the Nobel prize in physiology or medicine was given to Sir James W. Black, Gertrude B. Elion and George H. Hitchings for their discoveries of important principles for ... ?*drug treatment
In 1989 the Nobel prize in chemistry was given to Sidney Altman and Thomas R. Cech for their discovery of catalytic properties of ... ?*RNA
In 1989 the Nobel prize in literature was given to Camilio Jose Cela for a rich and intensive prose, which with restrained compassion forms a challenging vision of man's ... ?*vulnerability
In 1989 the Nobel prize in physics was divided between Norman F. Ramsey for the invention of the separated oscillatory fields method and its use in the hydrogen maser and other atomic clocks. and the other half jointly to Hans G. Dehmelt and Wolfgang Paul for the development of the ion ... ?*trap technique
In 1989 the Nobel prize in physiology or medicine was given to J. Michael Bishop and Harold E. Varmus for their discovery of the cellular origin of retroviral ... ?*oncogenes
In 1990 the Nobel prize in chemistry was given to Elias James Corey for his development of the theory and methodology of ... ?*organic synthesis
In 1990 the Nobel prize in literature was given to Octavio Paz, for impassioned writing with wide horizons, characterized by sensuous intelligence and humanistic ... ?*integrity
In 1990 the Nobel prize in physics was awarded jointly to: Jerome I. Friedman, Henry W. Kendall and Richard E. Taylor for their pioneering investigations concerning deep inelastic scattering of electrons on protons and bound neutrons, which have been of essential importance for the development of the quark model in ... ?*particle physics
In 1990 the Nobel prize in physiology or medicine was given to Joseph Murray and E. Donnall Thomas for their discoveries concerning organ and cell transplantation in the treatment of ... ?*human disease
In 1991 the Nobel prize in chemistry was given to Richard R. Ernst for his contributions to the development of the methodology of high resolution nuclear magnetic resonance (NMR) ... ?*spectroscopy
In 1991 the Nobel prize in literature was given to Nadine Gordimer, who through her magnificent epic writing has - in the words of Alfred Nobel - been of very great benefit to ... ?*humanity
In 1991 the Nobel prize in physics was given to Pierre-Gilles de Gennes for discovering that methods developed for studying order phenomena in simple systems can be generalized to more complex forms of matter, in particular to liquid crystals and ... ?*polymers
In 1991 the Nobel prize in physiology or medicine was given to Erwin Neher and Bert Sakmann for their discoveries concerning the function of single ion channels ... ?*in cells
In 1992 the Nobel prize in chemistry was given to Rudolph A. Marcus for his contributions to the theory of electron transfer reactions in ... ?*chemical systems
In 1992 the Nobel prize in literature was given to Derek Walcott, for a poetic oeuvre of great luminosity, sustained by a historical vision, the outcome of a ... ?*multicultural commitment
In 1992 the Nobel prize in physics was given to Georges Charpak for his invention and development of particle detectors, in particular the multiwire proportional ... ?*chamber
In 1992 the Nobel prize in physiology or medicine was given to Edmond Fischer and Edwin G. Krebs for their discoveries concerning reversible protein phosphorylation as a biological ... ?*regulatory mechanism
In 1993 the Nobel prize in chemistry was awarded for contributions to the developments of methods within DNA-based chemistry equally between Kary B. Mullis for his invention of the polymerase chain reaction (PCR) method. and Michael Smith for his fundamental contributions to the establishment of oligonucleiotide-based, site-directed mutagenesis and its development for ... ?*protein studies
In 1993 the Nobel prize in literature was given to Toni Morrison, who in novels characterized by visionary force and poetic import, gives life to an essential aspect of*American reality
In 1993 the Nobel prize in physics was awarded jointly to Russel A. Hulse and Joseph H. Taylor Jr. for the discovery of a new type of pulsar, a discovery that has opened up new possibilities for the study of ... ?*gravitation
In 1993 the Nobel prize in physiology or medicine was given to Richard J. Roberts and Phillip A. Sharp for their independent discoveries of ... ?*split genes
In 1994 the Nobel prize in chemistry was given to George A. Olah for his contribution to ... ?*carbocation chemistry
In 1994 the Nobel prize in literature was given to Kenzaburo Oe who with poetic force creates an imagined world, where life and myth condense to form a disconcerting picture of the human ... ?*predicament today
In 1994 the Nobel prize in physics was awarded for pioneering contributions to the development of neutron scattering techniques for studies of condensed matter to Bertram N. Brockhouse for the development of neutron spectroscopy and Clifford G. Shull for the development of the neutron diffraction ... ?*technique
In 1994 the Nobel prize in physiology or medicine was given to Alfred G. Gilman and Martin Rodbell for their discovery of G-proteins and the role of these proteins in ... ?*signal transduction in cells
In 1995 the Nobel prize in chemistry was awarded jointly to Paul Crutzen, Mario Molina and F. Sherwood Rowland for their work in atmospheric chemistry, particularly concerning the formation and decomposition of ... ?*ozone
In 1995 the Nobel prize in literature was given to Seamus Heaney for works of lyrical beauty and ethical depth, which exalt everyday miracles and the... ?*living past
In 1995 the Nobel prize in physics was awarded for pioneering experimental contributions to lepton physics, with one half to Martin L. Perl for the discovery of the tau lepton. and the other half to Frederick Reines for the detection of the ... ?*neutrino
In 1995 the Nobel prize in physiology or medicine was given to Edward B. Lewis, Christiane Nusslein-Volhard and Eric F. Wieschaus for their discoveries concerning the genetic control of early ... ?*embryonic development
In 1996 the Nobel prize in chemistry was given to Robert F. Curl Jr., Sir Harold W. Kroto, and Richard E. Smalley for their discovery of ... ?*fullerenes
In 1996 the Nobel prize in literature was given to  Wislawa Szymborska, for poetry that with ironic precision allows the historical and biological context to come to light in fragments of ... ?*human reality
In 1996 the Nobel prize in physics was awarded jointly to David M. Lee, Douglas D. Osheroff and Robert C. Richardson for their discovery of superfluidity in ... ?*helium-3
In 1996 the Nobel prize in physiology or medicine was given to Peter C. Doherty and Rolf M. Zinkernagel for their discoveries concerning the specificity of the cell mediated ... ?*immune defence
In 1997 the Nobel prize in chemistry was divided, one half being awarded jointly to  Paul D. Boyer and John E. Walker for their elucidation of the enzymatic mechanism underlying the synthesis of adenosine triphosphate (ATP) and with one half to Jens C. Skou for the first discovery of an ion-transporting ... ?*enzyme
In 1997 the Nobel prize in literature was given to Dario Fo, who emulates the jesters of the Middle Ages in scourging authority and upholding the dignity of the ... ?*downtrodden
In 1997 the Nobel prize in physics was awarded jointly to Steven Chu, Claude Cohen-Tannoudji and William D. Philips for development of methods to cool and trap atoms with ... ?*laser light
In 1997 the Nobel prize in physiology or medicine was given to Stanley B. Prusiner for his discovery of Prions - a new biological principle of ... ?*infections
In 1998 the Nobel prize in chemistry was awarded for pioneering contributions in developing methods that can be used for theoretical studies of the properties of molecules and the chemical processes in which they are involved. The prize was divided equally between Walter Kohn for his development of the density-functional theory and John A. Pople for his development of computational methods in ... ?*quantum chemistry
In 1998 the Nobel prize in literature was given to Jose Saramago, who with parables sustained by imagination, compassion and irony continually enables us once again to apprehend an ... ?*elusory reality
In 1998 the Nobel prize in physics was awarded jointly to Robert B. Laughlin, Horst L. Stormer and Daniel C. Tsui for their discovery of a new form of quantum fluid with fractionally charged ... ?*excitations
In 1998 the Nobel prize in physiology or medicine was given to Robert Furchgott, Louis J. Ignarro and Ferid Murad for their discoveries concerning nitric oxide as a signalling molecule in the ... ?*cardiovascular system
In 1999 the Nobel prize in chemistry was given to Ahmed Zewail for his studies of the transition states of chemical reactions using femtosecond ... ?*spectroscopy
In 1999 the Nobel prize in literature was given to Gunter Grass whose frolicsome black fables portray the forgotten face ... ?*of history
In 1999 the Nobel prize in physics was awarded jointly to Gerardus Thooft and Martinus J.G. Veltman for elucidating the quantum structure of ... ?*electroweak interactions in physics
In 1999 the Nobel prize in physiology or medicine was given to Gunter Blobel for the discovery that proteins have intrinsic signals that govern their transport and localization in the ... ?*cell
In 2000 the Nobel prize in chemistry was given to Alan J. Heeger, Alan G. Macdiarmid and Hideki Shirikawa for the discovery and development of conductive ... ?*polymers
In 2000 the Nobel prize in literature was given to Gao Xingjian for an oeuvre of universal validity, bitter insights and linguistic ingenuity, which has opened new paths for the ... ?*Chinese novel and drama
In 2000 the Nobel prize in physics was given to Zhores I. Alferov, and Herbert Kroemer for developing semiconductor heterostructures used in high-speed- and opto-electronics and and one half to Jack St. Clair Kilby for his part in the invention of the ... ?*integrated circuit
In 2000 the Nobel prize in physiology or medicine was given to Arvid Carlsson, Paul Greengard and Eric Kandel for their discoveries concerning signal transduction in the ... ?*nervous system
Experience is the name everyone gives to ... ?*their mistakes
There's only one thing  in the world worse than being talked about, and that is ... ?*not being talked about
... 's distinctive and controversial brand of pragmatism expresses itself along two main axes. One is negative---a critical diagnosis of what he takes to be defining projects of modern philosophy. The other is positive---an attempt to show what intellectual culture might look like, once we free ourselves from the governing metaphors of mind and knowledge in which the traditional problems of epistemology and metaphysics (and indeed, in his view, the self-conception of modern philosophy) are rooted*Richard Rorty
... asserts the relativity of morality*moral relativism
... asserts the relativity of truth*Cognitive relativism
... can be defined as an abstract object or term which ranges over particular things. The classic problem involves whether abstract objects such as "largeness" exist in a realm independent of human thought. Realists argue that they do*A universal
... can be thought of as the metaphysical theory that attempts to account for the truth of claims like 'It is possible that there are Aliens' without appealing to any nonactual objects whatsoever. What makes this theory so philosophically interesting, is that there is no obviously correct way to account for the truth of claims like 'It is possible that there are Aliens' without appealing to possible but nonactual objects*Actualism
... denied the soundness of metaphysics and traditional philosophy; they asserted that many philosophical problems are indeed meaningless. During 1930s the most important representatives emigrated to USA, where they influenced American philosophy. Until 1950s it was the leading philosophy of science; today its influence persists especially in the way of doing philosophy, in the great attention given to the analysis of scientific thought and in the definitely acquired results of the technical researches on formal logic and the theory of probability*Logical positivism
... denotes any metaphysical theory which claims that reality consists of a multiplicity of distinct, fundamental entities. The term was first used by Christian Wolff (1679-1754), and later popularized by William James in The Will to Believe. It is distinguished from both monism, the view that one kind of thing exists, and dualism, the view that two kinds of things exist*Pluralism
... entails that the individual self is either the motivating moral force and is, or should, be the end of moral action*Egoism
... have sought to deflate the universal pretensions of liberal theory. The main target has been Rawls description of the original position as an 'Archemedian point' from which the structure of a social system can be appraised, a position whose special virtue is that it allows us to regard the human condition 'from the perspective of eternity', from all social and temporal points of view*Communitarians
... having studied science at the University of Vienna, moved into philosophy for his doctoral thesis, made a name for himself both as an expositor and (later) as a critic of Karl Popper's 'critical rationalism', and went on to become one of this century's most famous philosophers of science. An imaginative maverick, he became a critic of philosophy of science itself, particularly of 'rationalist' attempts to lay down or discover rules of scientific method*Paul Feyerabend
... is a formulation utilitarianism which maintains that a behavioral code or rule is morally right if the consequences of adopting that rule are more favorable than unfavorable to everyone. It is contrasted with act utilitarianism which maintains that the morality of each action is to be determined in relation to the favorable or unfavorable consequences that emerge from that action*Rule utilitarianism
... is a movement in cognitive science which hopes to explain human intellectual abilities using artificial neural networks (also known as 'neural networks' or 'neural nets'). Neural networks are simplified models of the brain composed of large numbers of units (the analogs of neurons) together with weights that measure the strength of connections between the units. These weights model the effects of the synapses that link one neuron to another. Experiments on models of this kind have demonstrated an ability to learn such skills as face recognition, reading, and the detection of simple grammatical structure*Connectionism
... is a name given to a group of ancient philosophers who, from the existing philosophical beliefs, tried to select the doctrines that seemed to them most reasonable, and out of these constructed a new system*Eclecticism
... is a philosophical position maintaining that our minds gain knowledge independently of experience through innate ideas or mental faculties*A priorism
... is a term used in both ethics and epistemology. In ethics it deals with determining right actions and appropriate beliefs. In epistemology, it is the central component to knowledge as justified true belief*Justification
... is a term used to identify a type of knowledge which is obtained independently of experience*A priori
... is a theory in the philosophy of mind which maintains that talk of mental events should be translated into talk about observable behavior*Behaviorism
... is an epistemological position that we do not have knowledge or justification for believing in objective moral principles. It does not involve the rejection of moral values themselves, but simply the denial that we have knowledge of an objective realm of morals*Moral skepticism
... is defined as the complex method of obtaining information about our surrounding world, specifically through our senses, and apprehending this information as beliefs*Perception
... is sometimes identified (usually by its critics) as the thesis that all points of view are equally valid. In ethics, this amounts to saying that all moralities are equally good; in epistemology it implies that all beliefs, or belief systems, are equally true*Relativism
... is the belief that all values are baseless and that nothing can be known or communicated. It is often associated with extreme pessimism and a radical skepticism that condemns existence*Nihilism
... is the description and study of appearances. The term has come to be closely associated with the method of inquiry that was originated by Brentano and further developed by Husserl. The movement originally placed an emphasis on human experience descriptions, as the human experience was directed onto objects*Phenomenology
... is the idea, often associated with the political theories of John Locke and the "founders" of the American republic, that government can and should be legally limited in its powers, and that its authority depends on its observing these limitations*Constitutionalism
... is the position that arises out of the difficulties present in the dualism between the phenomenon and the object. It maintains that all we know are phenomena; we know nothing of the external things causing the phenomena*phenomenalism
... is the theory that God or the ultimate nature of reality is to be conceived as some form of will (or conation). This theory is contrasted to intellectualism, which gives primacy to God's reason*Voluntarism
... is the view that moral principles have an objective foundation, and are not based on subjective human convention.The term in its broadest sense, applies to moral theories that emphasize the use of reason or a rational procedure in moral decision making*moral rationalism
... is the view that moral utterances are neither true nor false statements about the world. They are, instead, expressions of feelings or prescriptive utterances. The key to this issue is distinguishing between two types of utterances: (1) propositional utterances, and (2) nonpropositional utterances*noncognitivism
... lived at a critical juncture of western culture when the arrival of the Aristotelian corpus in Latin translation reopened the question of the relation between faith and reason, calling into question the modus vivendi that had obtained for centuries. This crisis flared up just as universities were being founded*Saint Thomas Aquinas
... maintained that knowledge comes from foundational concepts known intuitively through reason, such as innate ideas. Other concepts are then deductively drawn from these*Continental Rationalists
... names both a political theory of the legitimacy of political authority and a moral theory about the origin and/or legitimate content of moral norms. The political theory of authority claims that legitimate authority of government must derive from the consent of the governed, where the form and content of this consent derives from the idea of contract or mutual agreement. The moral theory of (this theory) claims that moral norms derive their normative force from the idea of contract or mutual agreement and is thus skeptical of the possibility of grounding morality or political authority in either divine will or some perfectionist ideal of the nature of humanity*Contractarianism
... refers to a class of normative moral theories which maintain that an action is morally right if the consequences of that action are more favorable than unfavorable. Thus, correct moral conduct is determined solely by a cost-benefit analysis of an action's consequences*Consequentialism
... refers to the 18th century philosophical movement in Great Britain which maintained that all knowledge comes from experience*British Empiricism
... refers to the view that the truth of a thing is independent from the observing subject. The notion entails that certain things exist independently from the mind, or that they are at least in an external sphere*objectivity
... traditionally refers to a 17th century philosophical movement begun by Descartes. After Descartes, several dozen scientists and philosophers continued his teachings throughout continental Europe and, accordingly were titled "Cartesians." Some Cartesians strayed little from Descartes' scientific and metaphysical theories. Others incorporated his theories into Calvinistic theology. But a handful of philosophers influenced by Descartes were more original in developing their own views and these people are included under this more more general title*Continental rationalism
... was a profound and prolific writer in the Danish "golden age" of intellectual and artistic activity. His work crosses the boundaries of philosophy, theology, psychology, literary criticism, devotional literature and fiction. Kierkegaard brought this potent mixture of discourses to bear as social critique and for the purpose of renewing Christian faith within Christendom. At the same time he made many original conceptual contributions to each of the disciplines he employed. He is known as the "father of existentialism", but at least as important are his critiques of Hegel and of the German romantics, his contributions to the development of modernism, his literary experimentation, his vivid re-presentation of biblical figures to bring out their modern relevance, his invention of key concepts which have been explored and redeployed by thinkers ever since, his interventions in contemporary Danish church politics, and his fervent attempts to analyse and revitalise Christian faith*Soren Aabye Kierkegaard
... was the founder of American pragmatism (which he called "pragmaticism"), an extender of the Scotistic theory of signs (which he called "semeiotic"), an extraordinarily prolific logician and mathematician, and a developer of an evolutionary, psycho-physically monistic metaphysical system. A practicing chemist and geodesist by profession, he nevertheless considered scientific philosophy, and especially logic, to be his vocation. In the course of his polymathic researches, he wrote on a wide range of topics, ranging from mathematical logic to psychology*Charles Sanders Peirce
..., British philosopher, economist, moral and political theorist, and administrator, was the most influential English-speaking philosopher of the nineteenth century. His views are of continuing significance, and are generally recognized to be among the deepest and certainly the most effective defenses of empiricism and of a liberal political view of society and culture. The overall aim of his philosophy is to develop a positive view of the universe and the place of humans in it, one which contributes to the progress of human knowledge, individual freedom and human well-being. His views are not entirely original, having their roots in the British empiricism of John Locke, George Berkeley and David Hume, and in the utilitarianism of Jeremy Bentham. But he gave them a new depth, and his formulations were sufficiently articulate to gain for them a continuing influence among a broad public*John Stuart Mill
Along with J. G. Fichte and F. W. J. Schelling, ... belongs to the period of "German idealism" in the decades following Kant. The most systematic of the post-Kantian idealists, he attempted, throughout his published writings as well as in his lectures, to elaborate a comprehensive and systematic ontology from a "logical" starting point. He is perhaps most well-known for his teleological account of history, an account which was later taken over by Marx and "inverted" into a materialist theory of an historical development culminating in communism*Georg Wilhelm Friedrich Hegel
An ... may be defined as an object that has been intentionally made or produced for a certain purpose. Often the word is used in a more restricted sense to refer to simple, hand-made objects (for example, tools) which represent a particular culture. (This might be termed the 'archaeological sense' of the word.) In experimental science, the expression is sometimes used to refer to experimental results which are not manifestations of the natural phenomena under investigation, but are due to the particular experimental arrangement*artifact
Generally regarded as the most important philosopher ever to write in English,  the last of the great triumvirate of "British empiricists", ... was also noted as an historian and essayist. A master stylist in any genre, his major philosophical works remain widely and deeply influential, despite their being denounced by many of his contemporaries as works of scepticism and atheism. While Hume's influence is evident in the moral philosophy and economic writings of his close friend Adam Smith, he also awakened Immanuel Kant from his "dogmatic slumbers" and "caused the scales to fall" from Jeremy Bentham's eyes. Charles Darwin counted him as a central influence, as did "Darwin's bulldog," Thomas Henry Huxley*David Hume
In its most general philosophical sense, a ... involves any object in the world around us that we perceive through our senses. It is that perception of an object which becomes visible to our consciousness*phenomenon
Philosophers are interested in a constellation of issues involving the concept of truth. A preliminary issue, although somewhat subsidiary, is to decide what sorts of things can be true. Is truth a property of sentences (which are linguistic entities in some language or other), or is truth a property of propositions (nonlinguistic, abstract and timeless entities)? The principal issue is: What is truth? It is the problem of being clear about what you are saying when you say some claim or other is true. The most important theories of truth are the ... ?*Correspondence Theory, the Semantic Theory, the Deflationary Theory, the Coherence Theory and the Pragmatic Theory
The ... theory claims that we perceive something called ... in place of the actual objects that are in the world around us. This concept was first introduced by Moore, and was later adopted by Russell and Broad. This theory has come under scrutiny from Ryle and Austin, who propose that the notion of ... only complicates our account of perceptions. We do not perceive discrete bits of information, but instead perceive objects in our surrounding world*sense-data
The concept of ... is that all derived or secondary things proceed or flow from the more primary. It is distinguished from the doctrine of creation by its elimination of a definite will in the first cause, from which all things are made to emanate according to natural laws and without conscious volition. It differs from the theory of formation at the hands of a supreme artisan who finds his matter ready to his hand, in teaching that all things, whether actually or only apparently material, flow from the primal principle*emanation
The exact point in time when the term ... was first adopted is unknown. It is, qhowever, certain that Italy and the re-adopting of Latin letters as the staple of human culture were responsible for the name of Humanists. Literoe humaniores was an expression coined in reference to the classic literature of Rome and the imitation and reproduction of its literary forms in the new learning; this was in contrast to and against the Literoe sacroe of scholasticism*humanism
The field of ethics, also called moral philosophy, involves systematizing, defending, and recommending concepts of right and wrong behavior. Philosophers today usually divide ethical theories into three general subject areas, ... ?*metaethics, normative ethics and applied ethics
The term ... is ambiguous. It refers to a type of moral theory, as well as to a type of legal theory, despite the fact that the core claims of the two kinds of theory are logically independent*natural law
The term ... refers to information obtained externally by means of the senses or internally through emotion. The term a posteriori is often used interchangeably*experience
The term ... was first used by Christian Wolff in his discussions of the mind-body problem to depict both philosophers who would only acknowledge the mind (idealism or mentalism) and philosophers who only acknowledged the body (materialism). The meaning Wolff originally intended by using the term has broadened in scope through the centuries, and today applies to any doctrine or theory that claims that all things, no matter how many or of what variety, can be reduced to one unified thing in time, space, or quality*monism
The term ... was originally coined by Thomas Hyde around the beginning of the eighteenth century. As a metaphysical theory, it states that the world is made up of two elemental categories which are incommensurable. This includes distinctions between mind and body, good and evil, universal and particular, and phenomena and noumena*dualism
Money isn't everything, says ... ?*my boss
Abortion is advocated only by persons who have themselves ... ?*been born
Anyone who goes to a psychiatrist ought to have his ... ?*head examined
Coffee isn't my ... ?*cup of tea
Color-television. I don't believe that, until I have seen it ... ?*black on white
Include me ... ?*out
I am not vicious. I am ... ?*scottish
Thank you for the compliment. But everything you see, I owe to ... ?*spaghetti
He was heavy-weight champion in boxing 1994 and known as Double M*Michael Moorer
He was heavy-weight champion in boxing from 1885-1892, and known as the Boston strong boy*John L. Sullivan
He was heavy-weight champion in boxing from 1892-1897, and known as Gentleman Jim*James J. Corbett
He was heavy-weight champion in boxing from 1897-1899, and known as Ruby Robert*Bob Fitzsimmons
He was heavy-weight champion in boxing from 1899-1905, and known as The boilermaker*James J. Jeffries
He was heavy-weight champion in boxing from 1905-1906, and known as The Fightin' Kentuckian*Marvin Hart
He was heavy-weight champion in boxing from 1906-1908, and known as The Little Giant of Hanover*Tommy Burns
He was heavy-weight champion in boxing from 1908-1915, and known as Lil' Arthur*Jack Johnson
He was heavy-weight champion in boxing from 1915-1919, and known as The Pottawatomie Giant*Jess Willard
He was heavy-weight champion in boxing from 1919-1926, and known as the Manassa Mauler*Jack Dempsey
He was heavy-weight champion in boxing from 1926-1928, and known as The fighting Marine*Gene Tunney
He was heavy-weight champion in boxing from 1930-1932, and known as The Black Uhlan of the Rhine*Max Schmeling
He was heavy-weight champion in boxing from 1932-1933, and known as The Boston Gob*Jack Sharkey
He was heavy-weight champion in boxing from 1933-1934, and known as The ambling Alp*Primo Carnera
He was heavy-weight champion in boxing from 1934-1935, and known as The Livermore Larruper*Max Baer
He was heavy-weight champion in boxing from 1935-1937, and known as The Cinderella Man*James J. Braddock
He was heavy-weight champion in boxing from 1937-1949, and known as The Black Bomber*Joe Louis
He was heavy-weight champion in boxing from 1949-1951, and known as The Cincinnatti Cobra*Ezzard Charles
He was heavy-weight champion in boxing from 1951-1952, and known as Jersey Joe*Joe Walcott
He was heavy-weight champion in boxing from 1952-1956, and known as The Brockton Blockbuster*Rocky Marciano
He was heavy-weight champion in boxing from 1956-1959, and from 1960-1962*Floyd Patterson
He was heavy-weight champion in boxing from 1959-1960*Ingemar Johansson
He was heavy-weight champion in boxing from 1962-1964, and known as Sonny*Charles Liston
He was heavy-weight champion in boxing from 1964-1970, and known as The Louisville Lip*Cassius Clay
He was heavy-weight champion in boxing from 1970-1973 and known as Smokin'*Joe Frazier
He was heavy-weight champion in boxing from 1973-1974 and from 1994-1997 and known as Big George*George Foreman
He was heavy-weight champion in boxing from 1974-1978 and from 1978-1979 and known as The Greatest*Muhammad Ali
He was heavy-weight champion in boxing from 1980-1985 and known as The Easton Assasin*Larry Holmes
He was heavy-weight champion in boxing from 1985-1988 and known as The Spinks Jinx*Michael Spinks
He was heavy-weight champion in boxing from 1988-1990 and known as Iron Mike*Mike Tyson
He was heavy-weight champion in boxing from 1990-1992 and from 1993-1994 and known as The real deal*Evander Holyfield
He was heavy-weight champion in boxing from 1992-1993 and known as Big Daddy*Riddick Bowe
He was heavy-weight champion in boxing from 1997-1998 and known as  The Cannon*Shannon Briggs
He was heavy-weight champion in boxing from 1998-2001 and known as The Lion*Lennox Lewis
He was heavy-weight champion in boxing in 1978 and known as Neon Leon*Leon Spinks
He was heavy-weight champion in boxing in 1990 and known as Buster*James Douglas
He was heavy-weight champion in boxing in 2001 and known as The Rock*Hasim Rahman
American Motion Pictures are written by the half-educated for the ... ?*half-witted
Women are always eagerly on the lookout for any ... ?*emotion
If you shoot at mimes, should you ... ?*use a silencer?
Logical consequences are the scarecrows of fools and the beacons of ... ?*wise men
Rather than love, than money, than fame, give me ... ?*truth
In 1909, which U.S. President became the first to be depicted on a coin?*Abraham Lincoln
What the W. of George W. Bush stand for?*Walker
The number unemployed consists of all those people in a country who are willing and able to work but are unable to ... ?*find jobs
Which Vice President was the only one to serve two full terms as President?*Thomas Jefferson
Who was the first Vice President to resign?*John C. Calhoun
I am free of all prejudice. I ... ?*hate everyone equally
Who removed the cork from my ... ?*lunch
Men will pay large sums to whores, for telling them they are not ... ?*bores
The only sure weapon against bad ideas is ... ?*better ideas
She tried to sit in my lap, while I was ... ?*standing
It is the vice of scholars to suppose that there is no knowledge in the world but that of ... ?*books
What if everything is an illusion and nothing exists? In that case, I definitely ... ?*overpaid for my carpet
I never hated a man so much, that I returned his ... ?*diamants
Britains say 'tarmac'; Americans say ----------*runway
Analogy: 'Ancient' is to 'old' as 'recent' is to ----------*current
Bull - cow as fox - ----------*vixen
Goose - geese as passerby - ----------*passersby
Your ---------- holds your head to your shoulders*neck
In what field of study would you find "flying buttresses"?*architecture
The study of building design is ----------*architecture
This statue was found on the Greek island of Melos in 1820*Venus de Milo
Three main types of Greek columns are Doric, Ionic, and ----------*Corinthian
What Dutch master painted 64 self-portraits?*Rembrandt
Where is the Louvre located?*Paris
Who painted the Mona Lisa?*Leonardo da Vinci
What is the only sign in the zodiac which doesn't represent a living thing?*Libra
Which month has a diamond as a birthstone?*April
A heavenly body moving under the attraction of the Sun and consisting of a nucleus and a tail is a(n) ----------*comet
Does Uranus have an aurora?*yes
From 1979 until 2000 the most distant planet from the earth was ----------*Neptune
How many planets are there in our solar system?*nine
If you're in the northern hemisphere, Polaris, the North Star, can be found by looking which direction?*north
Name the largest planet in the solar system*Jupiter
Name the second-largest planet in the solar system*saturn
Our galaxy is commonly known as the ----------*Milky Way
The Big Dipper is part of what constellation?*Ursa Major
The North Star is also known as ----------*Polaris
The fourth planet from the sun is ----------*Mars
The name for the group of stars which form a hunter with a club and shield is ----------*Orion
The planet closest to the sun is ----------*Mercury
The spiral galaxy nearest ours is the ---------- galaxy*Andromeda
The tides on the earth's oceans are actually created by gravitational pull from the ----------*Moon
This cluster of stars is also known as the Seven Sisters*Pleiades
This comet appears every 76.3 years*Comet Halley
This planet's diameter is most equal to that of the earth's*Venus
What does "Ursa Major" mean in everyday English?*great bear
What is the astronomical name for a group of stars?*constellation
What is the name for the theoretical end-product of the gravitational collapse of a massive star?*black hole
What is the name used to describe the "minor planets"?*asteroids
What is the ocean of air around the earth called?*atmosphere
What is the proper name for falling stars?*meteors
What is the term for the path followed a by a small body around a massive body in space?*orbit
What phenomenon is caused by the gravitational attraction of the moon?*tides
What planet boasts the Great Red Spot?*Jupiter
The two sexes of humans are male and ----------?*female
This poisonous, oily liquid occurs in tobacco leaves*nicotine
This protein makes the blood red in color*Haemoglobin
Every human has one of these on their tummies*navel
How many large holes are in your head?*seven
This complex substance makes up all living things*protoplasm
What is normal body temperature for an adult human (in degrees fahrenheit)?*98.6
The practice of joining the parts of two plants to make them grow as one is called ----------*grafting
These flowerless plants grow on bare rocks and tree stumps*lichen
This fruit has its seeds on the outside*strawberry
This term means 'cone-bearing trees'*conifers
What fruit bear the latin name "citrus grandis"?*grapefruit
How many days where there in 1976?*three hundred and sixty six
Alvin & Simon had a brother called ----------*Theodore
An Andy Panda cartoon gave birth to a famous, cantankerous bird. Name him*Woody Woodpecker
An adventurous penguin named Tennessee Tuxedo had a sidekick named ----------?*Chumley
An alien creature in a funny hat has opposed both Bugs Bunny and Daffy Duck.  Where is he from?*Mars
Before Olive Oil met Popeye she was engaged to someone.  Who was he?*Ham Gravy
Benny and Cecil were at odds with whom?*John
Bugs always finds himself at the wrong end of a gun, usually toted by either Elmer Fudd or who?*Yosemite Sam
Casper the Friendly Ghost frolicked with which witch?*Wendy
Charles Boyer inspired a cartoon skunk. Who?*Pepe le Pew
Famous Phrases: Who knows? The ----------*Shadow
Gadzookie has a large, green friend. Who is he?*Godzilla
Hanna-Barbera rose to fame by creating what duo for MGM?*Tom and Jerry
How does Wonder Woman control her invisible airplane?*mental powers
In the cartoons who was Hokie Wolf's sidekick?*Ding
In what city does Fat Albert live?*Philadelphia
Mentor of Titan had two children in the Marvel comics, Thanos and ----------?*Ero
Miss Buckley is secretary to what commanding officer?*General Halftrack
Name Alley Oop's girl friend*Oola
Name Cathy's on again/off again boy friend?*Irving
Name Dennis the Menace's next door neighbors*Mr and Mrs Wilson
Name Donald Duck's girlfriend?*Daisy
Name Hagar the Horrible's dog*Snert
Name Li'l Abner's favorite Indian drink*Kickapoo Joy Juice
Name the European hit, now an animated series about underwater people*The Snorks
Name the apartments the Jetson's live in*The Skypad Apartments
Name the dog in the Yankee Doodle cartoons*Chopper
Name the fastest mouse in all of Mexico*Speedy Gonzalez
Name the ranger who was always after Yogi Bear*Rick
Name the town that Fred, Wilma, Barney, and Betty lived in*Bedrock
On what T.V. show could Tom Terrific be found?*Captain Kangaroo
Popeye's chief adversary has two names, Bluto and ----------?*Brutus
Porky Pig had a girlfriend named ----------*Petunia
Tess Trueheart married which plainclothes detective?*Dick Tracy
What came out of Milton's head?*steam
What character did Tex Avery first create upon arriving at MGM?*Screwball Squirrel
What comic strip character is Beetle Bailey's sister?*Lois (of Hi and Lois)
What did Dagwood give up to marry Blondie?*A family inheritance
What did Peppermint Patty always call Charlie Brown?*Chuck
What is Batman's butler Alfred's last name*Pennyworth
What is Blondie's maiden name?*Oop
What is Dennis the Menace's last name?*Mitchell
What is Smokey Stover's job?*fireman
What is Super Chicken's partners name?*Fred
What is the mother's name in Family Circus?*Thelma
What is the name of Duddley Do-Right's horse?*Horse
What is the name of the Family Circus's dog?*Barf
What kind of dog is Scooby Doo?*great dane
What type of plant does Broom Hilda sell?*venus flytrap
What was Daffy Duck's favorite insult?*You're dispicable!
What was George of the Jungle always running in to?*A tree
What was the first cartoon to feature sound?*Steamboat Willie
What was the name of George of the Jungle's pet elephant?*Shep
What was the name of Speed Racer's car?*The Mach Five
What was the original name Charles Schultz had for Peanuts?*Li'l Folks
What was the relationship between Superman and Supergirl?*cousin
When Tweety exclaimed, "I thought I saw a putty tat!", who did he see?*Sylvester
When danger appeared, Quick Draw McGraw became which super hero?*El KaBong
When not a Birdman, what does Ray Randall do for a living?*police officer
When not fighting crime, what did Underdog do for a living?*shoeshine boy
Where are Rocket J. Squirel and Bullwinkle Moose from?*Frostbite Falls
Where did Clark Kent attend college?*Metropolis University
Where did George of the Jungle live?*Imgwee Gwee Valley
Where did Mighty Mouse get his superpowers?*supermarket
Where do Rocky and Bullwinkle play football?*What'samatta University
Where does George Jetson work?*Spacely Sprockets
Where does Yogi Bear Live?*Jellystone Park
Which comic strip was banned from "Stars and Stripes"?*Beetle Bailey
Which superhero loves peace enough to kill for it?*Peacemaker
Who always tried to kill Krazy Kat?*Captain Marvel
Who is Donald Duck's uncle?*Scrooge
Who is Sally Brown's sweet baboo?*Linus
Who is Scooby Doo's nephew??*Scrappy Doo
Who is Snoopy's arch enemy?*The Red Baron
Who is stationed at Camp Swampy in the comic strips?*Beetle Bailey
Who runs Andy Capp's favorite pub?*Jack and Jill
Who says, "Th-th-th-that's all folks!"*Porky Pig
Who shot Bruce Wayne's parents?*Chill
Who was Dick Dastardly's pet?*Muttley
Who was always trying to get rent from Andy Capp?*Percy
Who was the Hulk's first friend?*Rick Jones
Who was the original voice of Mickey Mouse?*Walt Disney
Hydrogen Hydroxide is more commonly known as what?*water
Nitrogen, a poisonous gas, makes up 78% of the ---------- that we breathe*air
Sodium Hydroxide is more commonly known as ----------*lye
Sodium bicarbonate is better known as ----------*baking soda
The process of removing salt from sea water is known as ----------*desalination
The smallest portion of a substance capable of existing independently and retaining its original properties is a(n) ----------*molecule
This ancient attempt to transmute base metals into gold was called ----------*alchemy
This is the heaviest naturally occurring element*uranium
This is the symbol for tin*Sn
This poisonous gas is in the exhaust fumes from cars*carbon monoxide
Water containing carbon dioxide under pressure is called ---------- ----------*soda water
What is it that turns blue litmus paper red?*acid
What is the abbreviation for trinitrotoluene?*TNT
What is the chemical symbol for copper?*Cu
What is the chemical symbol for gold?*Au
What is the main component of air?*nitrogen
What is the more scientific name for quicksilver?*mercury
What is the symbol for iron in chemistry?*Fe
What is the symbol for silver?*Ag
Name the loner rebel reindeer with the red shiny nose*Rudolph
Santa Claus reportedly lives at the ---------- Pole*north
What comic strip is set at Camp Swampy?*Beetle Bailey
That big square thing you're staring at right now is called a ----------?*monitor
The Internet Relay Chat Program, which normally connects to port 6667, is more commonly known as ----------*mIRC
--isms:  Exalting one's country above all others*nationalism
--isms:  The belief that there is no God*atheism
--isms:  The theory that man cannot prove the existence of a god*agnosticism
-ism: The belief in the existence of a god or gods*theism
-isms: A diseased condition resulting from the use of beverages such as whiskey*alcoholism
-isms: A one-party system of government in which control is maintained by force and regimentation*fascism
-isms: A painful stiffness of the muscles and joints*rheumatism
-isms: A severe or unfavorable judgment*criticism
-isms: An economic system characterized by private ownership and competition*capitalism
-isms: Excessive emphasis on financial gain*commercialism
-isms: Excessive enthusiasm or zeal for a cause*fanaticism
-isms: Poisoning caused by a toxin in improperly prepared food*botulism
-isms: Public ownership of the basic means of production, distribution, and exchange*socialism
-isms: The belief in living a very austere and self-denying life*asceticism
A clip, shaped like a bar to keep a woman's hair in place is a ----------*barrette
A depilatory is a substance used for removing ----------*hair
A device used to change the voltage of alternating currents is a ----------*transformer
A flat, round hat sometimes worn by soldiers is a ----------*beret
A government in which power is restricted to a few is a(n) ----------*oligarchy
A person in his eighties is called a(n) ----------*octogenarian
A person who starts fires maliciously is a(n) ----------*arsonist
A person with a strong desire to steal is a(n) ----------*kleptomaniac
A pugilist is a ----------*boxer
A receptacle for holy water is a(n) ----------*font
A sun-dried grape is known as a(n) ----------*raisin
A word like 'NASA' formed from the initials of other words is a(n) ----------*acronym
Acrophobia is a fear of ----------*heights
An "omniscient" person has unlimited ----------*knowledge
An anemometer measures ---------- ----------*wind velocity
An animal stuffer is a(n) ----------*taxidermist
An artist supports his canvas on a(n) ----------*easel
An instrument on a car to measure the distance travelled is called a(n) ----------*odometer
Androphobia is the fear of ----------*males
Animals that once existed but don't exist now are said to be ----------*extinct
Any object worn as a charm may be called a(n) ----------*amulet
Bibliophobia is a fear of ----------*books
Brontophobia is the fear of ----------*thunder
Dendrochronology is better known as ----------*ring dating
Doraphobia is the fear of ----------*fur
Eleutherophobia is a fear of ----------*freedom
Gymnophobia is the fear of ----------*naked bodies
Hills and ridges composed of drifting sand are known as ----------*dunes
Ichthyology is the study of ----------*fish
Legal Terms: A crime more serious than a misdemeanor*felony
Legal Terms: A formal agreement enforceable by law*contract
Legal Terms: A supplement to a will*codicil
Legal Terms: The people chosen to render a verdict in a court*jury
Legal Terms: To steal property entrusted to one's care*embezzle
Name the pain-inflicting person you go to to get your teeth fixed*dentist
Name the porceilan chair you sit on at least once a day*toilet
One who tells fortunes by the stars is a(n) ----------*astrologer
Rats, mice, beavers, and squirrels are all ----------*rodents
The art of tracing designs and taking impressions of them is ----------*lithography
The covering on the tip of a shoelace is a(n) ----------*aglet
The distance around the outside of a circle is its ----------*circumference
The earth's atmosphere and the space beyond is known as ----------*aerospace
The effect produced when sound is reflected back is known as a(n) ----------*echo
The feeling of having experienced something before is known as ----------*deja vu
The science of preparing and dispensing drugs is ----------*pharmacy
The science of providing men, equipment and supplies for military operations is called ----------*logistics
The study of human behaviour is ----------*psychology
The study of human pre-history is ----------*archaeology
The study of insects is ----------*entomology
The study of light and its relation to sight is called ----------*optics
The study of man and culture is known as ----------*anthropology
The study of natural phenomena: motion, forces, light, sound, etc. is called ----------*physics
The study of plants is ----------*botany
The study of religion is ----------*theology
The study of sound is ----------*acoustics
The study of the composition of substances and the changes that they undergo is ----------*chemistry
The study of the earth's physical divisions into mountains, seas, etc. is ----------*geography
The study of the manner in which organisms carry on their life processes is ----------*physiology
The symbols used on a map are explained by the ----------*legend
The treatment of disease by chemical substances which are toxic to the causative micro-organisms is called ----------*chemotherapy
The weight at the end of a pendulum is a(n) ----------*bob
The wide wall built along the banks of rivers to stop flooding is a(n) ----------*levee
The word "cumulus" refers to a type of ----------*cloud
This instrument is used for measuring the distance between two points, on a curved surface*calliper
This instrument measures atmospheric pressure*barometer
This instrument measures the velocity of the wind*anemometer
This is the fear of enclosed spaces*claustrophobia
This word is used as the international radio distress call*mayday
This word means "split personality"*schizophrenia
Throat, foxing, and platform are parts of a(n) ----------*shoe
Toxiphobia is a fear of ----------*poison
What does a brandophile collect?*cigar bands
What does a heliologist study?*the sun
What does an ornithologist study?*birds
What is a device to stem the flow of blood called?*tourniquet
What is a dried plum called?*prune
What is a female swan called?*pen
What is a figure with eight equal sides called?*octagon
What is another name for a tombstone inscription?*epitaph
What is measured by a chronometer?*time
What is the common name for a Japanese dwarf tree?*bonsai
What is the common name for the Aurora Borealis?*northern lights
What is the common term for a "somnambulist"?*sleepwalker
What is the name for a branch of a river?*tributary
What is the study of heredity called?*genetics
What is the study of prehistoric plants and animals called?*paleontology
What is the term for a castrated rooster?*capon
Which science studies weather?*meteorology
Name the clothing wrinkle remover that sounds like a kind of metal*iron
A 3 1/2" floppy disk measures ---------- and 1/2 inches across*three
A brown crayon is what color?*brown
A smoke detector will alarm if it detects ----------*smoke
A water heater keeps ---------- warm for you*water
Although not all come from France, ---------- fries are often served with hamburgers*french
An ---------- clock usually wakes you in the morning*alarm
During the American Revolution, the Boston Tea Party took place in ---------- Harbor*boston
How do you spell abbreviation?*abbreviation
How many pencils are there in a dozen?*twelve
"Our Town" is a play by whom?*Thornton Wilder
The play "Our Town" is set where?*Grover's Corners
What is a sorcerer who deals in black magic called?*necromancer
Boston butt, jowl, and picnic ham are parts of a ----------*pig
Even though it tastes nothing like grapes, a ---------- is often eaten for breakfast*grapefruit
From what animal do we get venison?*deer
From which fish is caviar obtained?*sturgeon
From which fruit is the liqueur Kirsh made?*cherry
Iceberg, Boston, and Bibb are types of ----------*lettuce
Laetrile is associated with the pit of which fruit?*apricot
Little round chocolate candies are known as ----------&m's*m
Mustard, ketchup and onions on a hotdog are all ----------*condiments
Name the only fruit named for its color*orange
Natural vanilla flavoring comes from this plant*orchid
Often drunk, this liquid is normally harvested from female cows*milk
Often eaten for breakfast, bacon is actually the flesh of what barnyard animal?*pig
Often eaten for breakfast, the egg comes from what barnyard animal?*chicken
Rum is made from this plant*sugar cane
Vermicelli literally means ----------*little worms
What do you get when you add fresh fruit to red wine?*sangria
What is Japanese "sake" made from?*rice
What is the name of the syrup drained from raw sugar?*molasses
What kind of nuts are used in marzipan?*almonds
Where is the best brandy bottled?*cognac
Where was Budweiser first brewed?*St. Louis
Cocktails: Bourbon, sugar and mint make a(n) ----------*mint julep
Cocktails: Cognac (brandy) and white creme de menthe make a(n) ----------*stinger
Cocktails: Creme de Cacao, cream, and brandy make a(n) ----------*brandy alexander
Cocktails: Gin and Collins mix make a(n) ----------*Tom Collins
Cocktails: Rum, lime, and cola drink make a(n) ----------*cuba libre
Cocktails: Triple sec, tequila, and lemon or lime juice make a(n) ----------*margarita
Cocktails: Vodka and Kahlua make a ----------*black russian
Cocktails: Vodka, orange juice and Galliano make a(n) ----------*Harvey Wallbanger
Cocktails: Whiskey, hot coffee, and whipped cream make a(n) ----------*Irish coffee
A bridge hand with no cards in one suit is said to have a ----------*void
A poker hand consisting of three of a kind and a pair is called a ----------*full house
How many balls are used in a game of snooker including the cue ball?*twenty two
How many dots are there on a pair of dice?*forty two
How many squares are there on a chessboard?*sixty four
If you "peg out" what game are you playing?*cribbage
In poker five cards of the same suit is called a(n) ----------*flush
In pool, what color is the eight ball?*black
In which game might a person have a "full house"?*poker
In which game or sport are "Staunton" pieces used?*chess
In which game or sport can a person be "skunked"?*cribbage
In which sport are terms "spare" and "gutter" used?*tenpin bowling
In which sport or game are the terms: 'pin', 'fork', and 'skewer' used?*chess
In which sport or game is the term "rook" used?*chess
Name the only flexible murder weapon in the game of "Cluedo"*rope
Name the only woman suspect in the game of "Cluedo" who isn't married*Miss Scarlett
These are the two highest valued letters in "Scrabble".  "Q" and ----------*Z
This ancient Chinese game is played with 156 small rectangular tiles*mah jongg
This is the lowest ranking suit in Bridge*clubs
This term denotes a chess move in which both the king and the rook are moved*castling
What bowling term means three straight strikes?*turkey
What game or sport is Bobby Fischer identified with?*chess
What number is on the opposite side of the "five" on dice?*two
What score is not possible for a cribbage hand?*nineteen
Which chess piece is usually valued as 5 points?*rook
Which game usually begins with, "Is it animal, vegetable, or mineral?"*twenty questions
Although it doesn't sound like a dog, ----------dust is ornamental wood chips often placed in flowerbeds*bark
Chicago Transit Authority is now known as which group?*Chicago
If you drive on a parkway, you park on a ----------?*driveway
To refuel your car you go to a ---------- station*petrol
Wedding rings are normally worn on what finger of your hand?*ring finger
What are catalogued under the Dewey decimal system?*books
What are the 3 big colleges of the Ivy League? (name them alphabetically)*Harvard, Princeton, Yale
What company makes Pampers disposable diapers?*Proctor & Gamble
What do the initials U.F.O stand for?*Unidentified Flying Object
What do the letters in SAM missiles refer to?*Surface-to-Air Missile
What does IRS stand for?*Internal Revenue Service
What does a compass needle point to?*north
What does the abbreviation a.m. stand for?*Ante Meridian
What does the acronym CIA stand for?*Central Intelligence Agency
What is the minimum IQ score for the genius category?*one hundred and forty
What is this sign called "&"?*ampersand
When using a telephone, you must wait for a ---------- tone before starting your call*dial
Which U.S. president is on the five-dollar bill?*Abraham Lincoln
Acadia was the original name of which Canadian province?*Nova Scotia
Bridgeport is the largest city in which state?*Connecticut
Bridgetown is the capital of ----------*Barbados
Brussels is the capital of which country?*Belgium
Frankfort is the capital of which state?*Kentucky
Guayaquil is the largest city in what country?*Ecuador
Halifax is the capital of which Canadian province?*Nova Scotia
Havana is the capital of which country?*Cuba
He invented the most common projection for world maps*Gerardus Mercator
He visited Australia and New Zealand, then surveyed the Pacific Coast of North America*Vancouver
How many stars are on the flag of New Zealand?*four
If its 4:00pm in Seattle, Washington what time is it in Portland, Oregon?*4:00pm
In what city are the famous Tivoli Gardens?*Copenhagen
In what city is the Leaning Tower?*Pisa
In what city is the Smithsonian Institute?*Washington
In what country is Banff National Park?*Canada
In what country is Lahore?*Pakistan
In what country is Mandalay?*Myanmar (formerly known as Burma)
In what country is Taipei?*Taiwan
In what country is Thunder Bay?*Canada
In what country is the Jutland peninsula?*Denmark
In what country is the Mekong River Delta?*Vietnam
In what country is the Waterloo battlefield?*Belgium
In what country is the highest point in South America?*Argentina
In what country is the lowest point in South America?*Argentina
In what country is the source of the Blue Nile?*Ethiopia
In what island group is Corregidor?*Philippines
In what mountain range is Kicking Horse Pass?*Rocky
In what state is Concord?*New Hampshire
In which city is Red Square?*Moscow
In which city is Saint Paul's Cathedral?*London
In which city is Wembley Stadium?*London
In which city is the Bridge of Sighs?*Venice
In which city is the C.N. Tower?*Toronto
In which city is the Canale Grande?*Venice
In which city is the Colliseum located?*Rome
In which city is the Wailing Wall?*Jerusalem
In which country is Angel Falls?*Venezuela
In which country is Brest? (NOT Breast!)*France
In which country is Chennai (formerly Madras)?*India
In which country is Cusco?*Peru
In which country is Loch Ness?*Scotland
In which country is Normandy?*France
In which country is Sapporo?*Japan
In which country is the Calabria region?*Italy
In which country is the Dalai Lama's palace?*Tibet
In which country is the Great Victoria Desert?*Australia
In which country is the Machu Picchu?*Peru
In which country would you find the Yucatan Peninsula?*Mexico
In which ocean or sea are the Seychelles?*Indian Ocean
In which state are Gettysburg and the Liberty Bell?*Pennsylvania
In which state are the Finger Lakes?*New York
In which state is Appomattax?*Virginia
In which state is Cape Hatteras?*North Carolina
In which state is Hoover Dam?*Arizona
In which state is Mount McKinley?*Alaska
In which state is Mount St. Helens?*Washington
In which state is Mount Vernon?*Virginia
In which state is Stone Mountain?*Georgia
In which state is Walla Walla?*Washington
In which state is the Kennedy Space Center?*Florida
In which state is the Mayo Clinic?*Minnesota
In which state is the Painted Desert?*Arizona
Into what bay does the Ganges River flow?*Bay Of Bengal
Into what body of water does the Danube River flow?*Black Sea
Into what body of water does the Yukon River flow?*Bering Sea
Into what sea does the Elbe River flow?*North Sea
Into what sea does the Mackenzie River flow?*Beaufort Sea
Khartoum is the capital of which country?*Sudan
Kingston is the capital of which country?*Jamaica
Linz, Austria is a leading port on which river?*Danube
Madrid and Lisbon are both located near this river*Tagus
Meridians converge at the ----------*poles
Mount Victoria is the highest peak of this island country*Fiji
Name a country which has the same name as a bird*Turkey
Name the U.S. state with the smallest population*Alaska
Name the capital city of Massachusetts*Boston
Name the capital city of Rhode Island*Providence
Name the capital of Argentina*Buenos Aires
Name the capital of Brazil*Brasilia
Name the capital of Italy*Rome
Name the city at the west end of Lake Superior*Duluth
Name the continent that consists of a single country*Australia
Name the desert located in south-east California*Mojave
Name the large mountain chain in the eastern U.S.A*The Appalachians
Name the largest cathedral in the world*St. Peter's
Name the largest city in Canada*Toronto
Name the largest island in the world*Greenland
Name the largest lake in Australia*Eyre
Name the largest river forming part of the U.S. - Mexican border*Rio Grande
Name the last province to become part of Canada*Newfoundland
Name the longest river in Asia*Yangtze
Name the longest river in Nigeria*Niger
Name the most north-easterly of the 48 contiguous states*Maine
Name the only Central American country without an Atlantic coastline*El Salvador
Name the sea between Asia Minor and Greece*Aegean
Name the sea between Korea and China*Yellow Sea
Name the sea north of Alaska*Beaufort
Name the sea north of Murmansk, Russia*Barents
Name the sea west of Alaska*Bering
Name the second largest country in Africa*Algeria
Name the second largest country in South America*Argentina
Name the second largest lake in North America*Huron
Name the smallest of the Great Lakes*Ontario
Name the strait joining the Atlantic Ocean and the Mediterranean Sea*Gibraltar
Name the world's most photographed and most climbed mountain*Fuji
Nassau is the capital of which country?*Bahamas
On what island is Honolulu?*Oahu
On what island is Pearl Harbour?*Oahu
On what island is the Blue Grotto?*Capri
On what island is the U.S. naval base, Guantanamo?*Cuba
On what mountain are four presidents' faces carved?*Rushmore
On what peninsula are Spain and Portugal located?*The Iberian peninsula
On what river is the capital city of Canada?*Ottawa
On which river is London, England?*Thames
On which river is Rome located?*Tiber
On which river is the Aswan High Dam?*Nile
Over 75% of the Earth's surface is covered by some form of ----------*Water
Rabat is the capital of which country?*Morocco
San Francisco Bay is located near what city?*San Francisco
Seoul is the capital of which country?*South Korea
St. George's is the capital city of what island country?*Grenada
Surfing is believed to have originated here*Hawaii
The Auckland Islands belong to which country?*New Zealand
The Hebrides are part of this country*Scotland
The Ionian and Cyclades are island groups of which country?*Greece
The Little Mermaid is found in the harbour of which city?*Copenhagen
The Nationalist Chinese occupy this island*Taiwan
The Thatcher Ferry Bridge crosses what canal?*Panama Canal
The United States is made up of ---------- states*fifty
The Volta is the largest river in which country?*Ghana
The countries of Belgium, Netherlands, and Luxembourg are together called ----------*Benelux
The longest river in Western Europe is ----------?*Rhine
The sun sets in the ----------?*west
This Canadian island is the world's fifth largest*Baffin
This Pacific island's puzzling monoliths attract ethnologists*Easter Island
This country is divided at the 38th parallel*Korea
This country occupies the "horn of Africa"*Somalia
This country's flag has a large "R" on it*Rwanda
This imaginary line approximately follows the 180 degree meridian through the Pacific Ocean*international date line
This is called the "Honeymoon Capital" of the world*Niagara Falls
This is the bridge with the longest span in the U.S.A*Verrazano Narrows
This is the only borough of New York City that is not on an island*The Bronx
This is the port city serving Tokyo*Yokohama
This is the residence of English monarchs*Buckingham Palace
This island group is off the east coast of southern South America*Falkland Islands
This re-opened in 1975 after being closed for 8 years*Suez Canal
This section of Manhattan is noted for its Negro and Latin American residents*Harlem
To what country do the Faeroe Islands belong?*Denmark
To what country does the Gaza Strip belong?*Egypt
True Or False:  The Easter Bunny is from Easter Island*false
True Or False:  There are only virgins on the Virgin Islands*false
Under what river does the Holland Tunnel run?*Hudson
Warsaw is the capital of what country?*Poland
What American city is known as Little Havana?*Miami
What Asian city was once called Edo?*Tokyo
What Canadian city is at the west end of Lake Ontario?*Hamilton
What European country administers the island of Martinique?*France
What European country has "Vaduz" as its capital city?*Liechtenstein
What London borough does the Prime Meridian pass through?*Greenwich
What U.S. city is known as Insurance City?*Hartford
What U.S. city is named after Saint Francis of Assisi?*San Francisco
What U.S. state is known as The Land of 10,000 Lakes?*Minnesota
What US state is completely surrounded by the Pacific Ocean?*Hawaii
What are drumlins and eskers formed by?*glaciers
What are the worlds four oceans - alphabetically?*Arctic, Atlantic, Indian and Pacific
What body of water borders Saudi Arabia to the east?*Persian Gulf
What canal connects Lake Ontario and Lake Erie?*Welland
What city boasts the Copacabana Beach and Ipanema?*Rio de Janeiro
What city is associated with Alcatraz?*San Francisco
What city is on Lake Erie at the mouth of the Cuyahoga River?*Cleveland
What city is the Christian Science Monitor based in?*Boston
What city is the Kremlin located in?*Moscow
What city was the setting for "Gone With the Wind"?*Atlanta
What color does the bride wear in China?*red
What continent is Cyprus considered to be part of?*Asia
What continent is the home to the greatest number of countries?*Africa
What country are the Islands of Quemoy and Matsu part of?*Taiwan
What country borders Egypt on the west?*Libya
What country borders Egypt to the south?*Sudan
What country does the island of Mykonos belong to?*Greece
What country formed the union of Tanganyika and Zanzibar?*Tanzania
What country has the world's most southerly city?*Chile
What country is Phnom Penh the capital of?*Cambodia
What country is Santo Domingo the capital of?*Dominican Republic
What country is directly north of Israel?*Lebanon
What country is directly north of the continental United States?*Canada
What country is directly west of Spain?*Portugal
What country is known as the Hellenic Republic?*Greece
What country is located between Panama and Nicaragua?*Costa Rica
What country owns the island of Corfu?*Greece
What country was once known as Gaul?*France
What country was the setting for "Casablanca"?*Morocco
What country was the setting for "Doctor Zhivago"?*Russia
What famous geyser erupts regularly at the Yellowstone National Park?*Old Faithful
What is the capital of Alaska?*Juneau
What is the capital of Bangladesh?*Dacca
What is the capital of Burma?*Rangoon
What is the capital of Chile?*Santiago
What is the capital of Colorado?*Denver
What is the capital of Florida?*Tallahassee
What is the capital of India?*New Delhi
What is the capital of Indonesia?*Jakarta
What is the capital of Kansas?*Topeka
What is the capital of Maine?*Augusta
What is the capital of Nebraska*Lincoln
What is the capital of New Zealand?*Wellington
What is the capital of Washington state?*Olympia
What is the capital of West Virginia?*Charleston
What is the capital of Wisconsin?*Madison
What is the capital of Zimbabwe?*Harare
What is the capital of the United Arab Emirates?*Abu Dhabi
What is the current name for south-west Africa?*Namibia
What is the highest mountain in Canada?*Mt. Logan
What is the largest city in Australia, in terms of population?*Sydney
What is the largest city in China?*Shanghai
What is the largest of the countries in Central America?*Nicaragua
What is the monetary unit of India?*Rupee
What is the most sacred river in India?*Ganges
What is the official language of Egypt?*Arabic
What is the principal river of Ireland?*Shannon
What is the second largest of the United States?*Texas
What is the smallest Canadian province?*Prince Edward Island
What is the smallest independent state in the world?*Vatican City
What is the smallest of the Central American countries?*El Salvador
What is the world's highest city?*Lhasa
What island has Hamilton as its capital?*Bermuda
What island is known as the Spice Island?*Zanzibar
What mountain range separates Europe from Asia?*Ural
What prison island was off the coast of French Guiana?*Devil's Island
What river has the largest drainage basin?*Amazon
What river is Liverpool on?*Mersey
What river is called "Old Man River"?*Mississippi
What river is known as China's Sorrow?*Yellow
What river is the Temple of Karnak near?*Nile
What sea is between Italy and Yugoslavia?*Adriatic
What south American country has both a Pacific and Atlantic coastline?*Colombia
What state borders Alabama to the north?*Tennessee
What state is the Golden State?*California
What state was the home to Mayberry?*North Carolina
What symbol is on the flag of Vietnam?*star
What unit of currency will buy you dinner in Iraq, Jordan, Tunisia, and Yugoslavia?*Dinar
What volcano showers ash on Sicily?*Etna
What's the former name of Istanbul?*Constantinople
What's the highest mountain in the 48 contiguous U.S. states?*Whitney
Where are the Nazca Lines?*Peru
Where are the pyramids located?*Egypt
Where is Beacon Street?*Boston
Where is Euston Station?*London
Where is George Washington buried?*Mt. Vernon, Virginia
Where is Queen Maud Land located?*Antarctica
Where is Westminster Abbey located?*London
Where is the Admirality Arch?*London
Where is the Holy Kaaba?*Mecca
Where is the Parthenon located?*Athens
Where is the city of Brotherly Love?*Philadelphia
Where were the Pillars of Hercules located?*Gibraltar
Which Canadian province extends farthest north?*Quebec
Which Central American country extends furthest north?*Belize
Which European country has the highest population density?*Monaco
Which European country has the lowest population density?*Iceland
Which Irish city is famous for its crystal?*Waterford
Which U.S. city is known as Beantown?*Boston
Which U.S. city is known as the Biggest Little City in the World?*Reno
Which U.S. state borders a Canadian territory?*Alaska
Which U.S. state has the least rainfall?*Nevada
Which U.S. state receives the most rainfall?*Hawaii
Which capital is known as the Glass Capital of the World?*Toledo
Which city has the largest rodeo in the world?*Calgary
Which city is known as Motown?*Detroit
Which city is known as the Windy City?*Chicago
Which city is on the east side of San Francisco Bay?*Oakland
Which country administers Christmas Island?*Australia
Which country are the Galapagos Islands part of?*Ecuador
Which country borders Italy, Switzerland, West Germany, Czechoslovakia, Hungary, Yugoslavia, and Liechtenstein?*Austria
Which country developed "Tae-Kwan-Do"?*Korea
Which country has Ankara as its capital?*Turkey
Which country has Budapest as its capital?*Hungary
Which country has the longest land border?*China
Which country hosted the 1982 World Cup of soccer?*Spain
Which country uses the "yen" for currency?*Japan
Which country would come first in an alphabetical list of countries?*Afghanistan
Which islands were named after Prince Philip of Spain?*The Philippines
Which mainland Latin American country is in neither South America nor Central America?*Mexico
Which of the 48 contiguous states extends farthest north?*Minnesota
Which of the U.S. states borders only one other state?*Maine
Which river contains the most fresh water?*Amazon
Which state has the most hospitals?*California
Which state is divided into two parts by a large lake?*Michigan
Which state is the Evergreen State?*Washington
Which state is the Garden State?*New Jersey
Which state is the Wolverine State?*Michigan
With what country is Fidel Castro associated?*Cuba
With which country is Prince Rainier III identified?*Monaco
Name the capital city of Utah*Salt Lake City
Earth's outer layer of surface soil or crust is called the ----------*Lithosphere
Peat, lignite and bituminous are types of ----------*coal
Slate is formed by the metamorphosis of ----------*shale
The green variety of beryl is called ----------*emerald
The molten material from a volcano is ----------?*lava
The spot on the Earth's surface directly above an earthquake's focus is called the ----------*epicenter
The violet variety of quartz is called ----------*amethyst
There are three types of rocks: metamorphic, sedimentary, and ----------*igneous
These limestone deposits rise from the floor of caves*stalagmites
This is the hardest naturally occurring substance*diamond
What name is used to describe permanently-frozen subsoil?*permafrost
This normally has 4 legs and your butt is parked in it right now*chair
Name the first black nation to gain freedom from European colonial rule*Haiti
Britain and Argentina fought over these islands in 1982*Falklands Islands
Canadian Prime Minister: Pierre Elliott ----------*Trudeau
Churchill, F.D. Roosvelt and Stalin met here in 1945*Yalta
Eras are divided into units called ----------*periods
For what country did Columbus make his historic voyage?*Spain
Four Japanese carriers were destroyed in this battle*Midway
Frankish ruler Charles the Great is better known as ----------*Charlemagne
From what country did the U.S. buy the Virgin Islands?*Denmark
General Sherman burned this city in 1864*Atlanta
Germany's WW I allies were Austria-Hungary, Bulgaria, and ----------*Turkey
Germany's allies in WW II were Japan, Italy, Hungary, Bulgaria, Finland, Libya, and ----------*Romania
He allowed the bugging of the Democratic Committee headquarters*Richard Nixon
He discovered the Grand Canyon*Francisco Coronado
He is identified with the expression, "Eureka"*Archimedes
He is said to have fiddled while Rome burned*Nero
He led 900 followers in a mass suicide in 1979*Jim Jones
He ordered the persecution of Christians in which Peter and Paul died*Nero
He received the Nobel Peace Prize in 1964 for his civil rights leadership*Martin Luther King Jr.
He said, "I have nothing to offer but blood, tears, toil and sweat."*Sir Winston Churchill
He shot Lee Harvey Oswald*Jack Ruby
He taught Alexander the Great*Aristotle
He was assassinated on Dec. 8, 1980 in New York City*John Lennon
He was assassinated on Nov. 22, 1963*John Fitzgerald Kennedy
He was defeated at the Battle of Little Bighorn*General Custer
He was stabbed by Gaius Cassius Longinus*Julius Caesar
He was the American inventor of the Cotton Gin*Eli Whitney
He was the U.S. president during the Civil War*Abraham Lincoln
His ship was the H.M.S. Beagle*Charles Darwin
His wife was Roxana. His horse was Bacephalus. He was ----------*Alexander the Great
How do you write 69 in Roman numerals?*LXIX
How many astronauts manned each Apollo flight?*three
How old was John F. Kennedy when he became president?*forty three
In 1902 this volcano erupted, killing 30,000*Pelee
In what country did "Sepoy Mutiny" occur?*India
In which city was President Kennedy killed?*Dallas
In which city were the Hanging Gardens?*Babylon
In which country did the Boxer Rebellion take place?*China
In which country was Adolf Hitler born?*Austria
In which country was the Rosetta Stone found?*Egypt
Israel occupied the Golan Heights.  Whose territory was it?*Syria
Israel occupied the West Bank.  It belonged to ----------*Jordan
John F. Kennedy Airport in New York used to be called ----------*Idlewind
Mussolini invaded this country in 1935*Ethiopia
Name Jacques Cousteau's research ship*Calypso
Name the incident in which tea was dumped into the harbour*Boston Tea Party
One who fought professionally in Roman arenas was a(n) ----------*Gladiator
She overcame her handicaps to become a lecturer and a scholar*Helen Keller
She was Queen of Egypt and mistress of Julius Caesar*Cleopatra
She was called "The Maid of Orleans"*Joan of Arc
She was the first First Lady to be received privately by the Pope*Jackie Kennedy
She was the first woman premier of Israel*Golda Meir
She was the first woman to fly the Atlantic solo*Amelia Earhart
She was the first woman to swim the English channel*Gertrude Caroline Ederle
She was the greatest trick shot artist of all time*Annie Oakley
She won the 1979 Nobel Peace Prize for her work among the poor*Mother Teresa
Spain ceded Florida to Britain in exchange for this territory*Cuba
The "Bay of Pigs" fiasco took place in this country*Cuba
The Devonian Period is also known as the Age of ----------*fish
The Inquisition forced him to recant his belief in the Copernican Theory*Galileo
The Romans built these to convey water*aqueducts
The St. Valentine's Day massacre took place in this city*Chicago
The first dog in space was named ----------*Laika
The last line of this document is "Working men of all countries, unite."*Communist Manifesto
This F.B.I agent headed the investigation of Al Capone*Elliot Ness
This French peasant girl led the army to victories*Joan of Arc
This Indian group ruled in early Peru*Inca
This Nazi leader had his six children poisoned prior to his own death*Joseph Goebbels
This Queen of France was beheaded in 1793*Marie Antoinette
This Roman killed himself after his defeat at Actium*Marc Antony
This Sioux Indian toured with Buffalo Bill's Wild West Show*Sitting Bull
This Spaniard conquered Mexico*Hernando Cortez
This U.S. Secretary of State won the Nobel Peace Prize in 1973*Henry Kissinger
This U.S. president was fatally shot in 1881*James Garfield
This assassin of Julius Caesar was his friend*Brutus
This frontiersman and politician was killed at the Alamo*Davy Crockett
This is said to be history's greatest military evacuation*Dunkirk
This military attack took place on Dec. 7, 1941*Pearl Harbour
This organization was founded by William Booth*Salvation Army
This racist organization was formed in Tennessee in 1865*Ku Klux Klan
This war began on June 25, 1950*Korean
This was the largest real estate deal in U.S. history*Louisiana Purchase
This word describes the Nazi annihilation of Jews*holocaust
Those big black CD's that you see at garage sales that people call "albums" are made of ----------*vinyl
Two 747's collided here in 1977*Canary Islands
U.S. President, Chester Alan ----------*Arthur
U.S. President, Herbert C. ----------*Hoover
U.S. President, John Quincy ----------*Adams
U.S. President: Calvin ----------*Coolidge
What American feminist went bust as a silver dollar?*Susan B. Anthony
What Chinese dynasty was overthrown in 1911?*Manchu
What age preceded the Iron Age?*The Bronze Age
What did President J. Buchanan not have?*A wife
What did presidents Madison, Monroe, Polk, and Garfield have in common?*The first name "James"
What group landed in America in 1620?*The Pilgrims
What is the Roman numeral for fifty?*L
What is the name of the Russian Czar's daughter who might-or might not-have survived the Russian revolution?*Anastasia
What volcano destroyed Pompeii?*Vesuvius
What was the instrument of execution during the "Reign of Terror"?*guillotine
What was the name of Plato's school?*Academy
What was the name of the B-29 used at Hiroshima to drop the bomb?*Enola Gay
What was the nationality of the prisoners in the "Black hole of Calcutta"?*British
What was the third country to get the "bomb"?*Britain
What's the resting place of those buried at sea?*Davey Jones's Locker
Which U.S. president said, "The buck stops here"?*Truman
Which military battle took place in 1815?*Waterloo
Which nation was led by Genghis Khan?*The Mongols
Which president was responsible for the Louisiana Purchase?*Jefferson
Who is "The Iron Lady"?*Margaret Thatcher
Who led the attack on the Alamo?*Santa Ana
Who rode naked through the streets of Coventry?*Lady Godiva
Who said "Veni, vidi, vici" (I came, I saw, I conquered)?*Julius Caesar
Who said: "Let them eat cake"?*Marie Antoinette
Who sang Happy Birthday to John F. Kennedy for his 45th?*Marilyn Monroe
Who succeeded Churchill when he resigned in 1955?*Sir Anthony Eden
Who was "The Mad Monk"?*Rasputin
Who was George Washington's vice-president?*Adams
Who was the first chancellor of West Germany after WW II?*Konrad Adenauer
Who was the second man to set foot on the moon?*Edwin "Buzz" Aldrin
Who was, "First in war, first in peace and first in the hearts of his countrymen"?*George Washington
What three letters are overly used to indicate "Laugh Out Loud"?*lol
Many Meanings:  Fuel, vapor, flattulence, helium.  What is it?*gas
Multiple Meanings:  Drinking utensils or sight-enhancers*glasses
Multiple Meanings:  Slamming your hands together quickly, or a venereal disease*clap
Name the soda that is often confused with a drug*Coke
Subject, verb and object are parts of a ----------*sentence
This animal is found at the beginning of an (English) encyclopedia*aardvark
What is the only English word formed by the first three letters of the alphabet?*cab
What three letter word means the same as "to ingest"?*eat
"faux pas" means ----------*mistake
Mardi Gras is French for ----------*fat tuesday
The name Australia is derived from the Latin word "australis" which means ----------*southern
The name for this semi-precious stone comes from the Latin for "sea water"*aquamarine
What ONE word fits? ----------hood; ----------hole; ----------date*man
What ONE word fits? ----------stream; ----------hill; ----------pour*down
What does "c'est la vie" mean?*that's life
What word contains the combination of letters: "xop"?*saXOPhone
"The Diary of Anne Frank" was first published in English under what title?*The Diary of a Young Girl
Captain Hook, Tiger Lily, and Tinker Bell are characters in what story?*Peter Pan
Dr. Seuss wrote this book: The Cat in the ----------*Hat
Edgar Allen Poe wrote a famous poem about this animal*raven
From which Shakespeare play is this line taken: "Double, double ... "*Macbeth
From which Shakespeare play is this line taken: "Goodnight, goodnight! Parting is such sweet sorrow, That I should say goodnight till it be morrow."*Romeo and Juliet
From which Shakespeare play is this line taken: "To be or not to be?"*Hamlet
From which Shakespeare play is this line taken: "What in a name? That which we call a rose, by any other name would smell as sweet."*Romeo and Juliet
How many lines are in a sonnet?*fourteen
How many plays is Shakespeare generally credited with today?*thirty seven
Name the Shakespeare play from this ultra short plot summary: Urged on by his wife, a man murders his king in order to take his place*Macbeth
Stephen King's: "Pet ----------"*Semetary
Stephen King's: "Salem's ----------"*Lot
Stephen King's: "The Dead ----------"*Zone
This Shakespearean king was the actual king of Scotland for 17 years*Macbeth
This girl hid from the Nazis in Amsterdam*Anne Frank
What Shakespearean play features the line: A plague on both your houses?*Romeo and Juliet
What did Jeannie C. Riley describe as "a little Peyton Place"?*Harper Valley
What other name does Stephen King write under?*Richard Bachman
What publication was subtitled The What's New Magazine?*Popular Science
What was the name of Mother Goose's son?*Jack
What's Penthouse's sister publication for women?*Viva
Homer wrote this account of the Trojan War*Iliad
You are in a room where all walls face south.  A bear walks by.  What color is it?*White
A U.S. dime is worth ---------- cents*ten
A line drawn from an angle of a triangle to the mid-point of the opposite side is a(n) ----------*median
A triangle with three equal sides is called ----------*equilateral
A triangle with two equal sides is called ----------*isosceles
An angle greater than 180 degrees and less than 360 degrees is a(n) ---------- angle*reflex
An angle greater than 90 degrees and less than 180 degrees is said to be ----------*obtuse
An integer that is greater than 1 and is divisible only by itself and 1 is known as a(n) ----------*prime
Approximately how many inches are there in one meter?*thirty nine
Arc, radius, and sector are parts of a(n) ----------*circle
He is known as "The Father of Geometry"*Euclid
How many corners are there in a cube?*eight
How many nickles are there in 2.25?*forty five
If you cut through a solid sphere what shape will the flat area be?*circle
Name the number system which uses only the symbols 1 and 0*The binary system
The angles inside a square total ---------- degrees*three hundred and sixty
The mathematical study of properties of lines, angels, etc., is ----------*geometry
The space occupied by a body is called its ----------*volume
The square root of 1 is?*one
Two angles that total 180 degrees are called ----------*supplementary
What geometric shape has 4 equal sides?*square
A baby doctor is a ----------*pediatrician
A bone specialist is a(n) ----------*osteopath
A loss of memory is known as ----------*amnesia
A medicine that hastens the emptying of the bowels is called a ----------*laxative
A non-cancerous tumor is said to be ----------*benign
A thread used in surgery to tie a bleeding blood vessel is called a(n) ----------*ligature
Carditis, affects the ----------*heart
Doctors often have this instrument around their neck*stethoscope
Due to a lack of vitamin C, sailors used to contract this disease*scurvy
Encephalitis affects the ----------*brain
Gastritis affects the ----------*stomach
Hammer, anvil, and stirrup are parts of the ----------*ear
He discovered the process of vaccination for prevention of smallpox*Edward Jenner
Hepatitis affects the ----------*Liver
How many bones are there in the human body?*two hundred and six
How many chambers does the human heart have?*four
How many pints of blood does the average human have in his/her body?*twelve
In the field of psychiatry this term means self-love*narcissism
In what organ of the body is insulin produced?*pancreas
In which organ is a clear watery solution known as the "aqueous humor" found?*eye
In which organ is a pulmonary disease located?*lung
In which organ is your "hypothalmus" located?*brain
Infantile Paralysis is commonly known as ----------*polio
Meningitis affects the ----------*brain
Myositis affects the ----------*muscles
Name the hardest substance in the human body*enamel
Name the largest artery in the human body*aorta
Name the largest gland in the human body*liver
Osteomyelitis affects the ----------*bones
Peritonitis, affects the ----------*abdomen
Prosthetics deals with the making of ----------*artificial limbs
Tarsus, metatarsus, and phalanges are parts of a(n) ----------*foot
The branch of medicine dealing with curing by operative procedures is ----------*surgery
The lack of this element in the diet is a cause of goitre*iodine
The largest single organ of the human body is the ----------*skin
The medical name for the voice box is the ----------*larynx
The teeth used for biting or cutting are known as ----------*incisors
These animals were once used to bleed the sick*leeches
These attach muscles to bones or cartilage*tendons
This branch of medicine deals with old age and its diseases*geriatrics
This disease consists of a purposeless, continual growth of white blood cells*leukemia
This fingerlike projection is attached to the large intestine*appendix
This is known as "The Royal Disease"*haemophilia
This large bean-shaped lymph gland can expand and contract as needed*spleen
This membrane controls the amount of light entering the eye*iris
This organ is a small pouch that stores bile*gall bladder
This parasite lives in the intestines of man and animals*tapeworm
This small gland attached to the brain exerts a control over growth*pituitary
Ulna, radius, and clavicle are types of ----------*bone
What disease is also known as "rubella"?*German measles
What hormone is produced by the adrenal glands?*adrenaline
What is a skin specialist called?*dermatologist
What is an organism called that lives on or in a host animal?*parasite
What is the biggest disqualifying factor for prospective astronauts?*eyesight
What is the medical term for cancer of the blood?*leukemia
What is the name of the bone in the lower leg?*tibia
What kind of poisoning is known as plumbism?*lead poisoning
What toe is the foot reflexology pressure point for the head?*big toe
Where do you find the medulla oblongata?*brain
Where in the body is the tiniest human muscle?*ear
Which disease is also known as "Hansen's Disease"?*leprosy
With what body part is otology involved?*ear
If the groundhog sees his shadow on Feb. 2, there will be how many more weeks of bad weather?*six
This dry, warm wind flows eastward down the slopes of the Rocky Mountains*Chinook
Actor: ---------- Borgnine*Ernest
Actor: ---------- Hackman*Gene
Actor: ---------- Nimoy*Leonard
Actor: ---------- Savalas*Telly
At the end of "Planet of the Apes" what protruded from the rocks?*Statue of Liberty
Barbara Streisand was the female lead in "Hello, Dolly".  Who was the male lead?*Walter Matthau
Charles Laughton played Quasimodo in this epic film*Hunchback of Notre Dame
Darth Vader was the villan in the movie, "---------- Wars"*Star
Film Title: An Officer and a(n) ----------*Gentleman
Film Title: Fahrenheit ----------. (a number)*451
Film Title: The Last Days of ----------. (a city)*Pompeii
Film Title: ---------- (a number) Leagues Under the Sea*20,000
Forrest ---------- liked shrimp*Gump
He directed "The Godfather"*Francis Ford Coppola
He directed the movie E.T*Stephen Spielberg
He played Superman in the 1978 movie version*Christopher Reeve
He starred in "Conan the Barbarian"*Arnold Schwarzenegger
He starred in, "City Lights"*Charlie Chaplin
He was known as the "Elephant Man"*Joseph Merrick
He was the villain in "Star Wars"*Darth Vader
His films include: Giant, Written on the Wind, and A Farewell to Arms*Rock Hudson
His films include: Spartacus, The Vikings, and Ulysses*Kirk Douglas
In "Gone With the Wind", Scarlett regains her wealth by investing in what type of business?*sawmill
In 1975 Jack Nicholson won the best actor Oscar for his role in this film*One Flew Over the Cuckoo's Nest
In the 1996 version of "Romeo and Juliet", who played Juliet?*Claire Danes
In the film "Bright Eyes", Shirley Temple sang about this boat*The Good Ship Lollipop
In this 1968 film the husband of an unsuspecting young wife becomes involved with a witch's coven*Rosemary's Baby
In which Disney movie is the song "So This Is Love"?*Cinderella
Name the Disney cartoon in which the character "Belle" appears*Beauty and the Beast
Name the actor who played the leading role in "The Good, the Bad, and the Ugly"*Clint Eastwood
Name the first full length color cartoon talking picture*Snow White
Name the musical film named after a state*Oklahoma
Nick Nolte and Eddie Murphy star in this 1982 film*48 Hours
Peter Sellers is best known for his role as Inspector ----------*Clouseau
She played Lois Lane in the 1978 film version of "Superman"*Margot Kidder
She played a Polish refugee in "Sophie's Choice"*Meryl Streep
She played the lead role in "Coal Miner's Daughter"*Sissy Spacek
She starred in the 1952 film, "Niagara"*Marilyn Monroe
The first 18 minutes of this movie is black and white*Wizard of Oz
The song "Matchmaker, Matchmaker" came from which musical play?*Fiddler On The Roof
The two rival gangs in "West Side Story" were the Sharks and the ----------*Jets
This 1974 film started a run of nostalgia culminating in the TV series "Happy Days"*American Graffiti
This Disney movie relies heavily on computer animation*TRON
This actress appeared in "St. Elmo's Fire", "The Scarlett Letter", and "Striptease"*Demi Moore
This actress was Miss Hungary of 1936*Zsa-Zsa Gabor
This film starring Julie Andrews and Christopher Plummer wont he best picture Oscar for 1965*Sound of Music
This film starring Richard Beymer and Natalie Wood won the best picture Oscar for 1961*West Side Story
This film was an ambitious concert sequence of cartoons by Walt Disney*Fantasia
This is a classic film about a huge gorilla*King Kong
This magic word was in the movie, "Mary Poppins"*Supercalifragilisticexpialidocious
This movie directed by Woody Allen won the best picture Oscar in 1978*Annie Hall
This movie is about the migration of poor workers from the dust bowl to the Californian fruit valleys*The Grapes of Wrath
This movie starring Marlon Brando won the best picture award in 1972*The Godfather
This was the first 3-D film*Bwana Devil
This was the first cartoon talking picture*Steamboat Willie
This was the sequel to "Star Wars"*The Empire Strikes Back
This was the sequel to "The Empire Strikes Back"*Return of the Jedi
What actor appeared in all three of these films, Straw Dogs, Midnight Cowboy, and The Graduate?*Dustin Hoffman
What actress has received the most Oscar nominations?*Katherine Hepburn
What actress's real name was Frances Gumm?*Judy Garland
What city's police force did Charlie Chan work with?*Honolulu
What color was Bullitt's car?*Green
What country was the setting for "The King and I"?*Siam (Thailand)
What detective duo was featured in Mystery at Devil's Paw?*Hardy
What does the statue of Oscar stand on?*A Reel of Film
What famous animal character called "Skull Island" home?*King Kong
What film did John Wayne win his only Oscar for?*True Grit
What is Hawkeye's full name in M.A.S.H.?*Benjamin Franklin Pierce
What is the destination of the plane at the end of the film "Casablanca"?*Lisbon
What is the name of Pearce Brosnan's first James Bond film?*Goldeneye
What is the name of the Volkswagen in the film, "The Love Bug"?*Herbie
What is the name of the rabbit in the film, "Bambi"?*Thumper
What is the name of the skunk in the film, "Bambi"?*Flower
What is the name of the whale that swallowed Pinocchio*Monstro
What is the stage name of Greta Gustafson?*Greta Garbo
What is the title of the 1996 sequel to "Terms of Endearment"?*Morning Star
What kind of creature was Chewbacca in "Star Wars"?*Wookiee
What two words are normally at the end of most movies?*The End
What was "Rocky's" last name?*Balboa
What was Citizen Kane's dying word?*Rosebud
What was Dorothy's last name in "The Wizard of Oz"?*Gale
What was Rocky's nickname in the ring?*The Italian Stallion
What was Sir Alec Guinness's role in "Star Wars"?*Obi-Wan Kenobi
What was the first film directed by Robert Redford?*Ordinary People
What was the name of Ashley Wilkes' plantation in "Gone With the Wind"?*Twelve Oaks
What was the name of Han Solo's spaceship in "Star Wars"?*Millennium Falcon
What was the name of Luke's strange little advisor in "The Empire Strikes Back"?*Yoda
What was the name of the motel in the film "Psycho"?*Bates Motel
What was the setting for "The Sound of Music"?*Austria
Which actor said, "Love means never having to say you're sorry"?*Ryan O'Neil
Which character in "Forrest Gump" loved shrimp?*Bubba
Which character sang "Come Out, Come Out, Wherever You Are" in "The Wizard of Oz"?*Glinda
Which character sang, "When you wish upon a star.." in Disney's "Pinocchio"?*Jiminy Cricket
Which comedy duo did the famous, "Who's on first?" routine?*Abbott and Costello
Which planet was the "Planet of the Apes"?*Earth
Who co-starred with Julie Andrews in "Mary Poppins"?*Dick Van Dyke
Who hosted the 1997 Grammy Awards?*Ellen Degeneres
Who is Warren Beatty's sister?*Shirley MacLaine
Who is the fastest mouse in all of Mexico?*Speedy Gonzalez
Who is the male lead in the "Naked Gun" movies?*Leslie Nielsen
Who is the voice of Darth Vadar?*James Earl Jones
Who played "Robin" to Val Kilmer's "Batman"?*Christopher O'Donnell
Who played Brad Pitt's cop partner in the movie "Seven"?*Morgan Freeman
Who played Dorothy in "The Wizard of Oz"?*Judy Garland
Who played Garth in "Wayne's World"?*Dana Carvey
Who played Matt Helm in the movies?*Dean Martin
Who played Scarlette O'Hara in "Gone With the Wind"?*Vivien Leigh
Who played the 'Wicked Witch of the West' in "The Wizard of Oz"*Margaret Hamilton
Who played the role of Richard Blaine in Casablanca?*Humphrey Bogart
Who plays the character of the only escapee from Alcatraz in the movie "The Rock"?*Sean Connery
Who portrayed Han Solo in "Star Wars"?*Harrison Ford
Who portrayed Moses in "The Ten Commandments"?*Charlton Heston
Who produced, directed and starred in "Citizen Kane"?*Orson Welles
Who starred with John Travolta in the movie "Broken Arrow"?*Christian Slater
Who was C3PO's sidekick in "Star Wars"?*R2D2
"He's So Fine", "One Fine Day" and "A Love So Fine" where hits for what fine group?*The Chiffons
"Joy to the World" was a hit in 1971 for what band with three lead vocalists?*Three Dog Night
Besides "Auld Lang Syne" and "For He's a Jolly Good Fellow", what is the most frequently sung song in English?*Happy Birthday
Fifties rock "n" roll was revived by what greased hair, T-shirted, TV frequenting group?*Sha Na Na
Jerry Lee Lewis had Great ---------- Of Fire*balls!
Name Jerry Garcia's long lived group*The Grateful Dead
The Who's rock musical stars Elton John.  It's called ----------*Tommy
This female artist enjoyed sucess on both popular and country & western stations with such tunes as "Let Me Be There" and "Have You Never Been Mellow."*Olivia Newton-John
This was the Beatle's first film*A Hard Day's Night
Through 1963 this duo's total record sales exceeded 18 million with successes including "Cathy's Clown" and "Wake Up Little Suzie"*The Everly Brothers
What Procol Harem tune was based on the Bach cantata "Sleepers Awake"?*A Whiter Shade of Pale
What band recorded the 1978 hit album: "Briefcase Full of Blues"?*The Blues Brothers
What city is also known as Music City, U.S.A.?*Nashville
What country singer/songwriter (and sometimes actor) is known as "the country outlaw"?*Willie Nelson
What entertainer is allowing one of his songs to be used in a government campaign to beat drunk driving?*Michael Jackson
What famous classical composer continued to compose great music after becoming deaf?*Ludwig van Beethoven
What famous singer was known to give automobiles to complete strangers?*Elvis Presley
What group refused to have their pictures taken while they were not in their makeup?*Kiss
What group's biggest-ever hit was Be My Baby?*The Ronettes
What song by Don McLean talks about the day Buddy Holly died?*American Pie
What song did Aretha Franklin sing in "The Blues Brothers"?*Think
What was Elvis Presley's wife's name?*Priscilla
What was The Beatles' biggest hit single?*Hey Jude
What's the name of B.B. King's guitar?*Lucille
Which of Paul Simon's musical characters was told to hop on the bus?*Gus
Who invented the electrical bass?*Leo Fender
Who recorded the 1957 hit "Tammy"?*Debby Reynolds
Who recorded the lengthy song: "In-A-Gadda-Da-Vida" in 1969?*Iron Butterfly
Who was the first singer in Genesis?*Peter Gabriel
Whose theme song was Back In The Saddle Again?*Gene Autry's
The sea gods had a three-pronged spear called a(n) ----------*trident
This Greek mountain was known as the home of the gods*Olympus
What has no reflection, no shadow, and can't stand the smell of garlic?*vampire
What's heaven to fallen Norse warriors?*Valhalla
Where did Robin Hood supposedly live?*Sherwood Forest
Which ancient continent is said to be submerged?*Atlantis
Name the largest of the dinosaurs*brachiosaurus
The remains of prehistoric organisms that have been preserved in rocks are called ----------*fossils
A "sirocco" refers to a type of ----------*wind
A calm ocean region near the equator*doldrums
A coral island consisting of a ring of rock enclosing a central lagoon is a(n) ----------*atoll
A great wave resulting from an earthquakes is called a (an) ----------*tsunami
A group of gorillas is known as a ----------*band
A group of kangaroos is known as a ----------*troop
A hot spring which shoots steam into the air is a ----------*geyser
A one-humped camel is called a ----------*dromedary
A relationship between two different types of organism which live together for their mutual benefit*symbiosis
A terrapin is a type of ----------*turtle
An animal is a bird if it has ----------*feathers
An animal is a fish if it has ----------*gills
Corolla, filament and stigma are parts of a(n) ----------*flower
Dense sea-water swamps along coasts of hot countries are called ----------*mangroves
Dogs bark.  What do donkeys do?*bray
Excluding man, what is the longest-lived land mammal?*elephant
Fandible, lateral line, and dorsal fin are parts of a(n) ----------*fish
From what animal is "ambergis" obtained?*sperm whale
How man legs does a crab have?*ten
How many teats does a cow have?*four
How many tentacles does a squid have?*ten
Imperial, Buck, and Luna are types of ----------*moth
Knife, Clown, and Pencil are types of ----------*fish
Linseed oil is obtained from the seed of which plant?*flax
Maxillary palps, abdomen, and metathorax are parts of a(n) ----------*insect
Name the fastest land animal over a prolonged distance (1,000 yd. plus)*antelope
Name the heaviest breed of domestic dog*St. Bernard
Name the heaviest flying bird of prey*condor
Name the largest living bird*ostrich
Name the largest web-footed bird*albatross
Name the longest venomous snake*cobra
Name the mammal living at the highest altitude*yak
Name the most venomous spider*black widow
Name the only native North American marsupial*opossum
Name the slowest moving land mammal*sloth
Name the smallest breed of dog*chihuahua
Name the wild dogs of Australia*dingo
Paper is made from the pulp of ----------*wood
Snakes are reptiles.  What are frogs?*amphibians
Some animals spend the winter in a sleep-like state known as ----------*hibernation
The "canebrake", "timber" and "pygmy" are types of what?*rattlesnake
The fins of which fish are made into a soup?*shark
The four stages in the life-cycle of an insect are: egg, adult, pupa, and ----------*larva
The koala bear eats the leaves from this tree*eucalyptus
These marine crustaceans often attach themselves to the hulls of ships*barnacles
This African animal kills the most people*crocodile
This animal is armed with bony plates and rolls up into a ball if frightened*armadillo
This animal is kept as a house pet to kill cobras*mongoose
This animal is normally measured in "hands"*horse
This animal is the symbol of the U.S. Democratic Party*donkey
This animal is the symbol of the U.S. Republican Party*elephant
This animal's name is the same as that given to a high church official*cardinal
This animal's shell is used to make attractive jewelry*abalone
This bird lays its eggs in the nests of other birds*cuckoo
This is the largest of the deer family*moose
This is the main food of the blue whale*plankton
This is the only mammal with four knees*elephant
This order of insects contains the most species*beetle
This organic gem is a deep red secretion from a marine animal*coral
This small animal is trained to hunt rats and rabbits*ferret
This two ton animal can gallop at over 50 miles an hour*rhinoceros
This ugly creature has patches of red on his rear-end*mandrill
Walrus tusks are made of ----------*ivory
What are the pouched animals called?*marsupials
What bird is an excellent swimmer, but can't fly?*penguin
What does a camel store in its hump?*fat
What fish is the fastest?*sailfish
What is a group of larks called?*exaltation
What is a group of peacocks called?*muster
What is a group of whales called?*pod
What is a male goose called?*gander
What is a male swan called?*cob
What is a male swine called?*boar
What is a young goose called?*gosling
What is a young swan called?*cygnet
What is an emasculated stallion called?*gelding
What is another term for a black leopard?*panther
What is the horn of a rhinoceros made of?*hair
What is the only dog that doesn't have a pink tongue?*chow
What is the world's longest snake?*python
What large herbivore sleeps only one hour a night?*antelope
What name is given to a female calf?*heifer
What plant is opium derived from?*poppy
What travels in gaggles?*geese
What type of animal is a wallaby?*marsupial
What type of animal lives in a formicary?*ant
What word is used for a female fox?*vixen
What word is used for a female sheep?*ewe
What word is used for a male deer?*buck
What word is used for a male duck?*drake
What do oak trees grow from?*acorns
What would be kept in an "aviary"?*birds
Where is the Ocean of Storms located?*On the moon
Who is Robert Zimmermann?*Bob Dylan
What are white dwarfs and red giants?*stars
What are these: chrysolite, beryl, jasper, and tourmaline?*gems
Jack and Jill went up a ---------- to fetch a pail of water?*hill
Name the nursery rhyme mother whos last name is that of a bird?*goose
The energy which a body possesses by virtue of its motion is called ----------*kinetic
The pivot point of a lever is called the ----------*fulcrum
The process of water changing to water vapor is known as ----------*evaporation
The rate of change of velocity is known as ----------*acceleration
The visible spectrum of light ranges from red to ----------*violet
True Or False: Contrary to popular belief, a lightbulb actually absorbs darkness?*false
Two 1.5 volt batteries, when connected in series, produces ---------- volts*three
Water freezes at ---------- degrees Celcius*zero
Water freezes at ---------- degrees Fahrenheit*thirty two
What is measured by a Geiger counter?*radioactivity
When light waves pass from one medium into another they change direction.  This is called ----------*refraction
Work done, equals force multiplied by ----------*distance
For what does O.P.E.C. stand?*The Organization of Petroleum Exporting Countries
He was elected President of France, in 1981*Francios Mitterrand
Name Ronald Reagan's first wife*Wyman
U.S. President, Woodrow ----------*Wilson
How is the state of Mississippi spelled?*Mississippi
Name the circular plastic media you put in a CD-ROM drive*CD
On a standard computer keyboard, this is the letter just to the right of Z*x
The longest key on your keyboard is the ---------- bar*space
The season ---------- comes right after Spring*summer
There are ---------- seconds in a minute*sixty
This is required to make all electric things work?*electricity
What color is a blue crayon?*blue
What color is a green crayon?*green
What color is a red crayon?*red
What company makes Microsoft Windows 2000?*microsoft
Whats radar spelled backwards?*radar
Whats the abbreviation for United States of America?*USA
When spelled backwards, the word "retupmoc" becomes what?*computer
Followers of the Unification Church are called ----------*Moonies
He led the Israelites out of Egypt*Moses
He was the first King of the Hebrews*Saul
He was the second King of Israel*David
In what month is Christmas observed?*December
Name the holiest day in the Jewish calendar*Yom Kippur
On which mountain did Moses receive the Ten Commandments?*Sinai
Which city is sacred to Jews, Christians, and Muslims?*Jerusalem
Who founded Mormonism?*Joseph Smith
Who founded the People's Temple Commune?*Jim Jones
A flat-bottomed conical laboratory flask with a narrow neck is called a(n) ----------*erlenmeyer flask
A phrenologist reads ----------*skulls
A point to which rays of light converge is called a(n) ----------*focus
A shallow dish with a cover, used for science specimens is a(n) ----------*petri dish
Acetylsalicylic acid is more commonly known as ----------*aspirin
Botany and Zoology combined make up the science of ----------*biology
By what chemical process do plants manufacture food?*photosynthesis
By what name is Lysergic acid diethylamide better known?*LSD
Cetology is the study of ----------*whales
Circuits can be wired in series or in ----------*parallel
Cocci, Spirilla, and Streptococci are types of ----------*bacteria
Deoxyribonucleic acid is better known as ----------*DNA
Dermatitis affects the ----------*skin
Epidermal cells, palisade cells, and veins are parts of a(n) ----------*leaf
Ethylene glycol is frequently used in automobiles.. How?*anti-freeze
Forked, Sheet, and Ball are types of ----------*lightning
Growing plants in liquids rather than soil is known as ----------*hydroponics
He designed the first feasible automobile with an internal combustion engine*Karl Freidrich Benz
He invented "bifocal" lenses for eyeglasses*Benjamin Franklin
He is known for his theory of "Evolution"*Charles Darwin
He transmitted radio signals across the Atlantic in 1901*Enrico Marconi
He wrote "Sexual Behaviour in the Human Male" in 1948*Alfred Kinsey
How many degrees does the earth rotate each hour?*fifteen
In which country was the match invented?*France
Nitrous oxide is better known as ----------*laughing gas
Nuclear membrane, cytoplasm, and nucleus are parts of a(n) ----------*cell
Pulp, crown, and root are parts of a(n) ----------*tooth
Quinine is added to water to make ----------*tonic water
Sound Navigation Ranging is better known as ----------*sonar
The Kelvin scale is used to measure ----------*temperature
The filament of a regular light bulb is usually made of ----------*tungsten
The instrument used in geometry to measure angles is a(n) ----------*protractor
The name for the Russian equivalent of Skylab is ----------*Salyut
The second space shuttle was named ----------*Columbia
The vernal equinox is the beginning of ----------*spring
This Russian scientist used dogs to study conditioned reflexes*Ivan Pavlov
This is like an airplane but has its propeller on top instead*helicopter
This is the reading system used by the blind*Braille
This science deals with the motion of projectiles*ballistics
To make a car go backwards you have to put it in what gear?*reverse
What are these: Ceres, Juno, Iris, and Flora?*asteroids
What branch of science studies the motion of air and the forces acting on objects in air?*aerodynamics
What did Lewis E. Waterman invent in 1884?*fountain pen
What does the "lithosphere" refer to?*The earth's crust
What does the Binet test measure?*intelligence
Who invented dynamite?*Alfred Nobel
Who is known as the father of genetics?*Gregor Mendel
As bold as ----------*brass
As clean as a(n) ----------*whistle
As clear as a(n) ----------*bell
As cute as a(n) ----------*button
As easy as ----------*pie
As fit as a(n) ----------*fiddle
As good as ----------*gold
As graceful as a(n) ----------*swan
As large as ----------*life
As loud as ----------*thunder
As pale as a(n) ----------*ghost
As pleased as ----------*punch
As pretty as a(n) ----------*picture
As proud as a(n) ----------*peacock
As quiet as a(n) ----------*mouse
As sick as a(n) ----------*dog
As sly as a(n) ----------*fox
As smart as a(n) ----------*whip
Fresh as a(n) ----------*daisy
Baseball: The Baltimore ----------*Orioles
Baseball: The Toronto ----------*Bluejays
Basketball: The Boston ----------*Celtics
Basketball: The Denver ----------*Nuggets
Basketball: The Los Angeles ----------*Lakers
Basketball: The New York ----------*Knicks
Basketball: The Seattle ----------*Supersonics
Basketball: The Utah ----------*Jazz
Football: The Baltimore ----------*Colts
Football: The Buffalo ----------*Bills
Football: The Cincinnati ----------*Bengals
Football: The Cleveland ----------*Browns
Football: The Dallas ----------*Cowboys
Football: The Denver ----------*Broncos
Football: The Miami ----------*Dolphins
Football: The Minnesota ----------*Vikings
Football: The New Orleans ----------*Saints
Football: The Seattle ----------*Seahawks
He holds the NHL record for the most goals scored during the regular season*Wayne Gretzky
He was the NBA, MVP in 1976, 77, and 80*Kareem Abdul-Jabbar
Hockey: The Boston ----------*Bruins
Hockey: The Buffalo ----------*Sabres
Hockey: The Calgary ----------*Flames
Hockey: The Chicago ----------*Blackhawks
Hockey: The Detroit ----------*Red Wings
Hockey: The Edmonton ----------*Oilers
Hockey: The Los Angeles ----------*Kings
Hockey: The Montreal ----------*Canadians
Hockey: The Pittsburgh ----------*Penguins
Hockey: The St. Louis ----------*Blues
Hockey: The Toronto ----------*Maple Leafs
Hockey: The Vancouver ----------*Canucks
How many feet high is a basketball net?*ten
How many games must you win to win a normal set in tennis?*six
How many minutes is a major penalty in hockey?*five
How many minutes is each period of hockey?*twenty
How many players are there on a soccer team?*eleven
How many players are there on a water polo team?*seven
How many players make up a field hockey team?*eleven
How many points are awarded for a safety touch in football?*two
How many referees work a soccer game?*one
How many seams are there on a football? (American)*four
How many sides does a home-plate have?*five
In pro football a "sudden death" period lasts how many minutes long?*fifteen
In ten-pin bowling, how many points does a perfect game consist of?*three hundred
In this team sport each player gets a chance to play every position*volleyball
In what sport is the Heisman trophy awarded?*American football
In which city is the Cotton Bowl played?*Dallas
In which city is the Hockey Hall of Fame located?*Toronto
In which game or sport is a "Zamboni" used?*hockey
In which sport is a "hole-in-one" possible?*golf
In which sport is the America's Cup awarded?*sailing
In which sport is the Cy Young Trophy awarded?*baseball
In which sport is the Davis Cup awarded?*tennis
In which sport is the term "love" used?*tennis
In which sport is the term "wishbone" used?*football
In which sport is the term, "Hang ten" used?*surfing
In which sport would you find the "slapshot"*hockey
Name the hockey trophy awarded to the player demonstrating the best sportsmanship*The Lady Byng Trophy
On what type of surface are the tennis matches at Wimbledon played?*grass
She was "Sports Illustrated's" first female "Sportsman of the Year"*Billie Jean King
Soccer: The New York ----------*Cosmos
The 1976 Summer Olympics were held in this city*Montreal
The Japanese martial art of fencing is called ----------*kendo
The person who carries the golfer's clubs is called a(n) ----------*caddie
The white marks intersecting each five yard line are called ----------*hash marks
This is the most coveted trophy in Candian football*Grey Cup
This sport allows substitutions without a stoppage in play*hockey
This sport is called the "American pastime"*baseball
This team won their first World Series in 1969*New York Mets
This traditional Japanese wrestling sport takes place in a circular ring*sumo
Two under par on a hole of golf is called a(n) ----------*eagle
What are large snow bumps known as in skiing terms?*moguls
What do runners pass to each other in a relay race?*baton
What do the letters ERA mean in baseball?*Earned Run Average
What does TKO stand for?*Technical Knock Out
What football player rushed for 2,003 yards in 1973?*OJ Simpson
What game features the largest ball?*earthball
What is it called when a football team loses possession of the ball due to a misplay?*turnover
What is the heaviest class of weight-lifting?*super heavyweight
What number wood is a driver in golf?*one
What sport do the Harlem Globetrotters play?*basketball
What sport has a hooker in a scrum?*rugby
What sport would you helicopter to the Bugaboos for?*skiing
What trophy is awarded to the winner of the NHL play-offs?*Stanley Cup
What vehicles are involved in the "Tour de France"?*bicycles
What was football player Dick Lane's nickname?*Night Train
What's the nickname of the University of Georgia football team?*Bulldogs
Where are the U.S. Tennis Open Championships held?*Flushing Meadows, NY
Which NFL team's defensive unit was nicknamed "The Purple People Eaters"?*Minnesota Vikings
Which country won the World Cup of Soccer in 1982*Italy
Which football team was nicknamed the "Orange Crush"?*Denver Broncos
Which is the only position in soccer allowed to handle the ball?*goalkeeper
Which position is usually played by the tallest member on a basketball team?*centre
Which sport has a movement called a "telemark"?*skiing
Which sport uses stones and brooms?*curling
Which swimming stroke is named after an insect?*butterfly
Who was known as the "Sultan of Swat"?*Babe Ruth
Who was the 1978 Wimbledon Women's Singles champ?*Martina Navratilova
Who was the first NHL player to score 50 goals in one season?*Maurice Richard <-- pronounced "Reeee-shard"
Who was the first to win the grand slam of tennis?*Don Budge
With which sport is Chris Evert Lloyd identified?*tennis
When read upside down, what does the term "umop apisdn" signify?*upside down
If you look at the sun long enough, you go ----------*blind
This is the sandy area nearest the ocean*beach
What did TVs IMF stand for?*Impossible Mission Forces
What was the first network series devoted entirely to rock and roll?*American Bandstand
Who was Carl in Five Easy Pieces before going to Walton's Mountain?*Waite
Who was Chief Marshall of the Mickey Mouse Club?*Walt Disney
Hook, line and ----------*sinker
"7X" was used to refer to the secret ingredient of which drink?*Coca Cola
A "pigskin" is another name for a(n) ----------*football
A can of Pepsi holds ---------- fluid ounces*twelve
A foot-long ruler is ---------- inches long*twelve
Most men do this each morning, using a razor*shave
Most people wear a watch on their ---------- wrist*left
Name the implement that removes water from your windshield on your car?*wiper
Traffic Trivia:  Red means stop, ---------- means go*green
What US state has the Worlds Champion Chili Cookoff every year?*Texas
What bit of Bobby Goldsboro syrup focused on a dying young wife?*honey
What did Sally Rogers always wear in her hair?*a bow
What does the typical American eat 263 of each year? not pizza!!*eggs
What game show had people dressed up funny?*Let's Make A Deal
What is Grimace of the McDonald's characters?*A tastebud
What was Simple Simon fishing for in his mother's pail?*whale
What was the name of the Akla-Seltzer boy?*Speedy
What was the name of the dog in RCA Victor's trademark?*Nipper
What's the telephone area code for Chicago?*312
Where did Howdy Doody live?*Doodyville
Where were Archie and Edith Bunker's chairs enshrined?*The Smithsonian Institute
Who is also known as Mr. Warmth?*Don Rickles
Who is known as "the world's oldest teenager"?*Dick Clark
Ball and ----------*chain
Peace and ----------*quiet
Which word is related to these three: nap, walk, call?*cat
Which word is related to these three: painting, bowl, nail?*finger
Which word is related to these three: rat, blue, cottage?*cheese
A castle or enclosed place*fort
Make holes through something*perforate
Many trees*forest
Pardon*forgive
A choice cut of meat*tenderloin
What do the initials 'VCP' stand for?*Video Cassette Player
What do the initials 'VCR' stand for?*Video Cassette Recorder
What does 'A&W' (of root beer fame) stand for ?*Allen & Wright
What does 'AOL' stand for?*America Online
What does S.O.S. stand for?*Save Our Souls
What does N.A.S.A stand for?*National Aeronautics and Space Administration
What does the acronym 'scuba' mean?*Self Contained Underwater Breathing Apparatus
Who is the dog on the crackerjack box?*Bingo
What is ground being 'rested' for a season called?*fallow
After what were the B52 bombers named?*a fifties hairdo
How many engines are on a B52 bomber?*eight
How many gallons of fuel does a jumbo jet use during take off?*four thousand
What does a pilot drop to slow an airplane?*flaps
What is the world's fastest passenger aircraft?*Concorde
What type of craft is the US's Airforce One?*Boeing 747
Whic country developed the first jet fighter?*Germany
Which two nations built the concorde?*Britain and France
Who built the 'Cherokee' and 'Comanche' aircraft?*Piper
Who built the hurricane aircraft?*Hawker
From what is the liqueur kirsch made?*cherries
From which plant is tequila derived?*cactus
As what is California also known?*Golden State
As what is Minnesota also known?*Gopher State
What city is also known as Beantown?*Boston
What state is 'The Golden State'?*California
What state is also called the 'Garden State'?*New Jersey
What state is the 'Hoosier State'?*Indiana
Where are the headquarters of the CIA?*Langley, Virginia
Which date is inscribed on the book held by the Statue Of Liberty?*July 4 1776
What is the most air polluted city in the United States?*Los Angeles
Where is the Kitty Hawk?*The Smithsonian Institute
What mountain has the figures of three mounted confederate heroes of the Civil War?*Stone Mountain
What state is only part of the United States by treaty?*Texas
Which two fruits are an anagram of each other?*lemon and melon
How many litres of air is in an adult lung?*five
How many times do your ribs move every year during breathing?*five million
Like fingerprints, what other print is individual?*tongueprints
Of what does the typical man have 13,000?*whiskers
What do the auricularis muscles move?*ears
What is the Scientific name for the eardrum?*tympanic membrane
What is the common name for the scapula?*shoulder blade
What is the common name for the sternum?*breastbone
What is the common name for the tympanic membrane?*eardrum
What is the second largest bone in the foot?*talus
What is the smallest bone in the human body?*stirrup bone
Where are one quarter of the bones in the human body?*feet
Which is the most sensitive finger?*forefinger
As what is a camelopard also known?*giraffe
As what is a giraffe also known?*camelopard
As what is a moose also known?*algonquin
As what is an algonquin more commonaly known?*moose
At what age does a filly become a mare?*five
Does a wild rabbit live 10, 15 or 20 years?*ten
How fast (mph) can a kangaroo hop?*forty
How many hours a day does a ferret sleep?*twenty
How many hours does an antelope sleep at night?*one
How many teeth does a walrus have?*eighteen
If a robin's egg is put in vinegar for thirty days, what colour does it become?*yellow
Name one male fish that gives birth?*sea horse or pipe fish
Of what are walrus tusks made?*ivory
If body temperature was 86 degrees, how many years would a man man live?*two hundred
In 1986, what was the maximum fuel capacity (in litres) imposed in Formula 1 racing?*195
On what do honeybees have a type of hair?*eyes
Some animals always grow new teeth to replace the old. Name one of them!*crocodile
Some animals always grow new teeth to replace the old. Name one of them!*shark
What animal can get the disease 'heaves'?*horse
What animal can hop as fast as 40 mph?*kangaroo
What animal has red patches on its rear?*mandrill
What animal lives in a form?*hare
What animal lives in a warren?*rabbit
What animal's milk is more than 54% fat?*humpback whale
What are the only other animals on which the pill works?*gorillas
What bird is associated with Lundy Island?*puffin
What colour is a robin's egg?*blue
What comprises than 54% of humpback whale's milk?*fat
What dog is named after a Mexican state?*chihuahua
What herbivore sleeps one hour a night?*antelope
What insect has a type of hair on it's eyes?*honeybees
What is a female deer called?*doe
What is a male deer called?*buck
What is a marsupium?*pouch
What is a word for a castrated ram?*wether
What is a young whale called?*calf
What is another name for the coyote?*prairie wolf
What is another name for the prairie wolf?*coyote
What is the chihuahua named after?*A Mexican state
What is the heaviest snake?*anaconda
What is the largest lizard?*Komodo Dragon
What is the longest venomous snake?*king cobra
What is the scientific name for a turkey's wishbone?*furcula
What is the technical name for an animal's pouch?*marsupium
What lives in a formicary?*ants
What type of frog is the smallest frog?*gold frog
What was the first animal on the endangered species list?*peregrine falcon
What well known marsupial is the wallaby related to?*kangaroo
Where do ants live?*formicary
Which is the largest African bird of prey?*lammergeyer
Which is the largest aquatic bird?*albatross
Which is the largest known butterfly?*Queen Alexandra's Birdwing
Which is the only animal other than humans that can get leprosy?*armadillos
Which mammals fly?*bats
Which snake kills the most humans?*king cobra
With which island is the puffin associated?*Lundy Island
Where are there over 58 million dogs?*USA
What is the most venomous snake (no, it's not the king cobra!)?*Inland Taipan
Approximately how many years old is the first known written advertisement?*three thousand
In which ruins was the first known written advertisement found?*Thebes
What date is the 'Ides' of March?*Fifteenth
Which famous million dollar building cost more than a million dollars?*Sydney Opera House
Who painted 'Irises'?*Vincent Van Gogh
What is the art of tracing designs and making impressions of them called?*lithography
Which is the largest museum in the world?*Louvre
Who is a successful recording artist, talented landscape artist, and author of children's books?*Ricky Van Shelton
Who was born when Pluto, the astrological sign for death, was directly above Dallas, Texas?*John F. Kennedy
What is the astrological sign for death?*Pluto
What is the zodiacal symbol for Capricorn?*goat
Which constellation is represented by a goat?*Capricorn
As what is Polaris also known?*North Star
As what is the North Star also known?*Polaris
Saturday is named after which planet?*Saturn
What constellation is represented by scales?*Libra
What is the most essential tool in astronomy?*telescope
What is the name given to a group of stars?*constellation
What is the name of brightest asteroid visible from earth?*Vesta
What is the only day named after a planet?*Saturday
What is the small irregular white cloud that zips around Neptune approximately every 16 hours called?*Scooter
What is the technical name for 'falling stars'?*meteors
What planet is nearest the sun?*Mercury
When does a full moon rise?*sunset
Which is the only planet that rotates clockwise?*Venus
Who coined the theory that the earth revolves around the sun?*Nicolaus Copernicus
Who discovered the four largest moons of Jupiter?*Galileo
Who invented the telescope?*Galileo Galilei
What is the stratosphere immediately above?*troposphere
What is the troposphere immediately lower than?*stratosphere
How many 'Air Force One'(s) are there?*two
Who wore a cabbage leaf under his cap?*Babe Ruth
What drink is named after the queen of England who was famous for her 'sanguinary' persecution of the protestants?*Bloody Mary
What is made of fermented grape juice?*wine
Which animal has the largest eyes?*giant squid
As what is haemophilia also known?*royal disease
Of what is keratitis an inflammation?*cornea
On what side should you sleep to improve digestion?*right
To what disability can keratitis lead?*blindness
What appears when the sun activates melanocytes?*freckles
What body function is improved if you sleep on your right side?*digestion
What carries sensations from the tongue to the brain?*lingual nerve
What does the body release that dilates small blood vessels and so causes a person to blush?*peptides
What does the lack of iodine in the diet cause?*goitre
What does the pancreas produce?*insulin
What element is lacking in a diet when goitre occurs?*iodine
What falls out with phalacrosis?*hair
What falls out with phalacrosis?*hair
What fleshy muscular organ is joined to the hyoid bone?*tongue
What gland secretes fluid that washes the eyes?*tear gland
What is activated for freckles to appear?*melanocytes
What is the biological name for the shin bone?*tibia
What is the biological term for the voice box?*larynx
What is the common name for the larynx?*voice box
What is the hardest bone in the human body?*jawbone
What is the latin name for the top set of vertebrae?*cervical
What is the royal disease?*haemophilia
What is the tibia more commonly known as?*shin bone
What muscle is joined by the lingual nerve to the brain?*tongue
What muscles move the ears?*auricularis
What protein makes blood red?*Haemoglobin
What small region at end of medulla oblongata serves as 'bridge' to brain?*pons
When a tumour is cancerous, what is it said to be?*malignant
With age, what organ shrinks faster in males than in females?*brain
What is the birthstone for May?*emerald
What is the birthstone for September?*sapphire
Approximately how many years old are oak trees before they produce acorns?*fifty
One ragweed plant can release approximately how many grains of pollen?*one billion
To which family does the coffee plant belong?*madder
Which tree only produces acorns after it is fifty years old?*oak
How many inches tall are the bearskins worn by the guards at Buckingham Palace?*twenty
In the House of Lords, where does the Lord Chancellor sit?*woolsack
In which park are Queen Mary's gardens?*Regents Park
What are the only two london boroughs that start with the letter 'e'?*Ealing and Enfield
What does 'The Monument' in London commemorate?*Great Fire of London
What was the second bridge built across the Thames?*Westminster Bridge
Where is Selfridges?*Oxford Street, London
Which building commemorates the Great Fire of London?*Monument
Who at Buckingham Palace wears bearskins?*guards
What was Margaret Thatcher's nickname?*Iron Lady
What is the largest inhabited castle in the world?*Windsor Castle
Where is the 'whispering gallery'?*St. Paul's Cathedral
Where would you find a nave, apse, atrium and narthex?*Basilica
How is 75% of petrol in an engine wasted?*combustion
What make of car is a 'Thunderbird'?*Ford
What make of car is an 'Espace'?*Renault
Which country has the most cars per mile of road?*England
Donald Duck comics were banned in Finland because he didn't wear ----------?*pants
For which cartoon character was Beethoven a favourite composer?*Shroeder
How many freckles did Howdy Doody have?*forty eight
The maiden names of which two cartoon characters are Slaghoople and Mcbricker?*Wilma Flintstone and Betty Rubble
What are the names of Donald Duck's nephews?*Huey Dewey and Louey
What city do Batman and Robin patrol?*Gotham City
What expression did Clark Kent's newspaper boss like to use?*Great Caesar's ghost!
What is Dennis the Menace's surname?*Mitchell
What was the first cartoon character called?*Oswald the Rabbit
What were Wilma Flintstone and Betty Rubble's maiden names?*Slaghoople and Mcbricker
Which comic is drawn by Sam Keith?*The Maxx
Which magician did Lothar assist?*Mandrake
Who did the voices of Bugs Bunny, Sylvester and Tweety Pie?*Mel Blanc
Who drew the comic 'The Maxx'?*Sam Keith
Baseball: The San Diego ----------?*Padres
Who patrols Gotham City?*Batman and Robin
Who took dictation from Perry Mason?*Della Street
Who was Barney Rubble's best friend?*Fred Flintstone
Who was Fred Flinstone's best friend?*Barney Rubble
Who was born on Krypton?*Superman
Who was the black assistant of Mandrake the Magician?*Lothar
Who was the first voice of Mickey Mouse?*Walt Disney
What film was the last featuring Mel Blanc's voice?*Jetsons
20% of what is in the metal part at the end of a pencil?*sulphur
As what is sulphur also known?*brimstone
For what is the chemical formula H2O2?*hydrogen peroxide
For what metal is 'Au' the chemical symbol?*gold
Of what is 98% of the weight of water made?*oxygen
To what group of elements do cerium, praesiodymium and promethium belong?*rare earth metals
What does the symbol 'Am' represent?*americium
What is a corrosive substance with a pH value less than 7 called?*acid
What is calcium oxide commonly called?*lime
What is the atomic number for thalium?*eighty one
What is the atomic number of Bromine?*thirty five
What is the atomic number of Molybdenum?*forty two
What is the atomic number of sulphur?*sixteen
What is the atomic number of uranium?*ninety two
What is the chemical name for quicksilver?*mercury
What is the chemical symbol for gold?*Au
What is the chemical symbol for iron?*Fe
What is the heaviest naturally occuring element?*uranium
What is the symbol for copper?*Cu
What is the symbol for tin?*Sn
What term is applied to ethyl alcohol that has been treated with poison to make it unfit for human consumption?*denatured
What type of paper is used to test for acidity and alkalinity?*litmus
On what do approximately 100 people choke to death every year?*ballpoint pens
How long did it take God to create the Universe?*six days - he rested on the seventh
How many children did Noah have?*three
How many sayings did Jesus say from the cross?*seven
How many times did Peter deny Jesus?*three
How old was Sarah when she had a child?*ninety
On which day was the resurrection of Christ?*Easter Sunday
What are the first three words of The Bible?*In the beginning
What two biblical cities did God destroy with fire and brimstone?*Sodom and Gomorrah
Which two books in the Old Testament list the ten commandments? (in order of appearance)*Exodus and Deuteronomy
Who killed Goliath?*David
Who replaced Moses as the prophet of the Israelites?*Joshua
Whose name did God change to Israel?*Jacob
In the 'Twelve days of christmas', how many items in total are sent by 'my true love'?*seventy eight
A bird in the hand is worth ----------?*two in the bush
A stitch in time saves ----------?*nine
As clear as a ----------?*bell
As close as two ---------- in a pod?*peas
As easy as ----------?*pie
As hard as ----------?*nails
As mad as a ----------?*wet hen or hatter
As nervous as a long-tailed cat in a room full of ----------?*rocking chairs
As pretty as a ----------?*picture
As sick as a ----------?*dog
As sly as a ----------?*fox
Hell hath no fury like a ----------?*woman scorned
Time ---------- when you're having fun?*flies
On what is an espadrille worn?*foot
What are the essential ingredients of a daiquiri?*rum and lemon
What cocktail is based on rum and lemon?*daiquiri
In alphabet radio code, what word is used for 'c'?*charlie
In alphabet radio code, what word is used for 'f'?*foxtrot
In alphabet radio code, what word is used for 'h'?*hotel
In alphabet radio code, what word is used for 't'?*tango
In alphabet radio code, what word is used for 'x'?*X-ray
Using morse code, what does trasmitting using 3 dots, 3 dashes and 3 dots?*SOS
What is a south african coin containing 1 troy ounce of gold called?*Krugerrand
What is a group of donkeys called?*herd
What is a group of geese called?*gaggle
The De Beers group of companies controls more than 80% of the world's supply of ----------?*rough diamonds
What product built Hershey, Pennsylvania?*chocolate
Which company controls more than 80% of the world's rough diamond supply?*De Beers
What country did the operating system 'Linux' come from?*Finland
What does 'IBM' stand for?*International Business Machines
What does the 'x' mean when referring to the speed of a CD-rom (eg. 32x)?*times (faster than standard speed)
What type of printer did Seiko develop for the 1964 Tokyo Olympics?*dot matrix
What was invented over 3000 years ago that is now considered the first 'computer'?*abacus
What was the first version of Microsoft Windows to have networking capabilities?*Windows for Workgroups
What was the first version of Microsoft Windows?*Windows 286
Who is the CEO of Apple computers?*Steve Jobs
After who is the 'Ramses' brand condom named?*Pharaoh Ramses II
As what is America Online better known?*AOL
What is the most widely accepted theory for the creation of the universe?*Big Bang
What was created with the big bang?*Universe
What is kaolin?*pure china clay
What is liquid clay used in pottery called?*slip
What is pure china clay called?*kaolin
Because the emu and the kangaroo cannot walk backwards, they are on the Australian ----------?*coat of arms
For which country is the lotus flower the national symbol?*India
In which country is it polite to stick your tongue out at your guests?*Tibet
In which country is milk the most popular beverage?*USA
In which town does the famous 'running of the bulls' take place?*Pamplona
Israel has the highest per capital consumption of ----------?*turkey
What London landmark has an 11 foot long hand?*Big Ben
What animals are on the Australian coat of arms?*emu and kangaroo
What are the roads of Guam paved with?*coral
What are the sandals called that are worn in ceremonial japanese tradition?*tabi
What city do the Italians call the Monaco of bavaria?*Munich
What do the Italians call Munich?*Monaco of Bavaria
What famous building is located on the banks of the river Jumna?*Taj Mahal
What happened on screen for the first time in India in 1977?*Screen kiss
What is a water taxi known as in Venice?*gondola
What is the most common name in italy?*Mario Rossi
What is the name of a quarter of Jerusalem that can be translated as 'hundred gates'?*Mea Shearim
What is the name of the wrought iron tower in Paris?*Eiffel Tower
What is the national symbol for India?*lotus flower
What is the sacred river of Hinduism?*Ganges
What is the tribal african word for dowry?*lobola
When is turkey traditionally eaten in America?*thanksgiving
Where are the Hausa and Ibo tribes?*Nigeria
Where do the English monarchs live?*Buckingham Palace
Where is the Blarney Stone?*Blarney Castle, Ireland
Where was it once against the law to have a pet dog?*Iceland
Where would one eat a taco?*Mexico
Which country eats the most turkey per capita?*Israel
Which famous museum is in Paris, France?*Louvre
Which nationality calls Munich the 'Monaco of Bavaria'?*Italians
Germany's equivalent to the dollar is the ----------?*deutsche mark
Israel's equivalent to the dollar is the ----------?*shekel
Italy's equivalant to the dollar is the ----------?*lira
Japan's equivalent to the dollar is the ----------?*yen
Mexico's equivalent to the dollar is the ----------?*peso
Spain's equivalent to the dollar is the ----------?*peseta
The quetzal is the currency of ----------?*Guatemala
What is the Japanese currency?*yen
What is the currency of Guatemala?*quetzal
What is the monetary unit of India?*rupee
Which country has the currency 'yen'?*Japan
What country's currency is the Bolivar?*Venezuela
What is the currency of Venezuela?*Bolivar
What is armagnac?*brandy
What does a chronometer measure?*Time
What is a catalogue of languages called?*ethnologue
What is a gondola?*water taxi
What is a one-party system of government in which control is maintained by force and regimentation?*fascism
What is a word for a sorcerer who deals in black magic?*necromancer
What is another name for a sleepwalker?*somnambulist
What is ornamental work in silver or gold thread called?*filigree
What is the name given to a pregnant goldfish?*twit
What word means 'to chew the cud'?*ruminate
What has 336 dimples?*a golf ball
Sleeping sickness is carried by which insect?*tsetse fly
What disease is carried by the tsetse fly?*sleeping sickness
What is the international cry for help?*mayday
What degree do the intials 'DDS' stand for?*Doctor of Dental Surgery
What egyptian object is also known as 'the key to the Nile'?*Ankh
As what was Sony's video recorder known?*betamax
Circuits can be wired in parallel or ----------?*series
What is the only insect that can turn its head?*praying mantis
Recycling one glass jar saves enough energy to watch TV for how many hours?*three
Where is Sir Herbert Baker buried?*Westminster Abbey
For how much did an American urologist buy Napoleon's penis? (US Dollars)*$3800
For what are Allen and Wright most famous?*root beer
What Surrey town is famed for its salts?*Epsom
What is the most rural state in the USA?*North Dakota
If hell is a lake of fire, what would the temperature be? (in degrees Fahrenheit)*833
Who is the main character in 'Touched By An Angel'?*Monica
How many stars are there on the New Zealand flag?*four
What colours are on the Belgian flag?*yellow, black and red
Which country has a plain green flag?*Libya
Whose flag has the national arms on one side and the treasury seal on the other?*Paraguay
How many times its own length can the average flea jump?*150
Who was the tallest of Robin Hood's men?*Little John
How many herbs and spices are used in Kentucky Fried Chicken?*eleven
How many pieces of bun are in a Mcdonald's Big Mac?*three
In which country did edam cheese originate?*Holland
In which country did the word 'biscuit' originate?*France
What breakfast cereal was invented at Battle Creek Sanitarium?*Cornflakes
What did Charles Jung invent?*fortune cookies
What is another name for the carambula?*star fruit
What is the most widely used seasoning?*salt
What is the oldest known vegetable?*pea
Where were Cornflakes invented?*Battle Creek Sanitarium
Where were fortune cookies invented?*United States
Who invented fortune cookies?*Charles Jung
Who invented the Egg Mcmuffin?*Ed Peterson
What berries give gin its flavour?*juniper berries
A tayberry is a cross between which two fruits?*blackberry and raspberry
Unlike other oranges, what does a navel orange not have?*seeds
What fruits are usually served 'belle helene'?*pears
What is a cross between a blackberry and a raspberry?*tayberry
What is another name for the star fruit?*carambula
Where is most of the vitamin C in fruits?*skin
What is San Francisco's equivalent to Sydney's 'City To Surf' race?*Bay to Breakers footrace
What is the metal part of a lamp surrounding the bulb and supporting the shade called?*harp
Where did venetian blinds originate?*Japan
How many dots are on a twister mat?*thirty
How many folds does a Monopoly board have?*one
How many numbers are on the spinner in the game of 'Life'?*ten
How much does Park Place cost in Monopoly (in US Dollars)?*450
In a game of horseshoes, how many feet apart must the stakes be?*forty
In roulette, what number is green?*zero
Moving anti-clockwise on a dartboard, what is the number next to '4'?*eighteen
To what do opposite faces of a dice always add up?*seven
What is another name for the card game 'Blackjack'?*Twenty-one
What is another name for the card game 'Twenty-one'?*Blackjack
What is the best possible score in blackjack?*twenty one
What is the most popular sport in england?*darts
What is the tallest piece on a chessboard?*king
What number is at 12 o'clock on a dartboard?*twenty
What sport/game is Bobby Fischer associated with?*chess
Where did the card game 'bridge' originate?*Turkey
Where does the annual Poker World Series take place?*Las Vegas
Approximately how many pounds of cereal will the average american/canadian eat every year?*twelve
Peridot is the birthstone for ----------?*August
What is the birthstone for August?*peridot
During which month is the longest day in the Northern hemisphere?*June
During which month is the longest day in the Southern hemisphere?*December
During which month is the shortest day in the Northern hemisphere?*December
During which month is the shortest day in the Southern hemisphere?*June
What does a month beginning with a Sunday always have?*Friday the 13th
What game usually starts with 'is it animal, vegetable or mineral'?*twenty questions
What is the name of the office used by the president in the Whitehouse?*Oval office
What is viewed during a a pyrotechnic display?*fireworks
What system do the blind use for reading?*braille
Where will children as young as 15 be jailed for cheating on their finals?*Bangladesh
With what day does a month start if it has a Friday the 13th?*Sunday
How many chromosomes do each body cell contain?*forty six
What is the capital of Democratic Republic of the Congo?*Kinshasa
Abuja is the capital of ----------?*Nigeria
Accra is the capital of ----------?*Ghana
Albany is the capital of ----------?*New York
Ankara is the capital of ----------?*Turkey
As what is Constantinople now known?*Istanbul
As what is Formosa now known?*Taiwan
As what is Krung Thep is more commonly known?*Bangkok
As what is the South Pole also known?*Amundsen Scott Station
As what was the Taj Mahal originally built?*tomb
Austin is the capital of ----------?*Texas
Bamako is the capital of ----------?*Mali
Bangkok is the capital of ----------?*Thailand
Banjul is the capital of ----------?*Gambia
Bismarck is the capital of ----------?*North Dakota
Bissau is the capital of ----------?*Guinea-Bissau
Bogota is the capital of ----------?*Colombia
Boise is the capital of ----------?*Idaho
Bridgetown is the capital of ----------?*Barbados
Budapest is the capital of ----------?*Hungary
Cheyenne is the capital of ----------?*Wyoming
Columbus is the capital of ----------?*Ohio
Dakar is the capital of ----------?*Senegal
Des Moines is the capital of ----------?*Iowa
Dhaka is the capital of ----------?*Bangladesh
Djibouti is the capital of ----------?*Djibouti
Five US states border which ocean?*Pacific Ocean
Guatemala is the capital of ----------?*Guatemala
Helena is the capital of ----------?*Montana
How many Great Lakes are there?*five
How many countries border the black sea?*six
If you flew due West from Portugal, what is the first continent you would reach?*North America
In what state is Silicon Valley?*California
In which city is Westminster Abbey?*London
In which city is Westminster Abbey?*London
In which city is the Arch of Hadrian?*Athens
In which city is the famous Bond Street?*London
In which country is Tobruk?*Libya
In which country is the largest active volcano in the world?*Ecuador
In which county are all ten of England's highest peaks?*Cumbria
In which modern day country is ancient Troy?*Turkey
In which state is Tupelo?*Mississippi
In which state is the Natchez Trail?*Mississippi
Into what ocean does the Zambezi River empty?*Indian Ocean
Into which bay does the Golden Gate Strait lead?*San Francisco Bay
Into which estuary do the Trent and Ouse flow?*Humber
Is Belfast in Northern or Southern Ireland?*Northern
Is Dublin in Northern or Southern ireland?*Southern
Jefferson City is the capital of ----------?*Missouri
Kathmandu is the capital of ----------?*Nepal
Kigali is the capital of ----------?*Rwanda
Kingston is the capital of ----------?*Jamaica
Kinshasa is the capital of ----------?*Democratic Republic of the Congo
Kuwait City is the capital of ----------?*Kuwait
Lansing is the capital of ----------?*Michigan
Libreville is the capital of ----------?*Gabon
Lilongwe is the capital of ----------?*Malawi
Lome is the capital of ----------?*Togo
Luxembourg is the capital of ----------?*Luxembourg
Malabo is the capital of ----------?*Equatorial Guinea
Mayfair, London is a district of little streets near ----------?*Hyde Park
Mexico City is the capital of ----------?*Mexico
Montevideo is the capital of ----------?*Uruguay
Nashville is the capital of ----------?*Tennessee
Near what river is the Temple of Karnak?*Nile
New Delhi is the capital of ----------?*India
Nicosia is the capital of ----------?*Cyprus
Of what are Quemoy and Matsu part?*Taiwan
Of which country does the Kalahari Desert cover 84%?*Botswana
On the London Underground, which station has a different name on two of its platforms?*Bank and Monument
On the banks of which river is the Taj Mahal?*River Jumna
On what island is Pearl Harbor?*Oahu
On what river is Blackpool?*River Fylde
On what river is Liverpool?*Mersey
On what sea is the Crimea?*Black Sea
On which coast of Australia is Sydney?*East
Ouagadougou is the capital of ----------?*Burkina Faso
Port Louis is the capital of ----------?*Mauritius
Port Moresby is the capital of ----------?*Papua New Guinea
Raleigh is the capital of ----------?*North Carolina
Richmond is the capital of ----------?*Virginia
Riyadh is the capital of ----------?*Saudi Arabia
Rome is the capital of ----------?*Italy
Santiago is the capital of ----------?*Chile
Santo Domingo is the capital of ----------?*Dominican Republic
Singapore is the capital of ----------?*Singapore
Springfield is the capital of ----------?*Illinois
Sydney is on the east coast of ----------?*Australia
Tegucigalpa is the capital of ----------?*Honduras
Through which ocean does the International Date Line approximately follow the 180 degree meridian?*Pacific Ocean
Tirana is the capital of ----------?*Albania
Ulan Bator is the capital of ----------?*Mongolia
Vaduz is the capital of ----------?*Liechtenstein
What Central American country extends furthest north?*Belize
What Scandinavian capital begins and ends with the same letter?*Oslo
What city has the world's largest black population?*New York City
What continent is part of both the East and Aest hemispheres?*Antarctica
What country borders Egypt on the West?*Libya
What country borders Egypt to the South?*Sudan
What country borders Libya on the East?*Egypt
What country borders Sudan to the North?*Egypt
What country has the biggest population?*China
What country is situated between Panama and Nicaragua?*Costa Rica
What country is surrounded by Brazil, Argentina and Bolivia?*Paraguay
What country was once known as 'The Breadbasket of Russia'?*Ukraine
What country's capital is Caracas?*Venezuela
What divides the American North from the South?*The Mason-Dixon Line
What do Americans traditionally eat on thanksgiving day?*turkey
What does the George Washington Bridge span?*Hudson River
What is a peanut if it is not a pea or a nut?*legume
What is also known as Amundsen Scott Station?*South Pole
What is the Southernmost country in continental Europe?*Spain
What is the capital of Albania?*Tirana
What is the capital of Australia?*Canberra
What is the capital of Bangladesh?*Dhaka
What is the capital of Barbados?*Bridgetown
What is the capital of Brazil?*Brasilia
What is the capital of Burkina Faso?*Ouagadougou
What is the capital of California?*Sacramento
What is the capital of Chile?*Santiago
What is the capital of Colombia?*Bogota
What is the capital of Cyprus?*Nicosia
What is the capital of Djibouti?*Djibouti
What is the capital of Equatorial Guinea?*Malabo
What is the capital of Gabon?*Libreville
What is the capital of Gambia?*Banjul
What is the capital of Ghana?*Accra
What is the capital of Guatemala?*Guatemala
What is the capital of Guinea-Bissau?*Bissau
What is the capital of Honduras?*Tegucigalpa
What is the capital of Hungary?*Budapest
What is the capital of Idaho?*Boise
What is the capital of Illinois?*Springfield
What is the capital of India?*New Delhi
What is the capital of Iowa?*Des Moines
What is the capital of Italy?*Rome
What is the capital of Jamaica?*Kingston
What is the capital of Kuwait?*Kuwait
What is the capital of Liechtenstein?*Vaduz
What is the capital of Luxembourg?*Luxembourg
What is the capital of Luxembourg?*Luxembourg
What is the capital of Malawi?*Lilongwe
What is the capital of Mali?*Bamako
What is the capital of Mauritius?*Port Louis
What is the capital of Mexico?*Mexico City
What is the capital of Michigan?*Lansing
What is the capital of Missouri?*Jefferson City
What is the capital of Mongolia?*Ulan Bator
What is the capital of Montana?*Helena
What is the capital of Nepal?*Kathmandu
What is the capital of New York state?*Albany
What is the capital of Nigeria?*Abuja
What is the capital of North Carolina?*Raleigh
What is the capital of North Dakota?*Bismarck
What is the capital of Ohio?*Columbus
What is the capital of Papua New Guinea?*Port Moresby
What is the capital of Pennsylvania?*Harrisburg
What is the capital of Rwanda?*Kigali
What is the capital of Saudi Arabia?*Riyadh
What is the capital of Senegal?*Dakar
What is the capital of Singapore?*Singapore
What is the capital of Tennessee?*Nashville
What is the capital of Texas?*Austin
What is the capital of Thailand?*Bangkok
What is the capital of Togo?*Lome
What is the capital of Turkey?*Ankara
What is the capital of Uruguay?*Montevideo
What is the capital of Virginia?*Richmond
What is the capital of Wyoming?*Cheyenne
What is the capital of the Dominican Republic?*Santo Domingo
What is the circle of the earth at 0 degrees latitude called?*equator
What is the correct name of Bangkok?*Krung Thep
What is the deepest land gorge in the world?*Grand Canyon
What is the fifth largest country in the world?*Brazil
What is the highest peak in Fiji?*Mount Victoria
What is the largest city in China?*Shanghai
What is the largest city in Ecuador?*Guayaquil
What is the largest city in Switzerland?*Zurich
What is the largest country in Central America?*Nicaragua
What is the largest exclusively Indonesian island?*Sumatra
What is the largest ocean?*Pacific Ocean
What is the most mountainous country in Europe?*Switzerland
What is the oldest town in Belgium?*Tongeren
What is the only borough of New York City that is not on an island?*Bronx
What is the river capital of the world?*Akron
What is the saltiest sea in the world?*The Dead Sea
What is the second largest continent in the world?*Africa
What is the second largest ocean?*Atlantic Ocean
What is the second largest state in the USA?*Texas
What is the smallest Canadian province?*Prince Edward Island
What is the smallest state in the USA?*Rhode Island
What is the windiest place on earth?*Mount Washington, New Hampshire
What is the world's highest waterfall?*Angel Falls
What is the world's largest desert?*Sahara Desert
What is the world's largest lake?*Caspian Sea
What is the world's largest sea?*Mediterranean
What is the world's widest river?*Amazon
What lake is approximately 394,000 sq. km in area?*Caspian Sea
What ocean is found along the East border of Asia?*Pacific Ocean
What place is known as 'the land nowhere near'?*Cape Three Points
What seaport's name is spanish for 'white house'?*Casablanca
What small island is in the bay of Naples?*Isle of Capri
Where are the 'wallops'?*Hampshire
Where are the Nazca lines?*Peru
Where are the two steepest streets in the USA?*San Francisco
Where is Angel Falls?*Venezuela
Where is Calcutta?*India
Where is Cape Hatteras?*North Carolina
Where is Eurodisney?*Paris, France
Where is Gorky Park?*Moscow
Where is Lake Maracaibo?*Venezuela
Where is Mount Washington?*New Hampshire
Where is Tabasco?*Mexico
Where is Tongeren?*Belgium
Where is area 51 generally said to be?*Groom Lake
Where is the Blue Grotto - la Grotta Azzurra ?*Capri, Italy
Where is the Machu Picchu?*Peru
Where is the Taj Mahal?*India
Where is the bridge of San Luis Rey?*Peru
Where is the land of 10,000 lakes?*Minnesota
Where is the statue 'Le Petit Pissoir'?*Brussels
Where is the wailing wall?*Jerusalem
Where is the world's biggest prison camp?*Siberia
Where is the world's largest desert?*North Africa
Which Californian desert drops below sea level?*Death Valley
Which English county has the smallest perimeter?*Isle of Wight
Which Portuguese colony reverted to China in December 1999?*Macau
Which South American country has both a Pacific and Atlantic coastline?*Colombia
Which US state gets the most rainfall?*Hawaii
Which bridge spans the Hudson River?*George Washington Bridge
Which country administers Martinique?*France
Which country has the most emigrants?*Mexico
Which country is known as the roof of the world?*Tibet
Which country occupies the 'horn' of Africa?*Somalia
Which country owns Corfu?*Greece
Which imaginery line approximately follows the 180 degree meridian through the Pacific Ocean?*International Date Line
Which is the largest lake in South America?*Lake Maracaibo
Which is the most populated state/territory in Australia?*New South Wales
Which is the most remote island in the southern atlantic ocean?*Bouvet Island
Which is the only musical bird that can fly backwards?*hummingbird
Which is the only sea below sea level?*Dead Sea
Which is the smallest independent country?*Vatican City
Which island country lies immediately to the East of Reunion?*Mauritius
Which island country lies immediately to the West of Mauritius?*Reunion
Which island country lies to the East of Mauritius?*Australia
Which island country lies to the West of Australia?*Mauritius
Which large city is on the Southeastern coast of Australia?*Sydney
Which ocean has an area of approximately 166 sq. km?*Pacific Ocean
Which river passes through Germany, Austria, Slovakia, Hungary, Croatia, Yugoslavia, Romania, Bulgaria and Ukraine before arriving at the Black Sea?*Danube
Which tropic passes through Australia?*Tropic of Capricorn
Who owns the island of Bermuda?*Britain
Yaounde is the capital of ----------?*Cameroon
Approximately what percentage of the earth do the oceans cover?*seventy one
What is the highest active volcano in the world?*Cotopaxi
What is the most reliable geyser in the world?*Old Faithful
What type of rock is marble?*metamorphic
What is the sum of all the angles in a square? (in degrees)*three hundred and sixty
What is the glass capital of the world?*Toledo
By raising your legs slowly and laying on your back, in what can you not sink?*quicksand
Who invented the predecessor to today's computers?*Charles Babbage
After who was America named?*Amerigo Vespucci
After who was Mickey Mouse named?*Mickey Rooney
Approximately how many children did pharaoh Ramses II father?*one hundred and sixty
As what was Istanbul previously known?*Constantinople
As what was Taiwan formerly known?*Formosa
As what was winchester known by the Romans?*Venta Bulgarum
Between which countries was the shortest war in history?*Zanzibar and England
By who was Gerald Ford almost assassinated?*Squeaky Fromme
East Berlin was the capital of ----------?*East Germany
From what did Alexander the Great suffer?*epilepsy
George Washington Carver advocated planting peanuts and sweet potatoes to replace what?*cotton and tobacco
George Washington Carver advocated planting what to replace cotton and tobacco?*peanuts and sweet potatoes
Germany was split into two zones by which agreement?*Yalta agreement
His wife was Roxana, his horse was Bacephalus, he was?*Alexander the Great
How many British officers were forced by Indian troops into the Black Hole of Calcutta?*146
How many people were killed in the battle of Lexington?*eight
How many years was Nelson Mandela in prison?*twenty seven
In 1962, for what did Britain and France sign an agreement to build together?*Concorde
In 1975, what re-opened after an 8 year closure?*Suez Canal
In the 15th century, what was the war between the houses of Lancaster and York?*War of the Roses
In which battle was George A. Custer defeated?*Battle of Little Bighorn
In which country was paper money first used?*China
King Richard the ----------?*Lionhearted
Near what falls did Jimmy Angel crash his plane in 1937?*Angel Falls
Of which Cambodian party was Pol Pot the leader?*Khmer Rouge
Of which ship was Miles Standish captain?*The Mayflower
On what date did America become an independant nation?*July 4th, 1776
On what day of the week did Solomon Grundy die?*Saturday
On what was Pennsylvania incorrectly spelled?*Liberty Bell
The date of which Christian festival was fixed in 325AD by the Council of Nicaea?*Easter
The ---------- Tea Party?*Boston
Through the streets of what town did Lady Godiva ride naked?*Coventry
What 19th century war between Russia and England, Turkey, Britain and France, was named after a peninsula in the Black Sea?*Crimean War
What English city was known to the Romans as Venta Bulgarum?*Winchester
What Shakespearean king was actually king of Scotland for 17 years?*Macbeth
What United States president was in office during the civil war?*Abraham Lincoln
What colour was Diana Spencer's engagement photograph suit?*blue
What country was formerly known as Siam?*Thailand
What country was ruled by Pol Pot, leader of the Khmer Rouge party?*Cambodia (Kampuchea)
What did 'DMZ' stand for in the vietnam war?*Demilitarized Zone
What did David Stirling found?*SAS
What did Ed Peterson invent?*Egg Mcmuffin
What did Eli Whitney invent?*cotton gin
What did Erik Rotheim invent in 1926?*aerosol
What did Henry Shrapnel invent?*The exploding shell
What did Louis Cartier invent?*wristwatch
What did Marie Curie die of on 4th July, 1934?*radiation poisoning
What did Pennsylvania legalise before any other colony?*witchcraft
What did Victorian women bathe in to try to enlarge their breasts?*strawberries
What famous artist could write with both his left and right hand at the same time?*Leonardo da Vinci
What food was almost non-existent in Ireland in the 1840's?*potatoes
What is the 15' by 18' cell that 146 captured british officers were forced into by indian troops in the 19th century called?*Black Hole of Calcutta
What kind of teeth did George Washington have?*wooden
What missionary station was built by Albert Schweitzer?*Lambarene
What period is also known as the age of fish?*Devonian period
What pre-tv radio show turned film caused people to commit suicide when it was first aired?*War Of The Worlds
What war lasted from June 5 to June 11, 1967?*Six day war
What was Alaska called before 1867?*Russian America
What was Alexander The Great's wife's name?*Roxana
What was George A Custer's horses' name?*Comanche
What was King Arthur's mother's name?*Igraine
What was Russian America called after 1867?*Alaska
What was Thailand formerly known as?*Siam
What was named after Amerigo Vespucci?*America
What was the D-Day invasion password?*Mickey Mouse
What was the capital of East Germany?*East Berlin
What was the first American colony to legalise witchcraft?*Pennsylvania
What was the first fighting vehicle?*war chariot
What was the first product to have a barcode?*Wrigley's gum
What was the first ship to reach the Titanic after it sank?*Carpathia
What was the last chinese dynasty?*Manchu
What was the leading cause of death in the late 19th century?*tuberculosis
What was the name of the first ironclad warship ever launched?*HMS Warrior
What was the name of the scandal that resulted in the resignation of president Nixon?*Watergate
What wonder stood 32m high in rhodes harbour?*Colossus of Rhodes
When was D-day?*June 6th, 1944
When was Julius Caesar murdered?*Ides of March
When was the Greek alphabet first used?*800 BC
Where did 'The Mayflower' take the pilgrims?*New World
Where did Bill and Hilary Clinton switch on Christmas lights in 1995?*Belfast, Northern Ireland
Where did Churchill, Roosevelt and Stalin meet in 1945?*Yalta
Where did Guinevere retire to die?*Amesbury
When was the date of the Christian festival Easter fixed by the Council of Nicaea?*325 AD
Where did the Bay Of Pigs take place?*Cuba
Where did the Birkenhead sink?*Danger Point
Where did the bayonet originate?*Bayonne, France
Where was Napoleon defeated?*Waterloo
Where was Nelson mandela in prison?*Robben Island
Where were numerous French nuclear tests conducted?*Muraroa Atoll
Where were the Hanging Gardens?*Babylon
Where were the first books printed?*China
Which Apollo space mission put the first men on the moon ?*Apollo 11
Which Baltic seaport was the German rocket centre during WWII?*Peenemunde
What was the name of Buffy's doll in the 1970's show 'Family Affair'?*Mrs. Beasley
Which Spanish explorer first travelled to Jamaica?*Christopher Columbus
Which US president said 'the buck stops here'?*Harry Truman
Which british comedian was the first man to appear on the cover of 'playboy'?*Peter Sellers
Which country blew up a Greenpeace ship in New Zealand?*France
Which country was split into two zones by the Yalta agreement?*Germany
Which emperor made his horse a senator?*Caligula
Which famous actor is honored in a statue in Leicester Square?*Charlie Chaplin
Which famous explorer visited Australia and New Zealand, then surveyed the Pacific coast of North America?*Captain George Vancouver
Which frontiersman died at the Alamo?*Davy Crockett
Which houses fought the War of the Roses?*Lancaster and York
Which is the most ancient walled city?*Jericho
Which nation did Moshoeshoe found?*Basotho
Which nation was led by Genghis Khan?*Mongolia
Which nursery rhyme was the first gramophone recording?*Mary Had A Little Lamb
Which period was first, jurassic or carboniferous?*carboniferous
Which president was responsible for the 'Louisiana Purchase'?*Thomas Jefferson
Which racist organisation was formed in Tennessee in 1865?*Ku Klux Klan
Which film starring Julie Andrews and Christopher Plummer won the Oscar for best picture in 1965?*The Sound Of Music
Which ship did Charles Darwin captain?*HMS Beagle
Which was the first Chinese dynasty?*Shang
Which was the first magazine to publish a hologram on its cover?*National Geographic
Who advocated planting peanuts and sweet potatoes to replace cotton and tobacco?*George Washington Carver
Who appeared on the back of a US banknote in 1875?*Pocahontas
Who assassinated John Lennon?*Mark David Chapman
Who assassinated president Kennedy?*Lee Harvey Oswald
Who became president of South Africa in 1989?*F.W. de Klerk
Who built Camelot?*King Arthur
Who built the Lambarene missionary station?*Albert Schweitzer
Who built the Taj Mahal?*Shah Jahan
Who burned Atlanta in 1864?*General Sherman
Who captained the HMS Beagle?*Charles Darwin
Who committed the first daytime robbery?*Frank and Jesse James
Who developed the first nuclear submarine?*Soviet Union
Who did Squeaky Fromme try to assassinate?*Gerald Ford
Who died three days after Elvis Presley?*Groucho Marx
Who died three days before Groucho Marx?*Elvis Presley
Who discovered the Grand Canyon?*Francisco Coronado
Who drafted most of the American Declaration of Independence?*Thomas Jefferson
Who fiddled while Rome burned?*Nero
Who fixed the date of the Christian festival 'Easter'?*Council of Nicaea
Who forced 146 captured British officers into the Black Hole of Calcutta?*Indian troops
Who introduced bagpipes to the British Isles?*Romans
Who invented crop insurance?*Benjamin Franklin
Who invented the aerosol?*Erik Rotheim
Who invented the ballpoint pen?*Georgo and Laszlo Biro
Who invented the cotton gin?*Eli Whitney
Who invented the exploding shell?*Henry Shrapnel
If Brazil had won the 1998 tournament, how many times would they have won the soccer World Cup?*five
Who invented the gatling gun?*Richard Gatling
Who invented the wristwatch?*Louis Cartier
Who is considered the father of medicine?*Hippocrates
Who is identified with the word 'eureka'?*Archimedes
Who is known as the high priest of revenge?*Philip Seldon
Who is known for his 'theory of evolution'?*Charles Darwin
Who is recognised as the father of geometry?*Euclid
Who killed Jesse James?*Robert Ford
Who led 900 followers in a mass suicide in 1979?*Jim Jones
Who led the children of Israel out of Egypt?*Moses
Who led the mongols?*Genghis Khan
Who married actress Nancy Davis?*Ronald Reagan
Who met in Yalta in 1945 (in alphabetical order)?*Churchill Roosevelt Stalin
Who ordered the persecution of the Christians in which Peter and Paul died?*Nero
Who presided over the trial of Jesus?*Pontius Pilate
Who received the nobel peace prize in 1964 for civil rights leadership?*Martin Luther King Jr
Who ruled rome when Christ was born?*Augustus Caesar
Who said 'eureka'?*Archimedes
Who said 'public service is my motto'?*Al Capone
Who sailed to the new world in 'The mayflower'?*pilgrims
Who shot Abraham Lincoln?*John Wilkes Booth
Who signed the 'thanksgiving proclamation'?*Abraham Lincoln
Who started the second Punic war?*Carthage
Who succeeded Winston Churchill as Prime Minister of England?*Anthony Eden
Who tried to create the 'Great Society'?*Lyndon B Johnson
Who was 'The Elephant Man'?*Joseph Merrick
Who was George Washington's vice-president?*John Adams
Who was Joseph Merrick?*The Elephant Man
Who was King Arthur's foster-father?*Ector
Who was Ulysses' son, who grew to manhood in his absence?*Telemachus
Who was assassinated on December 8, 1980 in New York City?*John Lennon
Who was assassinated on November 22, 1963 in Dallas, Texas?*President John F. Kennedy
Who was captain of 'The Mayflower'?*Miles Standish
Who was defeated at the Battle of Little Bighorn?*George A. Custer
Who was forced by Indian troops into the Black Hole of Calcutta?*British officers
Who was given the only Nobel Peace Prize award during WWI?*International Red Cross
Who was kidnapped on the night of March 1, 1932?*Charles Lindbergh Jr
Who was known as 'the peanut president'?*Jimmy Carter
Who was the first (and last) catholic president?*John Fitzerald Kennedy
Who was the first person to break the sound barrier?*Chuck Yeager
Who was the first person to swim the English Channel?*Captain Matthew Webb
Who was the first woman in space?*Valentina Tereshkova
Who was the lead singer for Creedence Clearwater Revival, and recently released 'Blue Moon Swamp'?*John Fogerty
Who was the leader of the Khmer Rouge?*Pol Pot
Who was the only survivor of Custer's last stand?*his horse
Who were the first people to be elected into the Aviation Hall Of Fame?*The Wright Brothers
Who wrote 'The Starry Messenger'?*Galileo
What hobby was developed by the Palmer Paint Company?*Painting by numbers
Who invented painting by numbers?*palmer paint company
How many years in a vicennial?*twenty
Who hit the first golf shot on the moon?*Alan Sheppard
What does the abbreviation 'UNICEF' stand for?*United Nations Childrens' Emergency Fund
In mIRC, what colour does control-K + 4 give?*red
On IRC, how do you ask age, sex, location?*asl
On irc, what does a/s/l mean?*age/sex/location
What animal can live several weeks without its head?*cockroach
South africa is the biggest producer and exporter of ----------?*mohair
Which country is the biggest producer and exporter of mohair?*South Africa
Who founded the SAS?*David Stirling
More people are killed by donkeys every year than are killed in ----------?*plane crashes
Who invented 'bifocal' lenses for eyeglasses?*Benjamin Franklin
Who invented the most common projection for world maps?*Gerardus Mercator
In the old gag, where is Prince Albert?*In a can
From what Irish words is 'Dublin' derived?*dubh linn
From what language is the term 'finito'?*Italian
Merging the words 'melt' and 'weld' created which word?*meld
Other than Germany, whose official language is German?*Austria
The word rodent comes from the italian 'rodere', which means?*gnaw
What city's name is derived from the words 'dubh linn'?*Dublin
What does 'alma mater' mean in English?*bountiful mother
What does 'majuba' mean?*place of rock pigeons
What does 'shogun' mean in English?*military governor
What does the Irish 'dubh linn' mean?*black pool
What does the word 'karate' translate to in English?*open hand
What is 'blackpool' in Irish?*dubh linn
What is 'bountiful mother' in Latin?*alma mater
What is 'military governor' in Japanese?*shogun
What is the English word for 'fiesta'?*festival
What is the Israeli 'knesset'?*parliament
What is the Old English word for 'sneeze'?*fneasan
What is the Spanish word for 'festival'?*fiesta
What is the language of Hungary?*Magyar
What is the literal meaning of 'pince-nez'?*pinch nose
What is the meaning of the Mercedes Benz motto 'Das beste oder nichts'?*The best or nothing
What is the official language of Austria?*german
What two words make the word 'meld'?*melt and weld
What was the language of ancient India?*Sanskrit
In which country was it once against the law to slam your car door?*Switzerland
How fast does the tip of a standard rotary mower travel? (in km/h)*two hundred
What was the Lone Ranger's real name?*John Reid
Where does Nessie live?*Loch Ness
Who are santa's reindeer, in alphabetical order?*blitzen, comet, dancer, dasher, prancer and vixen
Who created the round table?*Merlin
Who was Bonnie Parker's partner?*Clyde Barrow
Who was Clyde Barrow's partner?*Bonnie Parker
Who was John Reid?*Lone Ranger
Who was the Lone Ranger's Indian companion?*Tonto
As what did H.G. Wells refer to Adolf Hitler?*A certifiable lunatic
How many books are there in Anne Rice's vampire series?*five
How many stories did enid blyton publish in 1959?*fifty nine
In 'A Christmas Carol', how many ghosts visited Scrooge?*four
In 'A Christmas Carol', what was the name of the miser?*Ebenezer Scrooge
In 'Alice In Wonderland', who never stopped sobbing?*Mock Turtle
In 'Romeo and Juliet', about what was Mercutio's long monologue?*Queen Mab
In 'Romeo and Juliet', who gave a long monologue about Queen Mab?*Mercutio
In 'Romeo and Juliet', who said 'I have a faint cold, fear thrills through my veins'?*Juliet
In 'Romeo and Juliet', who says 'make the bridal bed in that dim monument where Tybalt lies?*Juliet
In 'Romeo and Juliet', who says 'what must be must be'?*Juliet
In one of Donald Horne's novels, as what was Australia dubbed?*The lucky country
In one of Donald Horne's novels, which was 'the lucky country'?*Australia
In the Dr Seuss books, which elephant hatched an egg?*Horton
In which book did four ghosts visit Scrooge?*A Christmas Carol
On what book was 'Three Days Of The Condor' based?*Six Days Of The Condor
The Hardy Boys and ----------?*Nancy Drew
What Dr Seuss character steals Christmas?*Grinch
What controversial book did Germaine Greer write?*The Female Eunuch
What shakespearean play refers to the date of epiphany?*Twelfth Night
What story features flopsy, mopsy and cottontail?*Peter Rabbit
What subject did 'Mr. Chips' teach?*Latin
What was H.G Wells' first novel?*The Time Machine
What was Lestat's last name?*de Lioncourt
What were the dolls in the novel 'Valley Of The Dolls'?*pills
What were the two cities in 'A Tale Of Two Cities'?*London and Paris
Which Tennesee Williams play is about a Sicilian-American woman?*The Rose Tattoo
Which book featured the miser Scrooge?*A Christmas Carol
Which is the only book written by Margaret Mitchell?*Gone With The Wind
Who created 'Horton' the elephant?*Dr. Seuss
Who created 'Maudie Frickett'?*Jonathan Winters
Who created 'The Saint'?*Leslie Charteris
Who did Macduff kill?*Macbeth
Who did author Leslie Charteris create?*The Saint
Who dubbed Australia 'the lucky country'?*Donald Horne
Who killed Macbeth?*Macduff
Who said 'But, soft! what light through yonder window breaks'?*Romeo
Who was Winnie the Pooh's neighbour?*Piglet
Who was the author of 'Dracula'?*Bram Stoker
Who was the human companion of Willow?*Mad Mardigan
Who wrote '1984'?*George Orwell
Who wrote 'A Christmas Carol'?*Charles Dickens
Who wrote 'A Tale Of Two Cities'?*Charles Dickens
Who wrote 'A Tale Of Two Cities'?*Charles Dickens
Who wrote 'Alice In Wonderland'?*Lewis Carroll
Who wrote 'Gone With The Wind'?*Margaret Mitchell
Who wrote 'Psycho'?*Robert Bloch
Who wrote 'The Birds'?*Daphne du Maurier
Who wrote 'The Female Eunuch'?*Germaine Greer
Who wrote 'The Hobbit'?*J.R.R. Tolkien
Who wrote 'The Rose Tattoo'?*Tennessee Williams
Who wrote 'The Time Machine'?*H.G. Wells
Who wrote 'Valley Of The Dolls'?*Jacqueline Susann
Who wrote 'Weird Harold and Fat Albert'?*Bill Cosby
Who wrote 'little lamb, who made thee'?*William Blake
Who wrote the 'Dragonriders Of Pern' series?*Anne McCaffrey
Who wrote the 'Father Brown' crime stories?*G.K. Chesterton
Who wrote the 'Myth' series?*Robert Asprin
Who wrote the 'Noddy' books?*Enid Blyton
Who wrote the vampire series that featured Lestat as the main character?*Anne Rice
Who's last words were 'Thus with a kiss I die'?*Romeo
How old was the world's oldest man?*one hundred and forty one
Who was the world's oldest man?*Bir Narayan Chaudhari
On a ship, what is the line that indicates the maximum load that may be transported?*Plimsoll Line
How many different letters are used in the roman numeral system?*seven
What is 65% of 60?*thirty nine
What is the maximum number of integer degrees in a reflex angle?*three hundred and fifty nine
What is the maximum number of integer degrees in an acute angle?*eighty nine
What is the maximum number of integer degrees in an obtuse angle?*one hundred and seventy nine
What is the minimum number of integer degrees in a reflex angle?*one hundred and eighty one
What is the minimum number of integer degrees in an acute angle?*one
What is the minimum number of integer degrees in an obtuse angle?*ninety one
What is the only digit that has the same number of letters as its value?*four
What is the square root of one quarter?*one half
What instrument measures walking distance?*pedometer
In what body part does an osteopath specialise?*bones
A salt enema was given to children to rid them of ----------?*Threadworm
Heroin is the brand name of morphine once marketed by which pharmaceutical company?*Bayer
In the early 20th century, rattlesnake venom was used to treat which illness?*epilepsy
North American Indians ate watercress to dissolve gravel and stones in the ----------?*bladder
North American Indians ate watercress to dissolve what in the bladder?*gravel and stones
On what part of the body is an 'LTK procedure' performed?*eyes
What did North American Indians eat to dissolve gravel and stones in the bladder?*watercress
What does a sphygmomanometer measure?*blood pressure
What does hepatitis affect?*liver
What instrument measures blood pressure?*sphygmomanometer
What is a the technical name for a heart attack?*myocardial infarct
What is acute nasopharyngitis?*A cold
What is another name for consumption?*tuberculosis
What is another name for tuberculosis?*consumption
What was given to children to rid them of threadworm?*salt enema
Who ate watercress to dissolve gravel and stones in the bladder?*North American Indians
Who invented the smallpox vaccine?*Edward Jenner
Who was the first to use rubber gloves during surgery?*Dr. W. S. Halstead
As a result of their wearing high leather collars to protect their necks from sabres, as what were the first US marines known?*leathernecks
What do the letters 'SAM' mean in SAM missiles?*Surface To Air
What is the mascot of the US naval academy?*goat
What is the naval equivalent of an army Major?*Lieutenant Commander
With which hand do soldiers salute?*right
What forms when a diamond is cut with a laser?*graphite dust
What is also known as the 'bishop's stone'?*amethyst
What is the violet variety of quartz otherwise known as?*amethyst
Approximately how deep are the deepest mines? (in km)*four
In which country is the largest gold refinery?*South Africa
What is the deepest mine in the world?*Western Deep Levels Mine
What is the name of the largest gold refinery?*Rand Refinery
Where are the deepest mines?*South Africa
Graphite dust is formed when what is cut with a laser?*diamond
What are Swedish buns called?*Danishes
What has no reflection, no shadow, and can't stand the smell of garlic?*vampire
Which man has the most monuments erected in his honour?*Buddha
Which woman has the most monuments erected in her honour?*Virgin Mary
From which team did Marlboro switch its backing to Mclaren in the 1974 season?*BRM
How many pole positions did Ayrton Senna score?*sixty five
In 1976, James Hunt was disqualified after winning which Grand Prix?*British
Name the first automobile racetrack in America*Indianapolis Motor Speedway
Over what time period is the Le Mans endurance motor race?*Twenty four hours
To which team did Marlboro switch its backing from BRM in the 1974 season?*Mclaren
What colours was the Ferrari Formula 1 car in the 1964 US Grand prix?*blue and white
What event marked the 1954 french grand prix?*The return of Mercedes
Where do the Italians host the Grand Prix?*Monza
Which car won the 1953 Italian Grand Prix?*Maserati
Which new engine regulation replaced the 2.5 litre rule at the start of the 1961 season?*1.5 litre rule
Who hosts the Monza Grand Prix?*Italy
Who qualified for pole position in the 1984 Brazilian Grand Prix?*Elio de Angelis
Who was disqualified after winning the 1976 British Grand Prix?*James Hunt
Who was the driver for the Jordan team in the 1998 Grand Prix?*Damon Hill
Who won the 1966 F1 championship?*Jack Brabham
Whose motto is 'Be prepared'?*Boy Scouts
Who conquered the Matterhorn in 1865?*Edward Whymper
Who was Lauren Bacall's first husband?*Humphrey Bogart
About which family are the Godfather films?*Corleone
For which film did Art Carney win best actor Oscar in 1974?*Harry and Tonto
In 'Star Wars', who was C3P0's sidekick?*R2D2
In 'The Shining' what was the child's imaginary friend's name (the one who told him things that were going to happen)?*Tony
In the 'Nightmare On Elm Street' films, who played Freddy Krueger?*Robert Englund
In the film 'American Hot Wax', who did Jay Leno play?*Mookie
In the film 'American Hot Wax', who played the 'Mookie'?*Jay Leno
In the film 'Hackers', how old was 'zero----------kool' when he was first arrested?*eleven
In the film 'Home Alone', who played the baddies?*Joe Pesci and Daniel Stern
In the film 'Pretty Woman', for who was Goldie Hawn the body double?*Julia Roberts
In the film 'The Day Of The Jackal', who played the Jackal?*Edward Fox
In what did someone squish her hands to make the sound of e.t walking?*jelly
In what film did Whoopi Goldberg make her screen debut?*The Color Purple
In which James Bond film does the villain cheat at golf?*Goldfinger
In which film did Henry Fonda play a fallen priest?*The Fugitive
In which film did Paul Newman and Robert Redford hold hands and jump into a river?*Butch Cassidy and the Sundance Kid
In which film was Goldie Hawn the body double for Julia Roberts?*Pretty Woman
Juliette Binoche won an academy award for best supporting role in which film?*English Patient
Pancho was whose faithful sidekick?*Cisco Kid's
The film 'The Wizard Of ----------'?*Oz
Tippi Hedren is best known for her lead role in which film?*The Birds
Was Shirley Temple 21, 25 or 29 when she made her last film?*twenty one
What Marlon Brando film was widely banned?*Last Tango In Paris
What did Dorothy's house land on in 'The Wizard Of Oz'?*The Wicked Witch of the West
What film featured a cat named Mr. Bigglesworth?*Austin Powers
What film is generally considered the worst film ever made?*Attack of the Killer Tomatoes
What film marked James Cagney's return to the screen after 20 years?*Ragtime
What film starred Helen Hunt, Cary Elwes and Bill Paxton?*Twister
What film starred Rosie O'Donnell, Rita Wilson and Meg Ryan?*Sleepless in Seattle
What is the name of the film in which Steven Segal's character dies?*Executive Decision
What is the sequel to the film 'Every Which Way But Loose'?*Every Which Way You Can
What was Ben Stiller's character called in 'Mystery Men'?*Mr. Furious
What was Eddie Murphy's character name in 'Beverley Hills Cop'?*Axel Foley
What was Garth's last name in 'Wayne's World'?*Algar
What was John Wayne's real name?*Marion Morrison
What was Keanu Reeves' computer world alias in 'The Matrix'?*Neo
What was Keanu Reeves' first big film?*Point Break
What was Kevin Bacon's first big hit?*Footloose
What was painted on Peter Fonda's helmet motorcycle helmet in 'Easy Rider'?*stars and stripes
What was the first film directed by Robert Redford?*Ordinary People
What was the name of the pinball machine in the film 'Tommy'?*Wizard
What was the name of the two space shuttles in 'Armegeddon'?*Freedom and Independence
What was used for blood in the film 'psycho'?*chocolate syrup
Which basketball star played a genie in 'Kazaam'?*Shaquille O'Neal
Which film preceded 'Magnum Force' and 'The Enforcer'?*Dirty Harry
Which films are about the Corleone family?*The Godfather films
Which was the first 'Indiana Jones' film?*Raiders Of The Lost Ark
Who appeared in 'St. Elmo's Fire', 'The Scarlett Letter' and 'Striptease'?*Demi Moore
Who did Charlie Becker play in 'The Wizard of Oz'?*The mayor of the munchkins
Who directed 'The Shining'?*Stanley Kubrick
Who directed the Movie 'Psycho' from Robert Bloch?*Alfred Hitchcock
Who directed the classic thriller 'The Birds'?*Alfred Hitchcock
Who directed the film 'Ordinary People'?*Robert Redford
Who directed the film 'The Birds' from Daphne du Maurier?*Alfred Hitchcock
Who does the voice for Yoda in the Star Wars films?*Frank Oz
Who played 'Johnny Mnemonic'?*Keanu Reeves
Who played Clyde to Faye Dunaway's Bonnie?*Warren Beatty
Who played Dr. Frankenfurter in the pop-culture film 'The Rocky Horror Picture Show?*Tim Curry
Who played Dr. Kildare?*Richard Chamberlain
Who played Eddie in the pop-culture film 'The Rocky Horror Picture Show?*Meat Loaf
Who played Hopalong Cassidy?*William Boyd
Who played Louis in 'Interview With The Vampire'?*Brad Pitt
Who played Queen Amidala in the latest 'Star Wars' film?*Natalie Portman
Who played in the film 'Ragtime' after 20 years offscreen?*James Cagney
Who played the 'Universal Soldier'?*Jean-Claude Van Damme
Who played the mayor of the munchkins in 'The Wizard of Oz'?*Charlie Becker
Who played the murder victim in the original version of 'Psycho'?*Janet Leigh
Who played the president of the U.S in 'Air Force One'?*Harrison Ford
Who played the title role in the 'Mad Max' series of films?*Mel Gibson
Who played the title role in the 1978 version of 'Superman'?*Christopher Reeve
Who starred in 'Conan The Barbarian'?*Arnold Schwarzenegger
Who starred in the 1952 film 'Niagara'?*Marilyn Monroe
Who starred in the film 'The Man With Two Brains'?*Steve Martin
Who starred in the film version of 'To Kill A Mockingbird'?*Gregory Peck
Who was Dr. Zhivago's great love?*Lara
Who was John Wayne's musical co-star in true grit?*Glen Campbell
Who was Miss Hungary in 1936?*Zsa-Zsa Gabor
Who was the Cisco Kid's faithful sidekick?*Pancho
Who was the director of 'Terminator' and 'Titanic'?*James Cameron
Who was the villain in 'Star Wars'?*Darth Vader
Whose films include 'Giant', 'Written On The Wind' and 'A Farewell To Arms'?*Rock Hudson
In which film did Jay Leno play 'Mookie'?*American Hot Wax
What animal has the same name as a high church official?*cardinal
'Hang On Sloopy' was the official rock song of which band?*Ohio
'White Room' was a hit off which Eric Clapton album?*Cream
As what is Merle Haggard also known as?*Okie from Muskogee
Besides the Stones, which group had the longest touring career until the founder's death in 1995?*The Grateful Dead
Bill Justis was a studio musician when he recorded this 'sloppy' instrumental in october 1957?*Raunchy
Country singer Vince ----------?*Gill
Crosby, Stills and Nash's debut album included a song about a girl and the colour of her eyes. Name that song*Sweet Judy Blue Eyes
For whom did Colonel Tom Parker act as manager?*Elvis Presley
Formerly with Spencer Davis, he went on to form Traffic with Dave Mason. He is?*Steve Winwood
From what platform does the 'Chattanooga Choo Choo' leave Pennsylvania station?*twenty nine
From which station does the 'Chattanooga Choo Choo' leave?*Pennsylvania station
Hey! What was the name of the hit song released by 'The Romantics' in February 1980?*That's What I Like About You
How many members are in the 'fairfield four'?*five
How old was Leann Rhimes when she became a country music star?*fourteen
How old was Leann Rhimes when she recorded her first album?*eleven
In 'La Traviata', what does Violetta sing?*Sempre Libera
In 'La Traviata', who sings 'Sempre Libera'?*Violetta
In 1958, who had a pop music hit with 'Willie and the Hand Jive'?*Johnny Otis
In 1968, who released 'Carnival of life' and 'Recital'?*Lee Michaels
In 1981, who won song of the year with 'Sailing'?*Christopher Cross
In 1987, who released her second album 'Solitude Standing'?*Suzanne Vega
In a 1976 release, who wanted to 'fly like an eagle'?*Steve Miller Band
In late 1957, Buddy Holly's solo release 'Peggy Sue' challenged which song recorded with The Crickets?*Oh Boy
In the opera 'Don Giovanni', what was Leporello?*servant
In which Verdi opera does Violetta sing 'Sempre Libera'?*La Traviata
In which opera does Leporello entertain a vengeful jilted lover?*Don Giovanni
Michael di Lorenzo was one of the lead dancers on which Michael Jackson video?*Beat It
R. Kelly sings: 'If I can see it then I can do it, if I just believe it, there's nothing to it'. What's the song title?*I Believe I Can Fly
Randy Travis said his love was 'deeper than the ----------'?*holler
Savage Garden took 13 nominations and 10 wins at which awards?*ARIA awards
Singer Paula ----------?*Abdul
Sung by Robert Palmer, '---------- to love'?*Addicted
What Don Mclean song laments the day Buddy Holly died?*American Pie
What album holds the world record for copies sold?*Thriller
What are the separators on a guitar neck called?*frets
What classic rock band sang the song 'Paint It, Black'?*Rolling Stones
What did George Harrison discover on the Witwatersrand?*gold
What did Sheryl Crow do before she became a singer?*teach
What does the term 'DJ' mean?*Disc Jockey
What hardcore rock group sings, 'Blind' and 'Clown'?*Korn
What instrument are you playing when you perform a rim shot?*drums
What instrument does an organ grinder play?*hurdy gurdy
What is Cape Town's major choir called?*Philharmonic choir
What is Elton John's real name?*Reginald Dwight
What is Vanilla Ice's real name?*Robert van Winkle
What is a cello's full name?*violoncello
What is a violoncello usually called?*cello
What is the name given to the type of West Indian music made famous by artists such as Bob Marley and Peter Tosh?*reggae
What is the official birthplace of country music?*Bristol
What license plate number is on the Volkswagon on the cover of The Beatles' 'Abbey Road' Album?*281F
What song did Elton John and George Michael sing as a duet?*Don't Let The Sun Go Down On Me
What song was originally 'Good Morning To You' before the words were changed and it was published in 1935?*Happy Birthday To You
What song's words were changed and then published in 1935 as 'Happy Birthday To You'?*Good Morning To You
What was Elvis Presley's twin brother's name?*Garon
What was Jethro Tull before donating his name to a British epic rock group?*agriculturist
What was the average age of United States soldiers in the Vietnam war?*nineteen
What was the first CD pressed in the USA?*Born In The USA
What was the original name of Paul McCartney's fictional church cleaner 'Eleanor Rigby'?*Miss Daisy Hawkins
Where did George Harrison discover gold?*Witwatersrand
Where is the Rock and Roll Hall Of Fame?*Cleveland, Ohio
Which 1960's group sang a song inspired by 'Alice In Wonderland'?*The Jefferson Airplane
Which 1980's Pink Floyd album was made into a film that starred Bob Geldof, and featured the artwork of cartoonist Gerald Scarfe?*The Wall
Which Australian duo took 13 nominations and 10 wins at the ARIA awards?*Savage Garden
Which Elton John song was re-recorded as a requiem for Lady Diana Spencer?*Candle In The Wind
Which british group recorded the 1983 hit 'Owner Of A Lonely Heart'?*Yes
Which country and western singer is known as the 'okie from muskogee'?*Merle Haggard
Which singer is a former school teacher?*Sheryl Crow
Which singer/songwriter worked in a factory making toilets for airplanes before he recorded 'Aint No Sunshine'?*Bill Withers
Who 'imagined' a better world?*John Lennon
Who advised us to 'break on through to the other side'?*Jim Morrison (of The Doors)
Who appeared solo at the Woodstock festival after leaving 'The Lovin' Spoonful'?*John Sebastian
Who began his career with 'The Yardbirds' and established himself as one of the best rock guitarists of his generation?*Eric Clapton
Who began his professional career with Black Sabbath?*Ozzy Osbourne
Who collaborated with John Lennon on 'Whatever Gets You Through The Night'?*Elton John
Who did a version of 'One Bourbon, One Scotch, One Beer' on his 1977 debut album?*George Thorogood
Who did the music for the 1970's film 'Saturday Night Fever'?*Bee Gees
Who discovered gold on the Witwatersrand?*George Harrison
Who founded 'Live Aid' and 'Band Aid'?*Bob Geldof & Midge Ure
Who is Reginald Dwight known as?*Elton John
Who is Robert van Winkle?*Vanilla Ice
Who is the elder statesman of 'british blues', and fronted 'The Bluesbreakers'?*John Mayall
Who is the lead singer of 'The Doors'?*Jim Morrison
Who is the only singer to have no.1 hits in the 50's, 60's, 70's, 80's and 90's?*Cliff Richard
Who produced 'Sgt Pepper's Lonely Hearts Club Band'?*George Martin
Who recorded 'A Boy Named Sue'?*Johnny Cash
Who released 'Time, Love and Tenderness' in 1981?*Michael Bolton
Who released 'Tuesday Night Music Club' in 1993?*Sheryl Crow
Who released a chart-busting album in 1976 which featured 'The Lido Shuffle'?*Boz Scaggs
Who released the double album 'Goodbye Yellow Brick Road' in 1973?*Elton John
Who sang 'All Right Now'?*Free
Who sang 'Any Way You Want Me'?*Elvis Presley
Who sang 'Bad Case Of Loving You'?*Robert Palmer
Who sang 'Beat It'?*Michael Jackson
Who sang 'Beauty and the Beast'?*Celine Dion
Who sang 'Born In The USA'?*Bruce Springsteen
Who sang 'Forever and Ever, Amen'?*Randy Travis
Who sang 'Good Morning To You?*Mildred and Patty Hill
Who sang 'I'm A Believer'?*Monkees
Who sang 'In The Air Tonight'?*Phil Collins
Who sang 'Islands In The Stream' with Dolly Parton?*Kenny Rogers
Who sang 'Islands In The Stream' with Kenny Rogers?*Dolly Parton
Who sang 'Jet Airliner'?*Steve Miller Band
Who sang 'Rescue Me'?*Fontella Bass
Who sang 'That's Alright Mama'?*Elvis Presley
Who sang 'We've only just begun'?*Carpenters
Who sang 'You Can Call Me Al'?*Paul Simon
Who sang about 'Commitment'?*Leann Rhimes
Who sang about 'The Boogie Woogie Bugle Boy Of Company B'?*The Andrews Sisters
Who sang about Desmond and Molly Jones?*The Beatles
Who sang for 'Bad company' and 'Free', then went out on his own?*Paul Rodgers
Who sang with 'The Dakotas'?*Billy J. Kramer
Who sings 'Sweet Home Alabama'?*Lynyrd Skynyrd
Who wanted 'a lover with a slow hand'?*The Pointer Sisters
Who wanted 'a new drug'?*Huey Lewis and The News
Who was 'hooked on a feeling'?*Blue Suede
Who was a member of 'Crosby, Stills and Nash' and 'The Hollies'?*Graham Nash
Who was the Indian maiden in Johnny Preston's 'Running Bear'?*Little White Dove
Who was the first female to enter the Billboard charts in 1985?*Whitney Houston
Who was the oldest member of The Beatles?*Ringo Starr
Who was the only songwriter to win the Eurovision Song Contest twice?*Johnny Logan
Who wrote 'Roll Over Beethoven'?*Chuck Berry
Who wrote the opera 'The Giant'?*Sergei Prokofiev
Who wrote the opera 'The Masked Ball'?*Giuseppe Verdi
Who wrote the opera 'Tosca'?*Giacomo Puccini
Who wrote the opera 'norma'?*Vincenzo Bellini
Who wrote the oprea 'La Traviata'?*Guiseppe Verdi
Who wrote the song 'Do They Know It's Christmas' with Bob Geldof?*Midge Ure
Who wrote the song 'Do They Know It's Christmas' with Midge Ure?*Bob Geldof
Who's first release was 'Talking Heads 77'?*Psycho Killer
---------- in the name of love?*Stop
In Greek mythology, the riddle of what did Oedipus solve?*sphinx
Apollo was the Greek god of ----------?*prophecy and archery
Dionysus was the greek god of ----------?*Wine
In Egyptian mythology, what is the life force called?*Ka
In Egyptian mythology, who is known as the god of the desert?*Ash
In Egyptian mythology, who is the god of the underworld?*Osiris
In Egyptian mythology, who was Horus' mother?*Isis
In Egyptian mythology, who was Isis the wife of?*Osiris
In English mythology, who caused the death of the Lady of Shallot?*Sir Lancelot
In Greek mythology who did Athena turn into a spider?*Arachne
In Greek mythology, how many heads did Hydra have?*nine
In Greek mythology, into what did Athena turn Arachne?*spider
In Greek mythology, what animal is associated with Athena?*owl
In Greek mythology, what did Daedalus construct for Minos?*labyrinth
In Greek mythology, what was Minos the king of?*Crete
In Greek mythology, where did Perseus kill his grandfather?*Larrisan games
In Greek mythology, who did Jocasta marry?*Oedipus
In Greek mythology, who did Minos hire to construct the labyrinth?*Daedalus
In Greek mythology, who had nine heads?*Hydra
In Greek mythology, who hired Daedalus to construct the labyrinth?*Minos
In Greek mythology, who ruled over the island of Samos?*Polycrates
In Greek mythology, who solved the riddle of the Sphinx?*Oedipus
In Greek mythology, who turned Arachne into a spider?*Athena
In Greek mythology, who was Jason's wife?*Medea
In Greek mythology, who was Medea's husband?*Jason
In Greek mythology, who was the only mortal gorgon?*Medusa
In Greek mythology, who was the son of Peleus and Thetis?*Achilles
In Greek mythology, who were Achilles' parents?*Peleus and Thetis
Neptune was the Roman god of the ----------?*sea
Persephone was the Greek goddess of ----------?*spring
Poseidon was the Greek god of the ----------?*sea
What mythical Scottish town appears for one day every 100 years?*Brigadoon
Which Norse god had the Valkyries as handmaidens?*Odin
Which Titan had snakes for hair?*Medusa
Who did the Norse god Odin have as handmaidens?*Valkyries
Who is the Greek messenger god?*Hermes
Who is the Norse god of lightning?*Odin
Who is the Norse god of mischief?*Loki
Who is the Norse god of thunder and war?*Thor
Who is the mother of Apollo and Artemis?*Leto
Who was Hercules' father?*Zeus
Who was Hercules' stepmother?*Hera
Who was the Greek god of fire?*Hephaestus
Who was the Greek god of prophecy and archery?*Apollo
Who was the Greek god of wine?*Dionysus
Who was the Greek goddess of spring?*Persephone
Who, in Egyptian mythology, is the god of the dead?*Anubis
Which people invented the compass?*Chinese
Approximately how many pounds of salt is in every gallon of seawater?*one quarter
At which time of year do children grow fastest?*springtime
By what process is rock worn down by the weather?*erosion
During pregnancy, how many times its normal size does the human uterus expand?*five hundred
How many hearts do earthworms have?*five
Of what do earthworms have five?*hearts
What animal has bony plates and rolls up into a ball if it is frightened?*armadillo
What has approximately 1/4 pound of salt in every gallon?*seawater
What is the heart rate of the blue whale? (in beats per minute)*nine
When does the human uterus expand 500 times its normal size?*during pregnancy
Eras are divided into units called ----------?*periods
What can be either new, last or gibbous?*The moon
Who sat on her tuffet?*Little Miss Muffet
In the song 'Skip To My Lou', in what beverage are the flies?*Buttermilk
Where did Little Miss Muffet sit?*On her tuffet
Who is Mother Goose's son?*Jack
Is wholemeal bread brown or white?*brown
Lack of Vitamin D causes which disease?*rickets
Rickets is caused by a lack of which vitamin?*vitamin D
Six ounces of orange juice contains the minimum daily requirement for which vitamin?*vitamin C
What does iron deficiency cause?*anaemia
What makes brown bread healthier than white bread?*wholemeal
Basmati is a type of what?*Rice
What does a notaphile collect?*Banknotes
What does a philluminist collect?*Match box labels
What does an ombrometer measure?*rainfall
What is a 'funambulist'?*A tightrope walker
What is a pugilist?*boxer
What is another name for a  tightrope walker?*funambulist
What is someone who collects banknotes called?*Notaphile
With what is rainfall measured?*ombrometer
What is a 'somnambulist'?*sleepwalker
Which south african oil company has estblished the only commercially proven 'oil from coal' operations in the world?*Sasol
In ancient Greece, where were the original Olympics held?*Olympia
The Olympic motto 'citius, altius, fortius' means what?*Faster, higher, stronger
What is the Olympic motto in the original Latin?*Citius, altius, fortius
Where were the 1956 Summer Olympics held?*Melbourne, Australia
Where were the 1960 summer Olympics held?*Rome, Italy
Who did Zola Budd trip in the 1984 Los Angeles Olympics?*Mary Decker
Who tripped Mary Decker in the 1984 Los Angeles Olympics?*Zola Budd
How many episodes were there in the original Star Trek series?*seventy five
In 'Star Trek' Jean ---------- Picard?*Luc
In 'Star Trek', what is Data's rank?*Lieutenant Commander
In 'Star Trek', who was the captain of the 'Enterprise C'?*Rachel Garret
What is the registry number of the enterprise in the original Star Trek?*NCC 1701
Who played Deanna Troi in 'Star Trek The Next Generation'?*Marina Sirtis
What colour on black produces the most visible combination?*yellow
On maps, what is the technical name for the 'you are here' arrow?*ideo locator
The last line of which document is 'working men of all countries, unite!'?*Communist Manifesto
At what angle above the horizon must the sun be to create a rainbow? (in degrees)*forty
Waves 'break' when their height is how much more than the depth of the water?*seven tenths
What is the name given to elementary particles originating in the sun and other stars, that continuously rain down on the earth?*cosmic rays
In England, what is the Speaker of the House not allowed to do?*speak
Of which island do Ireland, Britain, Iceland and Norway dispute ownership?*Rockall
What does Israel call its parliament?*Knesset
Which island do the nationalist Chinese occupy?*Taiwan
Which nation calls its parliament 'The Knesset'?*Israel
Who succeeded Charles de Gaulle as president of France?*Georges Pompidour
What colour lenses are required to view a 3-D film?*red and green
Who invented popsicles?*Frank Epperson
As who is Vincent Furnier known?*Alice Cooper
What is Alice Cooper's real name?*Vincent Furnier
What is Conway Twitty's real name?*Harold Lloyd Jenkins
What is Harold Lloyd Jenkins' stage name?*Conway Twitty
What is Wynonna Judd's real name?*Christina Clair Ciminella
Who is Anne Mae Bullock better known as?*Tina Turner
Who is Christina Claire Ciminella otherwise known as?*Wynonna Judd
Of what did Sigmund Freud have a morbid fear?*ferns
Who wrote 'Sexual Behavior In The Human Male' in 1948?*Alfred Kinsey
With what branch of medicine is Franz Mesmer associated?*hypnotism
New York has the longest subway system in ----------?*North America
What city has the most underground stations in the world?*New York
Who is the spokesperson for the exercise tapes 'Tae Bo'?*Billy Blanks
What is the common name for lysergic acid diethylamide?*LSD
What does the 'c' in the equation e=mc^2 stand for?*speed of light
A catholic minister is known as a?*Priest
In what city does a certain church forbid burping or sneezing?*Omaha, Nebraska
Of the 266 popes, how many died violently?*thirty three
To where do Muslims make pilgrimage?*Mecca
What is God called in the Muslim faith?*Allah
What is a person who has made a pilgimage to Mecca?*Hajji
What religious movement was founded by William Booth?*Salvation Army
Who founded the Salvation Army?*William Booth
What are 35% of people using personal ads for dating?*married
St Frideswide the patron saint ----------?*Oxford
St Patrick the patron saint of ----------?*Ireland
St. Bernard the patron saint of ----------?*skiers
St. Christopher the patron saint of ----------?*travellers
Who is the patron saint of skiers?*St. Bernard
What country has the third most satellites in orbit?*France
As what is minus forty celcius the same?*minus forty fahrenheit
As what is minus forty fahrenheit the same?*minus forty celcius
How many beams of light are used to record a holograph?*two
In what does a rhinologist specialise?*human nose
In what was the strength of early lasers measured?*gillettes
In which branch of science are monocotyledon and dicotyledon terms?*Botany
Meteorology is the study of ----------?*weather
Of what did Aristotle say all things were made up?*air, earth, fire, and water
Of what is genetics the study?*heredity
Paedology is the study of ...... ?*soil
What did Wilhelm Roentgen discover in 1895?*X-rays
What does breaking the sound barrier cause?*A sonic boom
What does the Rankine scale measure?*temperature
What is name applied to the study of soil?*paedology
What is the number of blue razor blades a given beam can puncture?*gillette
What is the scientific name for brimstone?*sulphur
What is the scientific name for earth's outer layer of surface soil or crust?*lithosphere
What is the study of prehistoric plants and animals?*paleontology
What is the study of the composition of substances and the changes they undergo?*chemistry
What is the study of the earth's physical divisions termed?*Geography
What is the term that refers to the search for the existence of ghosts?*eidology
What was the first recorded message?*Mary had a little lamb
Which freezes faster - hot or cold water?*hot
Who developed the laws of electrolysis?*Michael Faraday
Who discovered X-rays?*Wilhelm Roentgen
Who first transmitted radio signals across the Atlantic?*Enrico Marconi
Who said all things were made up of air, earth, fire, and water?*Aristotle
Who spoke the first recorded message?*Thomas Edison
What is the boy scout motto?*Be prepared
Who founded the Boy Scouts?*Lord Baden Powell
How many legs does a crab have?*ten
What is cerumen?*earwax
What is the scientific name for earwax?*cerumen
What is a female calf called?*heifer
What is a female cat called?*queen
What is a male cat called?*tom
What is a person who makes barrels called?*cooper
What is a resident of Manchester called?*Mancunian
What is a resident of liverpool?*Liverpudlian
What is another word for a female sheep?*ewe
What is podobromhidrosis?*Smelly feet
What is the name given to male sheep?*ram
What is the name given to the switching of letters in an expression (e.g. saying Jag of Flapan instead of Flag of Japan)?*spoonerism
What is the study of weather technically called?*meteorology
What is the covering on the tip of a shoelace called?*aglet
After who was Deana Carter named?*Dean Martin
What instrument does Woody Allen play?*clarinet
What is Cher's maiden name?*Sarkassian
What is Tina Turner's real name?*Anne Mae Bullock
What is tattooed on Glen Campbell's arm?*dagger
What musical instrument did Jack Benny play?*violin
What was Betty Grable's nickname?*The Legs
What was Don Rickles' nickname?*Mr. Warmth
Who is Melanie Griffith's mother?*Tippi Hedren
Who is Tippi Hedren's daughter?*Melanie Griffith
Who is married to Eddie Van Halen?*Valerie Bertanelli
Who is married to Valerie Bertanelli?*Eddie Van Halen
Who married Mutt Lange?*Shania Twain
Who married Shania Twain?*Robert "Mutt" Lange
Who said 'you'd be surprised how much it costs to look this cheap'?*Dolly Parton
As neat as a ----------?*pin
In which state is the Houston Space Centre?*Texas
What is the biggest criterion for prospective astronauts?*eyesight
As who is Cassius Clay now known?*Mohammed Ali
Baseball: The Atlanta ----------?*Braves
Baseball: The Boston ----------?*Red Sox
Baseball: The Chicago ----------?*Cubs
Baseball: The Cleveland ----------?*Indians
Baseball: The Florida ----------?*Marlins
Baseball: The Houston ----------?*Astros
Baseball: The Kansas City ----------?*Royals
Baseball: The Milwaukee ----------?*Brewers
Baseball: The New York ----------?*Mets
Baseball: The Philadelphia ----------?*Phillies
Baseball: The St. Louis ----------?*Cardinals
Baseball: The Texas ----------?*Rangers
Basketball: The Denver ----------?*Nuggets
Football: The Chicago ----------?*Bears
Football: The Dallas ----------?*Cowboys
Football: The Denver ----------?*Broncos
Football: The Pittsburgh ----------?*Steelers
Football: The San Diego ----------?*Chargers
Hockey: The Calgary ----------?*Flames
Hockey: The Toronto ----------?*Maple Leafs
How many dimples does a golf ball have?*three hundred and thirty six
How many sides does a baseball homeplate have?*five
How many stitches are on a regulation baseball?*108
In baseball, who won their first world series in 1969?*New York Mets
In hockey, what is the equivalent of a rugby scrum?*face-off
In rugby, what is the equivalent of a hockey face-off?*scrum
In showjumping, how many points are incurred for knocking down a fence?*four
In what sport did the word 'crestfallen' originate?*cockfighting
In what sport do you find 'coursing'?*greyhound racing
In what sport is the term 'terminal speed' used?*Drag Racing
Other than England, which european country took part in the 1996 cricket World Cup?*Netherlands
Other than skiing, which sport takes place on a piste?*fencing
The first cricket one-day international was held between england and ----------?*Australia
What are the two basic aids in orienteering?*map and compass
What is soccer star Pele's real name?*Edson Arantes do Nascimento
What is the maximum number of clubs a golfer may use in a round?*fourteen
What is the misshapen ear that boxers often have called?*cauliflower ear
What is the name given to a rower who competes in an individual event?*sculler
What is the regulation height for a pin in tenpin bowling? (in inches)*fifteen
What is the score of a forfeited baseball game?*9-0
What is the score of a forfeited softball game?*7-0
What nationality is Gabriela Sabatini?*Argentinian
What sport has sprint, tandem and team pursuit events?*cycling
What sport is sometimes called 'rugger'?*rugby union
What sport/game is Chris Evert associated with?*tennis
What trophy is awarded to the winner of the NHL playoffs?*Stanley Cup
What was Jack Nicklaus' nickname?*Golden Bear
What was Mohammed Ali's original name?*Cassius Clay
Where is Capitol Hill?*Washington DC
Which country always leads the opening Olympic procession?*Greece
Which cricket player holds the world record for the highest individual score in first-class cricket?*Brian Lara
Which sport allows substitutions without stoppage in play?*hockey
Which tennis star wore denim shorts during matches?*Andre Agassi
Who did 'Tennis World' name rookie of the year in 1974?*Martina Navratilova
Who had the nickname 'Golden Bear'?*Jack Nicklaus
Who has played in the most consecutive baseball games?*Cal Ripken Jr
Who holds the NHL record for the most goals scored during a regular season?*Wayne Gretzky
Who hosted the 1999 cricket World Cup?*England
Who is Edson Arantes do Nascimento better known as?*Pele
Who was the 1990 Wimbledon women's singles runner-up?*Zina Garrison
Who was the NBA MVP in 1976, 1977 and 1980?*Kareem Abdul-Jabbar
Who was the NBA's most valuable player in 1976, 1977 and 1980?*Kareem Abdul-Jabbar
Who was the last Briton to win the men's singles at Wimbledon?*Fred Perry
Who was the only boxer to knock out Mohammed Ali?*Larry Holmes
Who won the 1982 soccer world cup?*Italy
With what did cricketer Mansoor Ali Khan Pataudi frequently play with in his hands?*glass eye
With what sport is Chris Boardman associated?*cycling
With what sport is Gabriela Sabatini associated?*tennis
With what sport is Jack Nicklaus associated?*golf
What swimming stroke is named after an insect?*butterfly
How many tentacles does a squid have?*ten
What were comfrey baths were believed to restore?*virginity
According to superstition, what do you make when you stub the toes on your right foot?*A wish
What did Captain Matthew Webb swim first?*English Channel
As who is Terry Bollea known?*Hulk Hogan
For which ad campaign was the line 'I can't believe I ate the whole thing' used?*Alka Seltzer
From where was Ricky in 'I Love Lucy'?*Cuba
In 'Coronation Street', who is Ken and Denise's son?*Daniel
In the TV series 'Seinfeld', who does Michael Richards play?*Kramer
In the TV series 'Seinfeld', who plays Kramer?*Michael Richards
In the TV series 'The Brady Bunch', what was Cindy's toy doll's name?*Kitty Carrie All
In the TV series 'The Fall Guy', who did Lee Majors play?*Colt Seavers
In the TV series 'The Fall Guy', who played Colt Seavers?*Lee Majors
In the TV sitcom 'Married With Children', what is the dog's name?*Buck
In the children's tv series 'Sesame Street', what two characters were roomates?*Bert and Ernie
Kelsey Grammer sings and plays the theme song for which TV show?*Frasier
On 'Dragnet', who played officer Bill Gannon?*Harry Morgan
On 'The Lucy Show', who played Vivian Bagley?*Vivian Vance
TV series: 'American ----------'?*Bandstand
To which elemetary school did TV's 'Brady Bunch' go?*Dixie Canyon Elementary
What TV network features programming just for children?*Nickelodeon
What TV series from 1970-1974 starred Susan Dey?*Partridge Family
What did Dr. David Banner become when he got angry?*The Incredible Hulk
What is Hulk Hogan's real name?*Terry Bollea
What is Kermit D Frog's girlfriend's name?*Miss Piggy
What is the drummer's name in 'The Muppet Show'?*Animal
What is the frog's name in 'The Muppet Show'?*Kermit D Frog
What is the name of Jaleel White's character in the tv series 'Family ties'?*Steve Urkel
What night club did Ricky work at on 'I Love Lucy'?*The Tropicana
What show/game has characters such as Bulbasaur and Pikachu?*Pokemon
What was Lucy's maiden name on 'I Love Lucy'?*McGillicuddy
What was the name of Ross' pet monkey on 'Friends'?*Marcel
Which famous male actor made his name in 'I Dream Of Jeannie'?*Larry Hagman
Who did Larry Hagman portray in the TV series 'Dallas'?*J.R. Ewing
Who did Pat Sajak play on the soapie 'Days Of Our Lives'?*Kevin Hathaway
Who did Patrick Duffy portray in the TV series 'Dallas'?*Bobby Ewing
Who did Vivian Vance play on 'The Lucy Show'?*Vivian Bagley
Who killed Kenny?*They
Who played Bobby Ewing in the TV series 'Dallas'?*Patrick Duffy
Who played George Costanza on 'Seinfeld'?*Jason Alexander
Who played Kevin Hathaway on the soapie 'Days Of Our Lives'?*Pat Sajak
Who played Steve Erkel in 'Family Matters'?*Jaleel White
Who played commander Riker in 'Star Trek'?*Jonathan Frakes
Who plays many voices, such as Dr Nick, and Moe on 'The Simpsons'?*Hank Azaria
Who sings and plays the theme song for the TV show 'Frasier'?*Kelsey Grammer
Who starred as 'ouboet' in the first TV series of 'Orkney Snork Nie'?*Frank Opperman
Who was the alter ego of 'The Incredible Hulk'?*Dr. David Banner
Who were Lucy and Ricky's next door neighbours and best friends?*Fred and Ethel
Over what place in india is it forbidden to fly an airplane?*Taj Mahal
Who has the world's largest double-decker tram fleet?*Hong Kong
What colour thread is used for filigree?*silver or gold
What was the name of the first space shuttle ever built?*Enterprise
In the opera 'La Traviata', what was Violetta's occupation?*courtesan
Which is the largest theme resort hotel?*Lost City
What can be tulip, balloon or flute?*wine glasses
From what were balloons originally made?*animal bladders
What toy was originally made from the bladder of an animal?*balloon
What city does Orly airport serve?*Paris
Which airline has the registration prefix 'VR'?*Cathay Pacific
Which city is served by Ringway Airport?*Manchester
5% of Canadians don't know the first seven words of the Canadian anthem, but know the first nine words of which anthem?*The American anthem
7% of Americans don't know the first nine words of the American anthem, but know the first seven words of which anthem?*Canadian anthem
Betsy Ross is the only real person to ever have been the head of a ----------?*Pez dispenser
How much do nine pennies weigh?*one ounce
Like what can a fully ripened cranberry be dribbled?*basketball
Of what are throat, foxing and platform parts?*shoe
What have woodpecker scalps, porpoise teeth and giraffe tails all been used as?*money
What is 'mother's ruin'?*gin
What is the range, in miles, of an Aim-7 Sparrow?*twenty eight
What keeps one from crying when peeling onions?*chewing gum
Who is the only real person to ever have been the head on a Pez dispenser?*Betsy Ross
How long is the longest tunnel? (in kms)*one hundred and sixty nine
The world's longest tunnel connects Delaware and ----------?*New York
The world's longest tunnel connects New York and ----------?*Delaware
What is the world's longest tunnel?*The Water Supply Tunnel
Where is the Kennedy Space Centre?*Cape Canaveral, Florida
In the USA, for how many years is a patent good?*seventeen
What is on a 5000 acre landfill at the head of Jamaica Bay near New York City?*John F. Kennedy Airport
What is the most popular street name in the US?*Park Street
What was Nancy Davis Reagan's birth name?*Anne Frances Robbins
What was the Statue Of Liberty originally named?*Liberty Enlightening The World
Where is Stone Mountain?*Atlanta
Where were Tommy Lee Jones and Al Gore freshman roommates?*Harvard University
Which city is a 'player with railroads, and the nation's freight handler'?*Chicago
Which two cities are known as the twin cities?*Minneapolis and Saint Paul
Who is the only man to have been both chief justice and president of the US?*William Taft
Who was Al Gore's freshman roommate at Harvard?*Tommy Lee Jones
Who was Tommy Lee Jones' freshman roommate at Harvard?*Al Gore
Who was born Anne Frances Robbins?*Nancy Davis Reagan
Who was born Sarah Jane Fulks?*Jane Wyman Reagan
Who was the first black mayor of Chicago?*Harold Washington
What was Jane Wyman Reagan's birth name?*Sarah Jane Fulks
If locked in a completely sealed room, of what will you die before you suffocate?*carbon monoxide poisoning
Where is the biggest calibre cannon?*Kremlin
Approximately how many times a minute does lightning strike the earth?*six thousand
Which US state holds the record for most snowfall in a day, recorded February 7, 1916?*Alaska
What is 9 metres high, 7 metres wide and 2,500 kilometres long?*Great Wall of China
Good Rhine wines are bottled in what colour bottles?*brown
How much wood can a woodchuck chuck if a woodchuck could chuck wood?*All the wood that a woodchuck could if a woodchuck could chuck wood
Which word is related to these three: rat, blue, cottage?*cheese
Complete the phrase: I need TP for my ----------*bunghole`
	listSplit := strings.Split(mirc, "\n")
	mircQuestions := make(questions, len(listSplit))

	for i, q := range listSplit {
		questionSplit := strings.Split(q, "*")
		if len(questionSplit) != 2 {
			continue
		}

		mircQuestions[i] = &question{
			Question: questionSplit[0],
			Answer:   questionSplit[1],
		}
	}
	triviaQuestions.Themes["mirc"] = mircQuestions
	triviaQuestions.All = append(triviaQuestions.All, mircQuestions...)
}
