import java.io.*;
import org.json.JSONArray;
import org.json.JSONTokener;

/*
 Simple helper to read/write JSON arrays for Opset.
 This class uses org.json (the JSON-Java library). If you don't have it on the
 classpath, the Java helper will still compile if you remove references or add
 the dependency later.
*/

public class OpsetHelper {
    public static JSONArray readArray(File f) throws IOException {
        try (FileReader fr = new FileReader(f)) {
            JSONTokener tok = new JSONTokener(fr);
            Object val = tok.nextValue();
            if (val instanceof JSONArray) {
                return (JSONArray) val;
            } else {
                return new JSONArray();
            }
        }
    }

    public static void writeFile(File f, JSONArray arr) throws IOException {
        try (FileWriter fw = new FileWriter(f)) {
            fw.write(arr.toString(2));
        }
    }
}
