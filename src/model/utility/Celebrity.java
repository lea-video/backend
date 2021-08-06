package model.utility;

import java.util.ArrayList;
import java.util.List;

public class Celebrity {
    String name;
    List<Tag> tags;

    public Celebrity(String name) {
        this.name = name;
        this.tags = new ArrayList<>();
    }
}
