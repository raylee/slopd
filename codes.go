package main

// I suspect these aren't universal, taken from http://www.minneapolismn.gov/www/groups/public/@mpd/documents/webcontent/convert_257091.pdf

// outcome of incident
const disposition_codes = `ADV Advised NOS No service
	AOK All okay
	NTF Notify
	AOT All quiet
	PEM Unit not sent
	AST Assisted
	RFD Refused service
	BKG Booking (an arrest)
	RPR Reprimand and release
	CNL Cancel (call cancelled)
	RPT Report made
	DTX Party to detox
	SCK Sick person
	FAL False alarm
	SEC Secured
	FTC Fail to clear
	SNT Sent one
	GOA Gone on arrival
	TAG Citation issued
	INF Information received
	TOW Vehicle towed
	INS In service
	TRN Transported one
	LFT Left one at the scene
	UNF Complaint unfounded
	MES Message delivered
	WRN Warning issued
`

// nature of the call
const nature_codes = `
	ABITE Animal bite
	DROWN Drowning
	ABURGB Attempted burglary of business
	DOMESW Domestic fight, with weapons
	ABURGD Attempted burglary of dwelling
	DOWN One down check for ambulance
	AC Animal call
	E ELEV Elevator emergency (one trapped)
	ACHILD Abandoned child
	E HELP EMS needs HELP
	ALERT Aircraft crash alert (possible)
	EXPLOS Explosion
	ALRMA Alarm (audible)
	F ALRM Fire alarm
	ALRMAB Alarm business
	F BLDG Fire in a building (assist fire department)
	ALRMAR Alarm residence
	F GARR Fire in a garage (assist fire department)
	ALRMH Alarm (hold up)
	F GRAS Grass fire
	ALRMR Alarm (recorded)
	F VEH Vehicle fire
	ALRMS Alarm (silent)
	F HELP Fire needs help
	AOA Assist other agency
	F MISC Miscellaneous fire
	ASLT Assault
	F OTS Fire outside
	ASLTP Assault in progress
	F OUT Fire out
	ASTDIS Replaces ASTINV
	F SMKA Fire smoke alarm
	ASTEMS Assist EMS (Urgent but not HELP)
	FC Firecrackers
	ASTOFF Assist an officer
	FCHILD Found child
	ASUIC Attempt suicide
	FIGHT Fight
	ATTPU Attempt pick -up (usually a law enforcement request)
	FIGHTW Fight with weapons
	AUTOTH Auto theft
	FLEE Fleeing suspect
	BABY Baby not breathing
	FORG Forgery
	BAIT Bait vehicle activation
	HAZMAT Hazardous material
	BARGE Loose barge
	HEART Heart attack
	BOMB Bomb (suspected)
	HELP Officer needs help
	BOMBT Bomb threat
	HOTROD Hotrodder (vehicle disturbance)
	BOOK Booking (police initiated arrest)
	INDEX Criminal sexual conduct (exposure)
	BURGB Burglary of business
	INFO Receive information
	BURGBP Burglary of business in progress
	JUMPER Bridge or building jumper (suicide attempt or threat)
	BURGD Burglary of dwelling
	KIDNAP Person kidnapped
	BURGDP Burglary of dwelling in progress
	KIDTRB Kid trouble
	CHASE Vehicle chase
	LCHILD Lost child
	CKHAZ Check for a hazard
	LKIN Person locked in (vehicle or building)
	CKWEL Check the welfare of a person
	MEDIC Medical assistance needed
	CRASH Aircraft crash
	MISC Miscellaneous (doesn’t fit other codes)
	COALRM Carbon monoxide alarm
	MPER Missing person
	CRANK Crank 911 caller
	MUSIC Loud music disturbance
	CSCM Criminal sexual conduct (molester)
	MYSDIS Mysterious disappearance
	CSCR Criminal sexual conduct (rape)
	NARC Narcotics call
	CURFEW Curfew violation detention
	NBRTBR Neighbor trouble
	CUSTRB Customer trouble
	NOPAY Non-paying customer (left scene)
	DABUSE Domestic abuse (family or household assault or threats)
	NOTIFY Notification; deliver message
	DAMPRP Damage to property
	OB Maternity run
	DIST Disturbance (various types)
	OD Overdose
	DK Drunk ODOR Noxious smell; combinations include
	DOA Dead body
	GDOR - gas order
	ODORIN - odor inside
	DOMES Domestic (family/household argument)
	PARTY Loud party causing a disturbance
	PD Property damage accident (vehicle)
	SUSPAK Suspicious package
	PDHR Hit & Run property damage (vehicle)
	SUSPP Suspicious person
	PI Personal injury accident
	SUSPV Suspicious vehicle
	PIHR Hit & Run personal injury accident
	TENTRB Tenant trouble
	PKG Parking problem
	THEFT Theft
	PKGBD Parking problem, blocked drive
	THEFTA Theft from an auto
	PPI Vehicle accident, possible injury
	THEFTH Theft holding (shoplifting)
	PROWL Prowler
	THEFTP Theft in progress
	PERGUN Person with a gun
	THREAT Threat made against another person
	PERWEA Person with a weapon
	TLE Traffic law enforcement (MV stop)
	RDHAZ Road hazard
	TOW Vehicle towed
	RECPRP Recovered property
	TRANS Transportation request
	RECVEH Recovered vehicle
	TRESP Trespassing
	RETPRP Retrieve personal property from former residence
	TRFCN Traffic control
	TRUANT Truancy, juvenile
	RISK High risk warrant served
	UNCON Unconscious person
	ROBBIZ Robbery of business
	UNKCEL Unknown trouble from cell phone
	ROBBZP Robbery of business in progress
	UNKTRB Unknown trouble
	ROBDWL Robbery of dwelling
	UNSBIZ Unsecured business
	ROBPER Robbery of person
	UNWANT Unwanted person
	S BLDG Smoke inside a building
	WALKTH Walk through a building
	SAFE Problem address
	WATEM Water emergency
	SEIZ Seizure (medical)
	WIREDN Wires down
	SHOOT Shooting victim
	SHOTS Sound of shots heard
	SICK Sick person
	SLUMP Person slumped over
	SOB Short of breath (medical)
	SPILL Hazardous spill
	STAB Stabbing/cutting
	STROKE Stroke
`
