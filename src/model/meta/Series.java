package model.meta;

import model.core.Playable;
import model.utility.Title;

import java.util.ArrayList;
import java.util.List;

public class Series implements Playable {
    Title title;
    List<Season> seasons;

    public Series(Title title) {
        this.title = title;
        this.seasons = new ArrayList<>();
    }

    public void addSeason(Season s) {
        this.seasons.add(s);
    }

    @Override
    public Title getTitle() {
        return this.title;
    }

    @Override
    public Playable getNext(int current) {
        if (current + 1 >= this.seasons.size()) {
            return null;
        }
        return seasons.get(current + 1);
    }
}
