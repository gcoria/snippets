# @param {Integer[]} arr
# @return {Void} Do not return anything, modify arr in-place instead.
def duplicate_zeros(arr)
  i = 0
  original_length = arr.length

  while i < original_length
    if arr[i] == 0
      arr.insert(i + 1, 0)
      i += 1
    end
    i += 1
  end

  # Truncate the array to the original length
  arr.slice!(original_length..-1)
end