/*
<!-- Ну мы тут с тобой оба разработчики, давай очевидности пропустим :) -->

<!-- Принцип работы -->
Чтобы быстро понять как он работает достаточно представить фразу лесника из сталкера чистое небо "Согнуло пространство в баранку" т.е. ring buffer
Через 0 он может прыгнуть в size - 1
Помогают ему getNext getPrev(честная копипаста из проекта в проект, а исконно он появился со стаковерфлоу офкос)
Я сделал чуть иначе, чем было написано в теории и у меня tail и head указывают на существующий элемент дека при условии, что count > 0
Работают они на принципе count > 0 && tail != head тем самым представляя из себя кольцевой буфер на массиве

<!-- Сложность -->
O(n) память т.к. надо записать это дело куда-то
O(1) любая операция. У нас всегда есть указатель на текущую голову и хвост.
Там даже значения при pop не меняются в буфере т.к. просты нет нужды. Он не выдаст то, где не прошли указатели head или tail
Надеюсь бред выше ты понял :)

А если нет, то 19мс выполнения это тоже док-во

<!-- Док-во корректности -->
Могу дать котика и ссылку на посылку
Ссылка на посылку: https://contest.yandex.ru/contest/22781/run-report/52373659/

Считаю весомые аргументы

kkkkkkkkkOOOOOOOOOkkkkkkkkkkkkOOOOOOkkkkkkkkkkkkkkkkOkkkkkkkkkkkkkkkkkkkkkkkkxxxxxxxxxddooc:,''..',;clccccccldxdl:;;;::cloolc;;;;,,,''',,,,,,,,,,''',,
kkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkxxxxxxxxxddolcc:;,,'''',;:ccccldxxl::;:::cloolc;;;;,,,''',,,,,,,,,'''',,
kOOkkkkkkkkkkkkkkkkkkkkkkkkkkkxxxxxkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkxxxxxxxxxdlc:::cccc:;,'',;::::ldkxl::::::cooolc;;;;,,'''',,,,,,,,,''',,,
kkkkkkkkkkkkkkkkkkkkkkkkkkkkdl::::coxkkkOOOOOOOOkkkOOOOOOkkkkkxxkkkkkkkkkkkkkkkxxxxxxdoc;'..',:lllc;,''',,,;cdkdl::::::loool:;;;;,,'..',,,,,,,,,,,,,,,
kkkkkkkkkkkkkkkkkkkkkkkkkkxo:;;,,,,;:coxkkkOOOOOOOOOOOkkkkxxdlllloddxxxkxxxxxxxkkxxxdl:,......,coolc;''''''':oxdl::::::loool:;;;;,,'..',,,,,,,,,,,,,,,
kkkkkkkkkkkkkkkkkkkkkkkkkxo;,:cllclcc::coxkkkOOOOkkxxdlc::c:;;:::;;::clllccccllodddlc,'',''''..;odol:,'''''';lddl:::::cloool:;;;,,,'..',,,,,,,,,,,,,,,
kkkxxxxxxxxkkkkkkkkkkkkkxo:'';ldxxkxxolccldxxkkxdolc;,'...',,;:::;;;;;;;;;;;,''';;;;,',,;,''''.,cddolc:;,,,,;coolcccccclodol:;;;,,,''.',,,,,,,,,,,,,,,
kkkkkkkkxxxxxxxxxxxkkkkxdc,.';ldxkOOkkxdolllllcc;;;;:;;,,',;;:oxdlc:::::coooc;,,;;,'''''''...''':lollllccccccccccclloooooool:;;;,,,'''',,,,,,,,,,,,,,,
kkkkxxxxxxxxxxxxxxxxxxxxoc,'';ldxkOOOOkkxdlc:;,,,,;:loolc;;;,;oO0kxoollodddoc;',;;,.............;cllllllllllccccccloddddooolcc:;;,,'''',,,,,,,,,,,,,,,
ddddddooooooddddoooooooolc,..';coxkkkkxdolc:;;;,,,,;:lddl;,,;cdk0Okkxddxxdl:;..;cc,.............:lllllllllllllcccllodddddoooollcc:;,,',,,,,,,,,,,,,,,,
lllllllllllllllllooooooolc;...';lxkOOkkdlccc:cc:,,,,;ldol:;,:d0K0Okkkkkkkdol:',:cc;'',''.......'cooooloooollllllllooddddddooooooollc:::;;,,,,,,,,,,,,,
lllllllllloooooooooooooool:'...,cdOOOOxollllccc::;,,,cooc:::d0XX0OkOkkkkkdddl::;:llc:;,,,''....,looooooooooolllllloddddddddooooooollllccc:,'',,,,,,,,,
:ccllllloooooooooooooooool:,...';ok0Oxdoooolcccccc:;;coddoldOKXX0Okkkkkxxxxxdc:::looolc:;,'....;looooooooooollllllodddddddddooooooollllcc:,'';:;;;,,,'
'',,;;::cllloooooooooooool:,'..';codddooodooollllollc:::cldk0KXK0kkOOkxxxxxxdc,.....'';:;,'...':oooooooooooollllllodxxddddddooooooolllcc:;'.';:cc:::;;
....'''',,;;::ccllloooooooc;'''',:loddoddxddolllooc,''....,ck0KKOkkOOkkxxxdol;.  ......','''..,cooooooooooooolllllodxxdddddolllloollcc:;,'...,;::cccc:
.........''''',,,;::cclllooc:;;,;ldddddxxxddddooo;',,......':ok0Oxkkkkxddolc;.  .... ...,,'...':looooooooooooollllodxxddddolc::ccllc:;,'......'',;::cc
..............'''''',,,;;:ccc:::lddddddddollllldo,,;...... ..':ddxkkxxdolc::,.   ...  .;::,....,coooooooooooooolloodxxxdddol;,',;::;,'...........';::c
................''''''''',;clolcodddolc:;;;:ccdkkdc;'.........';looolc::;;,,'........,:llc;'....:looooooooooooooooddxxxdddoc,'..',,'..............';::
.......................'',cdxxdoddolc:,'':odxkkOkkxdolccllol:;,;:::;;;;,,'......',;::cc::,,'....;loooooooooooooooodxxxxdddo:,......................,;:
........................'';lxkkdoolc;,..,cloooolloddddddxxdolclloddl;'.......;clcccc::;;,''.....,coooooooooooooooodxxxxxddo:'.......................,:
........................'';cdxxoool:'...',;;;;;;:clloddddolcccldxkOkdl:,..;ldkkdlc:;;::;''......'coooooooooooooooodxxxxxddo:'.......................';
.......................'',;coxdddoc,...,;;,,;;;;::cldxxxdddxxxxkkkkkxxdlcldxkOkdoc;;:::;,'.''''.,cooooooooooolllloodddddddoc,'.......................,
.......................'',;coxxxxo:'',:lllll:;,,,;ldxxxxxkOOOOOOOkkkxxddodkOkkkkxdollcc:;,'....':looooooolllllllllllcccllllc;'.......................'
......................''',;cldddolc::codxxdoc;,,;coxxkxxxkOOOOOOOOOkxxoc;:lodxkkkxdlc:;;,,'....;loooooolllllcc:::cccc;,;;:cc:,'.......................
....................'',;:ccccccc::::ccodxdoc,'',;:lldxxxkO00OO000Okxdodkkxolccoxkxolc:;;'......:looollllooolcc:::ccllc:;;::::;,'....................''
...................',:cloolllll:::cclllool:,'...''',;:cldk000000OkkkO0000Okkxc;:llc::;;,'......;cllccccloddolllcllllllc::::c:;,'.....',;;;;;;:::::::::
.................'';coddooooooc::clllllc:;:;'.....''''',;codxxkO0KKXXKKK0OOOxl:,',,,'''''..'',:clllllooodxxdooloolloooc::ccc:;''...',::cc:::::::::::::
.................';cdddoooodoolccclollc:,,::;'....''''',,,,;:cokKXXKKKK0OOkxxdolllloooooooddddxxxxxxxxxddxxxdoooooollllccllc:,'....,::cccc::::::::::::
................';coxxxddoddddolllooolc;,,:::,'''',,,;:clloodddxkkkkkkkkkkkkkkkkkkkxxxxxxxxxxxxxxxxxxxxdddxxxdlclooc;:cclllc;,'..',:cccccc:::::::::::;
............'',;:coxkkkxdddxxxxolldddoc;,,:c:;,,,;;;:codxxxxxxxxxxkkkkkkkkkkkkxkkkkxxxxxxxxxxxxxxxxxxddddddddoccllllccccllcc:;,,;::cccccc:::::::::::::
;;:::cccllllooodddxkO0OOxddkkkxdlldxxdc;,;cc:,',,;;:codxxxxxxxxxxxkkkkkkkkkkkxxxxkxxxxxxxxxxxxxxxxxxdddddddddooooolllllllllcccccccccccc:::::::::::::::
ddddddddddddddddxxxxO0K0kddkOkxxoldkxdl:;cll:;',,;cldxxxxxxxxxxxxxxkkkxxxxkkkxxxxxxxxxxxxxxxxxxxxxdddddddddddoooooolllllllllccccccccc::::::::::::::;;;
dddddddddddddddddxxxkOOOdookOOkxolodolc::llc:;;:cldxxxxxxxxxxxxkxxkkkxxxxxxxxxxxxxxxxxxxxxxxxxxdddddddddddoooooooolllllllllccccccccc:::::::::::::;;;;;
dddddddddddddddddxxxxxxxdodxkxoolccllllllllllooddxxxxxxxxxxxxxxkkxxkxxxxxxxxxxxxxxxxxxxxxxxxddxddddddddddoooooooollllllllccccccccc:::::::::::::::;;;;;
ddddddddddddddddddxxxxxxdddddddddddddddddddddxddxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxddddddddddddddoooooolllllllllcccccccccc:::::::::::::::;;;;;;
ddddddddddddddddddddxxxxxxddddddddxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxdddddddddddoooooooooolllllllcccccccccccc:::::::::::::;;;;;;;;
ddddddddddddddddddddxxxxxxxxddddxxxxxxxxxxxxxxddxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxddddddddddddddoooooooooooollllllccccccccccccc:::::::::::::;;;;;;;;;
ooodddddddddddddddddddddxddxxxxxxxxddxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxddddddddddddoooooooooooooolllllllcccccccccc::::::::::::::;;;;;;;;;;;;
oooooodddddddddddddddddddddddxxxxdddxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxdddddddddddddddddooooooooooolllllllllccccccccccc::::::::::::::;;;;;;;;;;;;;
oooooooddddddddddddddddddddddddddddxxxxxxxxxxxxxddddddxxxxxxxxddddddddddddddddddddddddddoooooooooooollllllllllccccccccccc::::::::::::::;;;;;;;;;;;;;;;
oooooooooooddddddddddddddddddddddddddddddxxxxxxdddddddddddddddddddddddddddddddddddddooooooooooollllllllllllllcccccccc::::::::::::::;;;;;;;;;;;;;;;;;;,
ooooooooooooddddddddddddddddddddddddddddddddxdddddddddddddddddddddddddddddddddddddoooooooooooollllllllllcccccccccc::::::::::::::::;;;;;;;;;;;;;;;;;;,,
*/

