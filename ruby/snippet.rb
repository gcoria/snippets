# replace regex
def maskify(cc)
  cc.gsub(/.(?=.{4})/,'#')
end

def is_int_array(arr)
  arr.all? {|e| e % 1.0 == 0} rescue false
end

def list names
  names = names.map { |name| name[:name] }
  last_name = names.pop
  return last_name.to_s if names.empty?
  "#{names.join(', ')} & #{last_name}"
end

def compute_depth n
  h={}
  1.step.find{|x|(?0..?9).all?(&h)||(
  d=n*x
  d.to_s.chars.each{|x|h[x]=1}
  !1)}-1
end

def up_array(arr)
  return nil if arr.empty? || arr.any? { |x| x < 0 || x > 9 }
  arr.join.next.chars.map(&:to_i)
end

def title_case(title, minor_words = '')
  title.capitalize.split().map{|a| minor_words.downcase.split().include?(a) ? a : a.capitalize}.join(' ')
end

def ip_to_int32(ip)
  x = IPAddr.new(ip).to_i
end
