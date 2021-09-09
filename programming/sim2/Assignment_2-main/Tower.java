public class Tower extends Proc {
    private Coords coordinations;
    int ts, tp, raduis;
    private Gateway gateway;

    public Tower(Gateway gateway, int x, int y, int ts, int tp, int raduis) {
        this.coordinations = new Coords(x, y);
        this.gateway = gateway;
        this.raduis = raduis;
        this.tp = tp;
        this.ts = ts;
    }

    public void TreatSignal(Signal x) {
        switch(x.signalType){
            case START: 
            if (canReachGateway()) {
                SignalList.SendSignal(START, gateway, time);
               
            }else{
                SignalList.SendSignal(FAILED_REACH, gateway, time);
            }
            SignalList.SendSignal(END, this, time + tp);
            break;
            case END:
                SignalList.SendSignal(START, this, time + Global.exp(ts));

            break;
        }
    }

    public Coords getCoords() {
        return this.coordinations;
    }

    public boolean canReachGateway() {
        return this.coordinations.distanceFrom(this.gateway.getCoords()) <= this.raduis;
    }

}
