#!/usr/bin/env julia

function main()
  elves = [0]
  #open("simple.txt") do f
  open("calories.txt") do f
    while ! eof(f)
      if (line = readline(f)) != ""
        elves[end] += parse(Int, line)
      else
        push!(elves, 0)
      end
    end
  end

  sort!(elves, rev=true)
  println("biggest: ", elves[1])
  println("top 3: ", sum(elves[1:3]))
end

main()

