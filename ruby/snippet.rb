# replace regex
def maskify(cc)
  cc.gsub(/.(?=.{4})/,'#')
end

def is_int_array(arr)
  arr.all? {|e| e % 1.0 == 0} rescue false
end
