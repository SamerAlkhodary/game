import java.util.Random;

public class Global {
	public static final int START = 1, END = 2, MEASURE = 3, FAILED = 4,FAILED_REACH = 5;
	public static double time = 0;

	public static double exp(int mean) {
		Random rand = new Random();
		double t = Math.log(1-rand.nextDouble())/(-1.0/mean);
		return t;

	}
}
