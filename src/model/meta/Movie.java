package model.meta;

import model.core.Media;
import model.core.Playable;
import model.utility.Title;

import java.util.ArrayList;
import java.util.List;

public class Movie implements Playable {
    Title title;
    Movie sequel;
    List<Media> cuts;

    public Movie(Title title) {
        this.title = title;
        this.cuts = new ArrayList<>();
    }

    public void addCut(Media m) {
        this.cuts.add(m);
    }

    public Media getCuts(int cut) {
        if (cut < 0 || cut >= this.cuts.size()) {
            return null;
        }
        return this.cuts.get(cut);
    }

    @Override
    public Title getTitle() {
        return this.title;
    }

    @Override
    public Playable getNext(int current) {
        return this.sequel;
    }

    public void setSequel(Movie m) {
        this.sequel = m;
    }
}
