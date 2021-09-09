public class Gateway extends Proc {
    private Coords coordinations;
    private int tp;
    public int arrived, failed, succeded, finised,tried_reach;

    public Gateway(int x, int y, int tp) {


        this.coordinations = new Coords(x, y);

        this.tp = tp;
    }

    private boolean isBusy() {
        return arrived != finised;
    }

    public void TreatSignal(Signal x) {
        switch (x.signalType) {
        case START:
        arrived++;

                SignalList.SendSignal(END, this, time + tp);
            break;

        case END:
            finised++;
            if (isBusy()) {
                SignalList.SendSignal(FAILED, this, time);
            } else {
                succeded++;
            }
            break;

        case FAILED_REACH:
            tried_reach++;
            
            break;

        case FAILED:
            failed++;
            break;

        }
    }

    public Coords getCoords() {
        return this.coordinations;
    }

}