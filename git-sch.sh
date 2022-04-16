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

mkdir -p $FileDir

echo "#$Today 프로그래머스" >> $FileDir/$FileName

CommitMsg= cat $FileName | head -1


#cd $GitRep
git add .
git status
git commit -m "#${CommitMsg:2}"
git push genie -f 15