package model.meta;

import model.core.Media;
import model.core.Playable;
import model.utility.Title;

import java.util.ArrayList;
import java.util.List;

public class ImageSet implements Playable {
    Title title;
    List<Media> images;

    public ImageSet(Title title) {
        this.title = title;
        this.images = new ArrayList<>();
    }

    public void addImage(Media m) {
        this.images.add(m);
    }

    @Override
    public Title getTitle() {
        return this.title;
    }

    @Override
    public Playable getNext(int current) {
        if (current + 1 >= this.images.size()) {
            return null;
        }
        return images.get(current + 1);
    }
}