package main

import (
	"bufio"
	"errors"
	"os"
	"strconv"
	"strings"
)

func main() {
	initReader()
	rawStrings := readStrings()
	dequeSize, err := strconv.Atoi(rawStrings[1])
	if err != nil {
		return
	}

	deque := NewDeque(dequeSize)
	for _, str := range rawStrings[2:] {
		toExecute := strings.Split(str, " ")
		if len(toExecute) > 1 {
			val, _ := strconv.Atoi(toExecute[1])
			switch strings.TrimSpace(toExecute[0]) {
			case "push_front":
				err := deque.PushFront(val)
				if err != nil {
					writeData("error")
				}
				break
			case "push_back":
				err := deque.PushBack(val)
				if err != nil {
					writeData("error")
				}
				break
			}
		} else {
			switch strings.TrimSpace(toExecute[0]) {
			case "pop_front":
				val, err := deque.PopFront()
				if err != nil {
					writeData("error")
				} else {
					strVal := strconv.Itoa(val)
					writeData(strVal)
				}
				break
			case "pop_back":
				val, err := deque.PopBack()
				if err != nil {
					writeData("error")
				} else {
					strVal := strconv.Itoa(val)
					writeData(strVal)
				}
				break
			}
		}
	}
}

//region Basic
var reader *bufio.Reader
var scanner *bufio.Scanner

