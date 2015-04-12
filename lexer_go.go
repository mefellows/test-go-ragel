
//line lexer_go.rl:1

//line lexer_go.rl:34


package main
import "fmt"


//line lexer_go.go:12
var _simple_lexer_actions []byte = []byte{
	0, 1, 1, 1, 2, 1, 3, 1, 4, 
	1, 5, 1, 6, 1, 7, 1, 8, 
	1, 9, 2, 3, 0, 
}

var _simple_lexer_key_offsets []byte = []byte{
	0, 0, 2, 4, 9, 21, 24, 26, 
}

var _simple_lexer_trans_keys []byte = []byte{
	48, 57, 48, 57, 95, 65, 90, 97, 
	122, 32, 43, 45, 61, 9, 13, 48, 
	57, 65, 90, 97, 122, 46, 48, 57, 
	48, 57, 95, 65, 90, 97, 122, 
}

var _simple_lexer_single_lengths []byte = []byte{
	0, 0, 0, 1, 4, 1, 0, 1, 
}

var _simple_lexer_range_lengths []byte = []byte{
	0, 1, 1, 2, 4, 1, 1, 2, 
}

var _simple_lexer_index_offsets []byte = []byte{
	0, 0, 2, 4, 8, 17, 20, 22, 
}

var _simple_lexer_trans_targs []byte = []byte{
	5, 0, 6, 4, 7, 7, 7, 0, 
	4, 1, 1, 4, 4, 5, 3, 3, 
	0, 2, 5, 4, 6, 4, 7, 7, 
	7, 4, 4, 4, 4, 4, 
}

var _simple_lexer_trans_actions []byte = []byte{
	19, 0, 0, 17, 0, 0, 0, 0, 
	9, 0, 0, 7, 9, 19, 0, 0, 
	0, 0, 5, 11, 0, 13, 0, 0, 
	0, 15, 17, 11, 13, 15, 
}

var _simple_lexer_to_state_actions []byte = []byte{
	0, 0, 0, 0, 1, 0, 0, 0, 
}

var _simple_lexer_from_state_actions []byte = []byte{
	0, 0, 0, 0, 3, 0, 0, 0, 
}

var _simple_lexer_eof_trans []byte = []byte{
	0, 0, 27, 0, 0, 28, 29, 30, 
}

const simple_lexer_start int = 4
const simple_lexer_first_final int = 4
const simple_lexer_error int = 0

const simple_lexer_en_gosh int = 4


//line lexer_go.rl:40
var cs, p, pe, ts, te, act, eof, mark int
var program *Program

func emit(token_name Token, data []byte, program *Program, ts int, te int) {
  value := string(data[ts:te])
  fmt.Printf("Inserting statement %v,%v onto array", token_name, value)
  program.append(Statement{Name: token_name, Value: value})
  fmt.Printf("Program %v", program)
}

func run_machine(input string) *Program {
  program = NewProgram()
  data := []byte(input)
  fmt.Printf("Running the state machine with input %v...\n", data)
  pe = len(data)
  //eof = -1
  eof = len(data)
  
//line lexer_go.go:94
	{
	cs = simple_lexer_start
	ts = 0
	te = 0
	act = 0
	}

//line lexer_go.rl:58
  
//line lexer_go.go:104
	{
	var _klen int
	var _trans int
	var _acts int
	var _nacts uint
	var _keys int
	if p == pe {
		goto _test_eof
	}
	if cs == 0 {
		goto _out
	}
_resume:
	_acts = int(_simple_lexer_from_state_actions[cs])
	_nacts = uint(_simple_lexer_actions[_acts]); _acts++
	for ; _nacts > 0; _nacts-- {
		 _acts++
		switch _simple_lexer_actions[_acts - 1] {
		case 2:
//line NONE:1
ts = p

//line lexer_go.go:127
		}
	}

	_keys = int(_simple_lexer_key_offsets[cs])
	_trans = int(_simple_lexer_index_offsets[cs])

	_klen = int(_simple_lexer_single_lengths[cs])
	if _klen > 0 {
		_lower := int(_keys)
		var _mid int
		_upper := int(_keys + _klen - 1)
		for {
			if _upper < _lower {
				break
			}

			_mid = _lower + ((_upper - _lower) >> 1)
			switch {
			case data[p] < _simple_lexer_trans_keys[_mid]:
				_upper = _mid - 1
			case data[p] > _simple_lexer_trans_keys[_mid]:
				_lower = _mid + 1
			default:
				_trans += int(_mid - int(_keys))
				goto _match
			}
		}
		_keys += _klen
		_trans += _klen
	}

	_klen = int(_simple_lexer_range_lengths[cs])
	if _klen > 0 {
		_lower := int(_keys)
		var _mid int
		_upper := int(_keys + (_klen << 1) - 2)
		for {
			if _upper < _lower {
				break
			}

			_mid = _lower + (((_upper - _lower) >> 1) & ^1)
			switch {
			case data[p] < _simple_lexer_trans_keys[_mid]:
				_upper = _mid - 2
			case data[p] > _simple_lexer_trans_keys[_mid + 1]:
				_lower = _mid + 2
			default:
				_trans += int((_mid - int(_keys)) >> 1)
				goto _match
			}
		}
		_trans += _klen
	}

_match:
_eof_trans:
	cs = int(_simple_lexer_trans_targs[_trans])

	if _simple_lexer_trans_actions[_trans] == 0 {
		goto _again
	}

	_acts = int(_simple_lexer_trans_actions[_trans])
	_nacts = uint(_simple_lexer_actions[_acts]); _acts++
	for ; _nacts > 0; _nacts-- {
		_acts++
		switch _simple_lexer_actions[_acts-1] {
		case 0:
//line lexer_go.rl:5
 fmt.Printf("Current mark: %d, fpc: %d", mark, p) 
		case 3:
//line NONE:1
te = p+1

		case 4:
//line lexer_go.rl:22
te = p+1
{ 
      emit(ASSIGNMENT, data, program, ts, te) 
    }
		case 5:
//line lexer_go.rl:30
te = p+1

		case 6:
//line lexer_go.rl:14
te = p
p--
{ 
      emit(INTEGER, data, program, ts, te) 
    }
		case 7:
//line lexer_go.rl:18
te = p
p--
{ 
      emit(FLOAT, data, program, ts, te) 
    }
		case 8:
//line lexer_go.rl:26
te = p
p--
{ 
      emit(IDENTIFIER, data, program, ts, te) 
    }
		case 9:
//line lexer_go.rl:14
p = (te) - 1
{ 
      emit(INTEGER, data, program, ts, te) 
    }
//line lexer_go.go:240
		}
	}

_again:
	_acts = int(_simple_lexer_to_state_actions[cs])
	_nacts = uint(_simple_lexer_actions[_acts]); _acts++
	for ; _nacts > 0; _nacts-- {
		_acts++
		switch _simple_lexer_actions[_acts-1] {
		case 1:
//line NONE:1
ts = 0

//line lexer_go.go:254
		}
	}

	if cs == 0 {
		goto _out
	}
	p++
	if p != pe {
		goto _resume
	}
	_test_eof: {}
	if p == eof {
		if _simple_lexer_eof_trans[cs] > 0 {
			_trans = int(_simple_lexer_eof_trans[cs] - 1)
			goto _eof_trans
		}
	}

	_out: {}
	}

//line lexer_go.rl:59
  
  fmt.Printf("Finished. The state of the machine is: %s\n", cs)
  fmt.Printf("p: %s pe: %s\n", p, pe)
  return program
}
