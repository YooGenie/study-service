Y=$(date +%Y)
M=$(date +%m)
D=$(date +%d)

Ym=$Y-$M
Day=$Y$M$D
Today=$Y-$M-$D
GitRep="study-service"

#HomeDir= "/home/ubuntu"
GitDir="$GitRep"
FileDir="$GitRep/auto"
FileName= "$Day_1".go

#mkdir -p $FileDir

echo "#$Today 프로그래머스" >> $FileDir/$FileName


#cd $GitRep
git add .
git commit -m "#$Today 프로그래머스"
git push genie -f 17