func initReader() {
	reader = bufio.NewReader(os.Stdin)
	scanner = bufio.NewScanner(reader)
	scanner.Split(bufio.ScanLines)
}

func readStrings() []string {
	var data []string

	for scanner.Scan() {
		rawString := scanner.Text()
		data = append(data, rawString)
	}

	return data
}

func writeData(data string) {
	writer := bufio.NewWriter(os.Stdout)
	writer.WriteString(data + "\n")
	writer.Flush()
}

//endregion

//region deque

type Deque struct {
	maxSize int
	buffer  []int
	head    int
	tail    int
	count   int
}

func NewDeque(size int) *Deque {
	return &Deque{
		maxSize: size,
		buffer:  make([]int, size),
		head:    0,
		tail:    0,
		count:   0,
	}
}

func (deque *Deque) PushFront(number int) error {
	if deque.head == deque.tail && deque.count == 0 {
		deque.buffer[deque.head] = number
		deque.count++
		return nil
	}

	nextHead := deque.getPrev(deque.head)
	if nextHead == deque.tail {
		return errors.New("is full")
	}

	deque.head = nextHead
	deque.buffer[deque.head] = number
	deque.count++
	return nil
}

func (deque *Deque) PushBack(number int) error {
	if deque.head == deque.tail && deque.count == 0 {
		deque.buffer[deque.tail] = number
		deque.count++
		return nil
	}

	nextTail := deque.getNext(deque.tail)
	if nextTail == deque.head {
		return errors.New("is full")
	}

	deque.tail = nextTail
	deque.buffer[deque.tail] = number
	deque.count++
	return nil
}

func (deque *Deque) PopFront() (int, error) {
	if deque.head == deque.tail {
		if deque.count != 0 {
			deque.count--
			return deque.buffer[deque.head], nil
		}
		return 0, errors.New("is empty")
	}

	oldHead := deque.head
	deque.head = deque.getNext(deque.head)
	deque.count--

	return deque.buffer[oldHead], nil
}

func (deque *Deque) PopBack() (int, error) {
	if deque.head == deque.tail {
		if deque.count != 0 {
			deque.count--
			return deque.buffer[deque.tail], nil
		}

		return 0, errors.New("is empty")
	}

	oldTail := deque.tail
	deque.tail = deque.getPrev(deque.tail)
	deque.count--

	return deque.buffer[oldTail], nil
}

func (deque *Deque) getNext(index int) int {
	return (index + deque.maxSize + 1) % deque.maxSize
}

func (deque *Deque) getPrev(index int) int {
	return (index + deque.maxSize - 1) % deque.maxSize
}

//endregion
