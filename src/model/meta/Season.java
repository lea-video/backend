package model.meta;

import model.core.Media;
import model.core.Playable;
import model.utility.Title;

import java.util.ArrayList;
import java.util.List;

public class Season implements Playable {
    Title title;
    Series series;
    List<Media> episode;

    public Season(Series series, Title title) {
        this.title = title;
        this.series = series;
        this.episode = new ArrayList<>();
    }

    public void addEpisode(Media m) {
        this.episode.add(m);
    }

    @Override
    public Title getTitle() {
        return this.title;
    }

    @Override
    public Playable getNext(int current) {
        if (current + 1 < this.episode.size()) {
            return episode.get(current + 1);
        }
        // Try to return next Season
        return this.series.getNext(this.series.seasons.indexOf(this));
    }
}
