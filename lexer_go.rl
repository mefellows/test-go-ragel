%%{

  machine simple_lexer;

  action print { fmt.Printf("Current mark: %d, fpc: %d -> %s, fc: %s\n", mark, p, data[mark:p], fc) }
  action history { fmt.Printf("history match\n") }
  action sudo { fmt.Printf("sudo match\n") }
  action copy { fmt.Printf("copy match\n") }
  
  integer     = ('+'|'-')?[0-9]+;
  float       = ('+'|'-')?[0-9]+'.'[0-9]+;
  assignment  = '=';
  identifier  = [a-zA-Z][a-zA-Z_]+; 
#  command     = ('sudo'|'history'|'copy');
  command = ('sudo'|'history'|'copy') >print;
  
  gosh := |*

    command any* => {
      emit(COMMAND, data, program, ts, te) 
    };

    float => { 
      emit(FLOAT, data, program, ts, te) 
    };
    
    assignment => { 
      emit(ASSIGNMENT, data, program, ts, te) 
    };
    
    identifier => { 
      emit(IDENTIFIER, data, program, ts, te) 
    };
    
    integer => { 
      emit(INTEGER, data, program, ts, te) 
    };
    
    space;
    
  *|;

}%%

package main
import "fmt"

%% write data;
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
  %% write init;
  %% write exec;
  
  //fmt.Printf("Finished. The state of the machine is: %s\n", cs)
  //fmt.Printf("p: %s pe: %s\n", p, pe)
  return program
}
