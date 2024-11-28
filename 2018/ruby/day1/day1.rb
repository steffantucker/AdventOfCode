# Calculates the sums of the alternating frequencies
# Provides answer to day 1's first star
def calcFrequency(freqs)
	sum = freqs.reduce(0, :+)
	puts sum
end

# Calculates repeated sum frequency after
# n loops through frequency list
def calcCalibrationFrequency(freqs)
	sum = 0
	sums = []
	cycles = 0
	i = 0
	puts "Calculating calibration freq"
	puts "this will take a bit"
	while (!(sums.include? sum))
		sums.push(sum)
		sum += freqs[i]
		if ((i + 1) >= freqs.length)
			i = 0
			cycles += 1
		else
			i += 1
		end
	end
	puts sum
end

frequencyList = []
File.foreach('./day1.inputs') { |x| frequencyList.push(x.to_i) }
calcFrequency(frequencyList)
calcCalibrationFrequency(frequencyList)
