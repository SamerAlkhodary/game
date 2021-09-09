public class Coords {
    private int x, y;

    public Coords(int x, int y) {
        this.x = x;
        this.y = y;
    }

    public double distanceFrom(Coords other) {

        return Math.hypot(this.x - other.x, this.y - other.y);
    }

}