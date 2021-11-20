
all: oghamtool
oghamtool: *.go
	go build
clean:
	rm -f oghamtool
test: all
# Convert into Latin.
	./oghamtool -latin "᚛ᚋᚐᚊ ᚉᚓᚏᚐᚅᚔ ᚐᚃᚔ ᚐᚈᚆᚓᚉᚓᚈᚐᚔᚋᚔᚅ᚜"
	#bivaidonas maqi mucoi 
	./oghamtool -latin "᚛ᚁᚔᚃᚐᚔᚇᚑᚅᚐᚄᚋᚐᚊᚔᚋᚒᚉᚑᚔ᚜"
	# cunavali
	./oghamtool -latin "᚛ᚉᚒᚅᚐᚃᚐᚂᚔ᚜"
	# legg sd legescad
	./oghamtool -latin "᚛ᚂᚓᚌᚌ ᚄᚇ ᚂᚓᚌᚓᚄᚉᚐᚇ᚜"
	# maq corrbri maq ammllogitt
	./oghamtool -latin "᚛ᚋᚐᚊ ᚉᚑᚏᚏᚁᚏᚔ ᚋᚐᚊ ᚐᚋᚋᚂᚂᚑᚌᚔᚈᚈ᚜"
# Convert back to Ogham.
	./oghamtool -ogham ">cunafali<" | grep --color "᚛ᚉᚒᚅ"
	# Hard letters
	./oghamtool -ogham ">ea<" | grep --color ᚕ
	# For Rob.
	./oghamtool -latin "᚛ᚏᚑᚁ ᚌᚏᚐᚆᚐᚋ ᚔᚄ ᚐ ᚈᚏᚑᚂᚂ᚜" | grep --color ROB
