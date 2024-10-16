ASSUME DS:DATA,SS:STACK,CS:CODE
DATA SEGMENT
      str  DB 'hElLo WoRlD','$'
DATA ENDS

STACK SEGMENT
            DB 10 DUP (0)
STACK ENDS

CODE SEGMENT
      start:
            mov  ax,data
            mov  ds,ax

            mov  bx,0
            mov  cx,11
      s:    
            mov  al,str[bx]      ;去除该字母的asc码
            or   al,100000B
      ;转大写：&1011111B
      ;转小写：|100000B
            mov  str[bx],al      ;转换后的asc码放回原地
            inc  bx              ;下一个字母
            loop s

      ;mov dx,offset str
            lea  dx,str
            mov  ah,9            ;将dx的内容显示到屏幕
            int  21H

            mov  ah,4cH
            int  21H

CODE ENDS
end start