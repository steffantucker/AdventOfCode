def calculateChecksum(ids)
	twoLetterCount = 0
	threeLetterCount = 0
	letterCounts = Hash.new
	# counts the double and triple lettered words
	ids.each do |id|
		letterCounts.clear
		id.each_char do |letter|
			if(letterCounts.has_key? letter)
				letterCounts[letter] += 1
			else
				letterCounts[letter] = 1
			end
		end
		if (letterCounts.has_value? 2)
			twoLetterCount += 1
		end
		if (letterCounts.has_value? 3)
			threeLetterCount += 1
		end
	end
	puts "#{twoLetterCount} IDs with double letters"
	puts "#{threeLetterCount} IDs with triple letters"
	puts "Checksum: #{twoLetterCount * threeLetterCount}"
end

def findSimilarIds(ids)
	counts = Hash.new
	# get the sum of each words code count
	# the assumption being words that are only a few letters off
	# will have similar sums
	ids.each do |id|
		sum = 0
		id.each_codepoint.inject(:+)
		counts[id] = sum
	end

	closeOnes = []
	counts.each do |id, count|
		counts.each do |idj, countj|
			# similar sums are probably similar strings
			if id != idj && (count + 27) > countj && (count - 27) < countj
				wrongCount = 0
				# compare each letter directly
				id.length.times do |i|
					if id[i] != idj[i]
						wrongCount += 1
					end
					break if wrongCount > 1
				end
				# found some that are only a letter apart
				# now extract only the similar letters and print
				if wrongCount <= 1
					answer = ""
					id.length.times do |i|
						if id[i] == idj[i]
							answer += id[i]
						end
					end
					closeOnes.push "similar letters: #{answer}"
				end
			end
		end
		counts.delete(id)
	end
	puts closeOnes
end

# collect ids from inputs file
boxIds = []
File.foreach("./day2.inputs") { |x| boxIds.push(x.strip) }
puts "#{boxIds.length} box IDs"

calculateChecksum(boxIds)
findSimilarIds(boxIds)
