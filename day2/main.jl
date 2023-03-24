#!/usr/bin/env julia

MAP= Dict("A" => 1, "B" => 2, "C" => 3, "X" => 1, "Y" => 2, "Z" => 3)
SCORES = [[4 8 3],
          [1 5 9],
          [7 2 6]]

function get_score(x, y )
  x, y = MAP[x], MAP[y]
  SCORES[x][y]
end

function main()
  score_rps = 0
  #open("simple.txt") do f
  open("guide.txt") do f
    while ! eof(f)
      x, y = split(readline(f), " ")
      score_rps += get_score(x, y)
    end
  end

  println("score (part 1): ", score_rps)
end

main()

