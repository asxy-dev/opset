import java.io.File;

public class ConfigCheck {
    public static void main(String[] args) {
        System.out.println("ConfigCheck: lightweight validator placeholder");
        if (args.length == 0) {
            System.out.println("Usage: java ConfigCheck <file1> [file2 ...]");
            return;
        }
        for (String a : args) {
            File f = new File(a);
            System.out.println((f.exists() ? "OK: " : "MISSING: ") + a);
        }
    }
}
