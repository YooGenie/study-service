Y=$(date +%Y)
M=$(date +%m)
D=$(date +%d)

Ym=$Y-$M
Day=$Y$M$D
Today=$Y-$M-$D
GitRep="study-service"

GitDir="$HomeDir/$GitRep"
FileDir="$HomeDir/$GitRep/auto"
FileName="$Day".go

#mkdir -p $FileDir

echo "#$Today 프로그래머스" >> $FileDir/$FileName


#cd $GitRep


git add .
git status
git commit -m "#$Today 프로그래머스"
git push genie -f 15