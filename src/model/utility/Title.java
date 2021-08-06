package model.utility;

import java.util.HashMap;
import java.util.Map;

public class Title {
    Map<String, String> translations;
    String def;

    public Title(String defaultTitle) {
        this.def = defaultTitle;
        this.translations = new HashMap<>();
    }
}
