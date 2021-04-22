# hrs (hours)
hrs is a small CLI tool to calculate working hours each month, it takes into account the danish holidays.

I personally use this tool for budgetting purposes when working as a contractor to calculate a baseline for number of hours I can work each month for a whole year.

## install
install by using 
```
go install github.com/eikc/hrs
```

## usage
Current hrs will default to the current year with a default of 7.5 working hours each day. 

```
hrs 

#outputs:
❯ hrs
Total working hours in 1/2021 : 150.000000 
Total working hours in 2/2021 : 150.000000 
Total working hours in 3/2021 : 172.500000 
Total working hours in 4/2021 : 135.000000 
Total working hours in 5/2021 : 142.500000 
Total working hours in 6/2021 : 165.000000 
Total working hours in 7/2021 : 165.000000 
Total working hours in 8/2021 : 165.000000 
Total working hours in 9/2021 : 165.000000 
Total working hours in 10/2021 : 157.500000 
Total working hours in 11/2021 : 165.000000 
Total working hours in 12/2021 : 172.500000
```

You can adjust the working hours or year by using the `-year` or `-hours` flag

```
hrs -year 2022 -hours 8

#outputs
❯ hrs -year 2022 -hours 8
Total working hours in 1/2022 : 168.000000 
Total working hours in 2/2022 : 160.000000 
Total working hours in 3/2022 : 184.000000 
Total working hours in 4/2022 : 144.000000 
Total working hours in 5/2022 : 160.000000 
Total working hours in 6/2022 : 168.000000 
Total working hours in 7/2022 : 168.000000 
Total working hours in 8/2022 : 184.000000 
Total working hours in 9/2022 : 176.000000 
Total working hours in 10/2022 : 168.000000 
Total working hours in 11/2022 : 176.000000 
Total working hours in 12/2022 : 168.000000 
```
