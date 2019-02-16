# @param {String} s
# @return {Boolean}
def is_valid(str)
  stack = []
  brackets = [
    { open: "(", close: ")" },
    { open: "[", close: "]" },
    { open: "{", close: "}" },
  ]

  str.each_char do |chr|
    brackets.each do |bracket|
      stack.push(bracket[:open]) if chr == bracket[:open]

      if chr == bracket[:close]
        return false unless stack.pop == bracket[:open]
      end
    end
  end

  return false unless stack.empty?

  true
end
