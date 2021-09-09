run:
	@echo  "Generating and running the simulation"
	python3 script.py 
	mkdir files 
	cp data.json config.txt files 
	javac *.java 
	java MainSimulation
clean: 
	@echo "removing .class files"
	rm *.class