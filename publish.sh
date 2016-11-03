go test . ./app ./px500
go fmt . ./app ./px500
go install

bin_dir=$GOPATH/bin
dev_dir=$GOPATH/src/github.com/pihao/desktop500px
out_dir=~/Downloads

cd $out_dir
f=$(desktop500px -v | tr ' ' '_').7z
if test -f "$f"; then rm $f; fi
7z a $f $bin_dir/desktop500px $dev_dir/key.json $dev_dir/README.md
