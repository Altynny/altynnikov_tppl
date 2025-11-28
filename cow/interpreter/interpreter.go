package interpreter

import (
	"errors"
	"fmt"
	"strings"
)

type command int

const (
	command_moo command = iota
	command_mOo
	command_moO
	command_mOO
	command_Moo
	command_MOo
	command_MoO
	command_MOO
	command_OOO
	command_MMM
	command_OOM
	command_oom
)

var commandName = map[command]string{
	command_moo: "moo",
	command_mOo: "mOo",
	command_moO: "moO",
	command_mOO: "mOO",
	command_Moo: "Moo",
	command_MOo: "MOo",
	command_MoO: "MoO",
	command_MOO: "MOO",
	command_OOO: "OOO",
	command_MMM: "MMM",
	command_OOM: "OOM",
	command_oom: "oom",
}

var nameToCommand map[string]command

func init() {
	nameToCommand = make(map[string]command, len(commandName))
	for cmd, name := range commandName {
		nameToCommand[name] = cmd
	}
}

type interpreter struct {
	code   []command
	pos    int
	loops  map[int]int
	buffer []byte
	ptr    int
}

func Interpreter() *interpreter {
	inter := interpreter{
		code:   make([]command, 0),
		pos:    0,
		loops:  make(map[int]int),
		buffer: make([]byte, 1),
		ptr:    0}
	return &inter
}

func (inter *interpreter) clear() {
	inter.code = make([]command, 0)
	inter.pos = 0
	inter.loops = make(map[int]int)
	inter.buffer = make([]byte, 1)
	inter.ptr = 0
}

func (inter *interpreter) Interpret(source string) error {
	inter.parseCode(source)
	for ; inter.pos < len(inter.code); inter.pos++ {
		cmd := inter.code[inter.pos]
		err := inter.executeCommand(cmd)
		if err != nil {
			inter.clear()
			return err
		}
	}
	inter.clear()
	return nil
}

func (inter *interpreter) parseCode(source string) {
	tokens := strings.Fields(source)
	var last_MOO_pos int
	pos := 0
	for _, t := range tokens {
		if cmd, prs := nameToCommand[t]; prs {
			inter.code = append(inter.code, cmd)
			switch cmd {
			// поиск циклов
			case command_MOO:
				last_MOO_pos = pos
			case command_moo:
				inter.loops[last_MOO_pos] = pos
				inter.loops[pos] = last_MOO_pos
			}
			pos++
		}
	}
}

func (inter *interpreter) executeCommand(cmd command) error {
	current_block := &inter.buffer[inter.ptr]
	switch cmd {

	// поиск команды "MOO" в обратном порядке и продолжение исполнения с неё
	case command_moo:
		inter.pos = inter.loops[inter.pos] - 1

	// предыдущая ячейка буффера
	case command_mOo:
		inter.ptr -= 1

	// следующая ячейка буффера
	case command_moO:
		inter.ptr += 1
		if inter.ptr >= len(inter.buffer) {
			inter.buffer = append(inter.buffer, 0)
		}

	// выполнить инструкцию с номером из текущей ячейки
	case command_mOO:
		number := *current_block
		new_cmd := command(number)
		_, prs := commandName[new_cmd]
		if new_cmd == command_mOO || !prs {
			return errors.New("invalid command")
		} else {
			inter.executeCommand(new_cmd)
		}

	// если значение в ячейке равно 0, то ввести с клавиатуры, иначе вывести на экран в виде ASCII
	case command_Moo:
		if *current_block == 0 {
			inter.executeCommand(command_oom)
		} else {
			fmt.Print(string(*current_block))
		}

	// значение ​текущей ячейки уменьшить на 1
	case command_MOo:
		*current_block -= 1

	// значение текущей ячейки увеличить на 1
	case command_MoO:
		*current_block += 1

	// при нулевом значении текущей ячейки все команды пропускаются и исполнение продолжается после "moo"
	case command_MOO:
		if *current_block == 0 {
			inter.pos = inter.loops[inter.pos]
		}

	// обнулить значение в ячейке
	case command_OOO:
		*current_block = 0

	// вывод int значения текущей ячейки
	case command_OOM:
		fmt.Print(string(*current_block))

	// ввод значения в текущую ячейку
	case command_oom:
		var newValue int
		_, err := fmt.Scanf("%d", &newValue)
		if err != nil {
			return err
		}
		*current_block = byte(newValue)
	}
	return nil
}
