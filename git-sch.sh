Y=$(date +%Y)
M=$(date +%m)
D=$(date +%d)

Day=$Y$M$D
Today=$Y-$M-$D
GitRep="study-service"

HomeDir="/home/ubuntu"
GitDir="$HomeDir/$GitRep"
FileDir="$HomeDir/$GitRep/auto"
FileName="$Day".go

mkdir -p $FileDir

#git checkout -b 23

echo "
import (
	"strconv"
)

func solution(s string) bool {
    var result bool

    if len(s)==4 || len(s)==6 {
        if c, err :=strconv.ParseInt(s,10	,64); err !=nil {
		result = false
	}else if c!=0 {
		result = true
	}
    }



    return result
    //연습
}
" >> $FileDir/$FileName


cd $GitDir
git add .
git commit -m "#$Today 프로그래머스2"
git pull genie master
git push genie -f  23
