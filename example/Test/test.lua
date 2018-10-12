local m = require("module")
function max(num1, num2)
   if (num1 > num2) then
      result = num1;
   else
      result = num2;
   end
   print(result)
   return result,num1,num2; 
end
print(test(100,2))
print(m.mFunc(1,2))
print(m.Name)