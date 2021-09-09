import java.util.*;
import java.io.*;

public class MainSimulation extends Global {

	public static void main(String[] args) throws IOException {

		// utplockade signalen i huvudloopen nedan.
		// The signal list is started and actSignal is declaree. actSignal is the latest
		// signal that has been fetched from the
		// signal list in the main loop below.

		Signal actSignal;
		ConfigParser parser = new ConfigParser("config.txt");
		new SignalList();

		parser.parse();
		for (Tower tower : parser.getTowers()) {
			SignalList.SendSignal(START, tower, time + Global.exp(parser.getTs()));

		}

		// Here process instances are created (two queues and one generator) and their
		// parameters are given values.

		// To start the simulation the first signals are put in the signal list

		// This is the main loop

		while (time < 100000) {
			actSignal = SignalList.FetchSignal();
			time = actSignal.arrivalTime;
			actSignal.destination.TreatSignal(actSignal);
		}
		System.out.println("load: "+(double) parser.getGateway().arrived / time);
		System.out.println("tput: "+(double) parser.getGateway().succeded / time);

		// Slutligen skrivs resultatet av simuleringen ut nedan:
		// Finally the result of the simulation is printed below:

	}
}