import java.io.File;
import java.util.ArrayList;
import java.util.List;
import java.util.Scanner;

public class ConfigParser {
    private String path;
    private int ts, tp, raduis, n, width, height;
    private List<Tower> towers = new ArrayList<>(n);
    private Gateway gateway;

    public ConfigParser(String path) {
        this.path = path;
    }

    public void parse() {
        File myObj = new File(path);
        try {
            Scanner myReader = new Scanner(myObj);
            String dim = myReader.nextLine();
            String[] dims = dim.split(",");
            this.width = Integer.parseInt(dims[0]);
            this.height = Integer.parseInt(dims[1]);
            this.tp = Integer.parseInt(myReader.nextLine());
            this.ts = Integer.parseInt(myReader.nextLine());
            this.raduis = Integer.parseInt(myReader.nextLine());
            dim = myReader.nextLine();
            dims = dim.split(",");
            this.gateway = new Gateway(Integer.parseInt(dims[0]), Integer.parseInt(dims[1]), this.tp);
            while (myReader.hasNextLine()) {
                String coord = myReader.nextLine();
                String[] coords = coord.split(",");
                towers.add(new Tower(this.gateway, Integer.parseInt(coords[0]), Integer.parseInt(coords[1]), this.ts,
                        this.tp, this.raduis));

            }
            myReader.close();
        } catch (Exception e) {
            System.out.println("An error occurred.");
            e.printStackTrace();
        }
    }

    public int getTs() {
        return this.ts;
    }

    public int getTp() {
        return this.tp;
    }

    public int getR() {
        return this.raduis;
    }

    public int getN() {
        return this.towers.size();
    }

    public int getWidth() {
        return this.width;
    }

    public int getHeight() {
        return this.height;
    }

    public List<Tower> getTowers() {
        return this.towers;
    }

    public Gateway getGateway() {
        return this.gateway;
    }

}
