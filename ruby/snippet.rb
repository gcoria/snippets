# replace regex
def maskify(cc)
  cc.gsub(/.(?=.{4})/,'#')
end