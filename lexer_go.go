
//line lexer_go.rl:1

//line lexer_go.rl:43


package main
import "fmt"


//line lexer_go.go:12
var _simple_lexer_actions []byte = []byte{
	0, 1, 0, 1, 1, 1, 2, 1, 3, 
	1, 4, 1, 5, 1, 6, 1, 7, 
	1, 8, 1, 9, 1, 10, 
}

var _simple_lexer_key_offsets []byte = []byte{
	0, 0, 2, 4, 9, 15, 21, 27, 
	42, 45, 47, 52, 58, 64, 69, 69, 
	75, 81, 87, 93, 99, 
}

var _simple_lexer_trans_keys []byte = []byte{
	48, 57, 48, 57, 95, 65, 90, 97, 
	122, 95, 111, 65, 90, 97, 122, 95, 
	105, 65, 90, 97, 122, 95, 117, 65, 
	90, 97, 122, 32, 43, 45, 61, 99, 
	104, 115, 9, 13, 48, 57, 65, 90, 
	97, 122, 46, 48, 57, 48, 57, 95, 
	65, 90, 97, 122, 95, 112, 65, 90, 
	97, 122, 95, 121, 65, 90, 97, 122, 
	95, 65, 90, 97, 122, 95, 115, 65, 
	90, 97, 122, 95, 116, 65, 90, 97, 
	122, 95, 111, 65, 90, 97, 122, 95, 
	114, 65, 90, 97, 122, 95, 100, 65, 
	90, 97, 122, 95, 111, 65, 90, 97, 
	122, 
}

var _simple_lexer_single_lengths []byte = []byte{
	0, 0, 0, 1, 2, 2, 2, 7, 
	1, 0, 1, 2, 2, 1, 0, 2, 
	2, 2, 2, 2, 2, 
}

var _simple_lexer_range_lengths []byte = []byte{
	0, 1, 1, 2, 2, 2, 2, 4, 
	1, 1, 2, 2, 2, 2, 0, 2, 
	2, 2, 2, 2, 2, 
}

var _simple_lexer_index_offsets []byte = []byte{
	0, 0, 2, 4, 8, 13, 18, 23, 
	35, 38, 40, 44, 49, 54, 58, 59, 
	64, 69, 74, 79, 84, 
}

var _simple_lexer_indicies []byte = []byte{
	0, 1, 3, 2, 4, 4, 4, 1, 
	4, 5, 4, 4, 1, 4, 6, 4, 
	4, 1, 4, 7, 4, 4, 1, 8, 
	9, 9, 10, 12, 13, 14, 8, 0, 
	11, 11, 1, 16, 0, 15, 3, 17, 
	4, 4, 4, 18, 4, 19, 4, 4, 
	18, 4, 20, 4, 4, 18, 20, 20, 
	20, 22, 22, 4, 23, 4, 4, 18, 
	4, 24, 4, 4, 18, 4, 25, 4, 
	4, 18, 4, 19, 4, 4, 18, 4, 
	26, 4, 4, 18, 4, 20, 4, 4, 
	18, 
}

var _simple_lexer_trans_targs []byte = []byte{
	8, 0, 7, 9, 10, 11, 15, 19, 
	7, 1, 7, 3, 4, 5, 6, 7, 
	2, 7, 7, 12, 13, 7, 14, 16, 
	17, 18, 20, 
}

var _simple_lexer_trans_actions []byte = []byte{
	7, 0, 21, 0, 0, 0, 0, 0, 
	11, 0, 9, 0, 1, 1, 1, 19, 
	0, 15, 17, 0, 0, 13, 0, 0, 
	0, 0, 0, 
}

var _simple_lexer_to_state_actions []byte = []byte{
	0, 0, 0, 0, 0, 0, 0, 3, 
	0, 0, 0, 0, 0, 0, 0, 0, 
	0, 0, 0, 0, 0, 
}

var _simple_lexer_from_state_actions []byte = []byte{
	0, 0, 0, 0, 0, 0, 0, 5, 
	0, 0, 0, 0, 0, 0, 0, 0, 
	0, 0, 0, 0, 0, 
}

var _simple_lexer_eof_trans []byte = []byte{
	0, 0, 3, 0, 0, 0, 0, 0, 
	16, 18, 19, 19, 19, 22, 22, 19, 
	19, 19, 19, 19, 19, 
}

const simple_lexer_start int = 7
const simple_lexer_first_final int = 7
const simple_lexer_error int = 0

const simple_lexer_en_gosh int = 7


//line lexer_go.rl:49
var cs, p, pe, ts, te, act, eof, mark int
var program *Program

func reset() {
  cs, p, pe, ts, te, act, eof, mark = 0,0,0,0,0,0,0,0
}
func emit(token_name Token, data []byte, program *Program, ts int, te int) {
  value := string(data[ts:te])
  program.append(Statement{Name: token_name, Value: value})
}

func run_machine(input string) *Program {
  reset()
  program = NewProgram()
  data := []byte(input)
  pe = len(data)
  eof = len(data)
  //eof = -1
  
//line lexer_go.go:134
	{
	cs = simple_lexer_start
	ts = 0
	te = 0
	act = 0
	}

//line lexer_go.rl:68
  
//line lexer_go.go:144
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

//line lexer_go.go:167
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
	_trans = int(_simple_lexer_indicies[_trans])
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
 fmt.Printf("Current mark: %d, fpc: %d -> %s, fc: %s\n", mark, p, data[mark:p], data[p]) 
		case 3:
//line NONE:1
te = p+1

		case 4:
//line lexer_go.rl:27
te = p+1
{ 
      emit(ASSIGNMENT, data, program, ts, te) 
    }
		case 5:
//line lexer_go.rl:39
te = p+1

		case 6:
//line lexer_go.rl:19
te = p
p--
{
      emit(COMMAND, data, program, ts, te) 
    }
		case 7:
//line lexer_go.rl:23
te = p
p--
{ 
      emit(FLOAT, data, program, ts, te) 
    }
		case 8:
//line lexer_go.rl:31
te = p
p--
{ 
      emit(IDENTIFIER, data, program, ts, te) 
    }
		case 9:
//line lexer_go.rl:35
te = p
p--
{ 
      emit(INTEGER, data, program, ts, te) 
    }
		case 10:
//line lexer_go.rl:35
p = (te) - 1
{ 
      emit(INTEGER, data, program, ts, te) 
    }
//line lexer_go.go:288
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

//line lexer_go.go:302
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

//line lexer_go.rl:69
  
  //fmt.Printf("Finished. The state of the machine is: %s\n", cs)
  //fmt.Printf("p: %s pe: %s\n", p, pe)
  return program
}
