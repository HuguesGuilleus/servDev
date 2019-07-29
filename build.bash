#!/bin/bash

# HTTP Server for developping project
# 2019 GUILLEUS Hugues <ghugues@netc.fr>
# BSD 3-Clause "New" or "Revised" License

clear

# include template file
cp templ/_data.go templ/data.go
for t in redirect dir notFound
do
	# cat templ/$t.gohtml > .cache_$t #######################
	tr '\n\t' '  ' <templ/$t.gohtml | sed -r -e 's/\s+/ /g' |
		sed -e 's/\(\W\) \(\W\)/\1\2/g' > .cache_$t
	sed -e "/FILE:$t/r .cache_$t" -e "/FILE:$t/d" -i templ/data.go
	rm .cache_$t
done

go build
exit $?
