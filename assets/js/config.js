(function () {
  const STORAGE_KEY = 'ccaSiteConfig';

  const defaultConfig = {
    branding: {
      orgName: 'Cadott Community Association',
      headerTitle: 'Neighbors building Cadott since 1970',
      headerSubtitle: 'Nabor Days · Music in the Park · Volunteering',
      footerNote: 'Built with local pride.',
      headerButtons: [
        {
          label: 'Get involved',
          href: 'contact.html#meetings',
          style: 'primary'
        }
      ]
    },
    meeting: {
      banner: 'Next meeting: TBD',
      heroTitle: 'Next meeting TBD.',
      heroBody: 'We will post the date, time, and location as soon as the board confirms it. Show up, listen, and plug into Nabor Days, Music in the Park, or neighborhood volunteering.'
    },
    home: {
      hero: {
        badge: 'Cadott, Wisconsin',
        title: 'Mission: To enhance our community by strengthening business and citizen involvement.',
        body: 'We keep things simple—short meetings, printed notes, and hometown events that anyone can join or enjoy.',
        figure: {
          src: 'https://www.cadottcommunity.com/uploads/6/4/1/1/64113323/19554523-681567728699006-1629257994525173032-n_orig.jpg',
          alt: 'Nabor Days parade rolling through downtown Cadott',
          caption: 'Nabor Days parade floats and volunteers keep Main Street buzzing.'
        }
      },
      involvement: {
        kicker: 'Get involved',
        title: 'How to plug in.',
        subtitle: 'Three fast steps so you can raise your hand and help neighbors this season.'
      },
      nabor: {
        badge: 'Nabor Days',
        title: 'The last full weekend in July.',
        body: 'Live music, sunrise runs, parades, and horsepower. Add your name to the schedule now so every lane, tent, and float feels ready when Cadott shows up.',
        highlightHeading: 'Weekend highlights',
        highlightSubhead: 'Drag, drop, reorder in admin mode.',
        note: 'Circle the dates now—we post the full schedule each spring at meetings and online.'
      },
      steps: [
        {
          label: 'Step 1',
          body: 'Next meeting TBD. Watch this page or the email list for the posted date and agenda.'
        },
        {
          label: 'Step 2',
          body: 'Add your name to the volunteer sheet. We pair you with bingo nights, music events, or Nabor Days shifts.'
        },
        {
          label: 'Step 3',
          body: 'Questions later? Email the board or message us on Facebook and we will walk you through it.'
        }
      ],
      naborHighlights: [
        'Live music, crafts, and vendor rows fill Riverview Park.',
        'Sunday parade down Main Street plus royalty coronation.',
        'Truck & Tractor Pull, Midwest Motorsports Mud Bog, and Lions Club chicken dinner with FFA cheese curds on Hwy 27.',
        'Fun Run/Walk and family activities each morning.'
      ],
      heroButtons: [
        {
          label: 'Next meeting details',
          href: 'contact.html#meetings',
          style: 'primary'
        },
        {
          label: 'Email or message us',
          href: 'mailto:cadott-community@googlegroups.com',
          style: 'secondary'
        }
      ],
      getInvolvedButtons: [
        {
          label: 'Meetings & members page',
          href: 'contact.html#meetings',
          style: 'primary'
        },
        {
          label: 'Simple FAQ',
          href: 'faq.html',
          style: 'secondary'
        }
      ]
    },
    contact: {
      email: 'cadott-community@googlegroups.com',
      facebook: 'https://facebook.com/CadottNaborDays',
      heroButtons: [
        {
          label: 'Email the board',
          href: 'mailto:cadott-community@googlegroups.com',
          style: 'primary'
        },
        {
          label: 'Sign up for the newsletter',
          href: 'https://landing.mailerlite.com/webforms/landing/s6z1s0',
          style: 'secondary',
          target: '_blank',
          rel: 'noreferrer noopener'
        }
      ],
      board: [
        'Anna Goodman — President',
        'Lucy Meinen — Vice President',
        'Jess Buckli — Secretary',
        'Nikki Ruhe — Treasurer',
        'Ashley Anderson — Media & parade contact',
        'Cindy Mayer — Director',
        'Nikki Vaughn — Director'
      ],
      cadence: [
        'Updates · fundraising, sponsors, quick event recaps.',
        'Planning · Nabor Days, Music in the Park, community events.',
        'Open floor · needs, ideas, and causes that need a boost.'
      ]
    },
    event: {
      kicker: 'Mid-July 2026',
      title: 'Nabor Days — Cadott\'s summer celebration!',
      description: 'Nabor Days is the biggest event of the summer, held right here in Cadott at Riverview Park. Bring the whole family for a weekend of live music, delicious food stands, the grand parade down Main Street, a softball tournament, and much more!',
      location: 'Riverview Park',
      heroButtons: [
        {
          label: 'RSVP on Facebook',
          href: 'https://www.facebook.com/cadottcommunityassociation',
          style: 'primary'
        },
        {
          label: 'Volunteer to help',
          href: 'mailto:cadott-community@googlegroups.com',
          style: 'secondary'
        }
      ],
      schedule: [
        'Friday evening · Vendor tents open, live music kicks off the weekend',
        'Saturday morning · Annual Nabor Days Parade down Main Street',
        'Saturday afternoon · Softball tournament, kid\'s games, and food stands in Riverview Park',
        'Saturday night · Street dance and headline band performance',
        'Sunday morning · Community breakfast and wrap-up'
      ],
      timeline: [
        {
          title: 'Opening Ceremonies',
          time: 'Friday Evening',
          description: 'The weekend officially begins with food vendors opening up and the first round of live entertainment.'
        },
        {
          title: 'Nabor Days Parade',
          time: 'Saturday Morning',
          description: 'Grab a spot on Main Street to watch the floats, marching bands, and local businesses in our annual parade.'
        },
        {
          title: 'Park Activities',
          time: 'Saturday Afternoon',
          description: 'Head down to Riverview Park for the softball tournament, carnival games, and delicious local food stands.'
        },
        {
          title: 'Street Dance',
          time: 'Saturday Night',
          description: 'Put on your dancing shoes for the headline performance and street dance outside. A Cadott tradition!'
        },
        {
          title: 'Wrap-Up',
          time: 'Sunday Morning',
          description: 'Join us for a community pancake breakfast and help with the Sunday town cleanup before heading out.'
        }
      ],
      highlights: [
        'Grand Parade down Main Street featuring local businesses and organizations.',
        'Live music and street dances running all weekend long.',
        'Charity softball tournament in Riverview Park.',
        'Fantastic local food stands, beer gardens, and family activities.',
        'Proceeds from CCA run stands go straight to community improvements.'
      ],
      preview: [
        {
          title: 'Where to park',
          details: 'Parking is available at the high school lot with a free shuttle, or on side streets near Riverview Park. Please be respectful of neighbors\' driveways.'
        },
        {
          title: 'What to Expect',
          details: 'A full weekend of family-friendly fun! Bring lawn chairs for the parade, sunscreen for the park, and get ready to catch up with old friends.'
        },
        {
          title: 'Enter a Float',
          details: 'Want your business or group in the parade? Email cadott-community@googlegroups.com to register your float.'
        }
      ]
    },
    eventListings: [
      {
        title: 'Nabor Days Festival',
        timing: 'Mid-July 2026 · Riverview Park',
        summary: 'Cadott\'s biggest summer festival! Enjoy the grand parade, softball tournaments, live music, and food stands all weekend.',
        details: 'Join us for this time-honored tradition. All proceeds from CCA run stands go directly back into the community.'
      },
      {
        title: 'Music in the Park',
        timing: 'Wednesdays in June & July · 6:30 PM · Riverview Park Bandshell',
        summary: 'Free midweek concerts featuring local bands, Lions Club concessions, and canned good drives.',
        details: 'Book bands, schedule food trucks, manage sound, and collect donations for the Cadott Food Pantry.'
      },
      {
        title: 'Nabor Days Weekend',
        timing: 'Last full weekend in July · Riverview Park, Hwy 27, Main Street',
        summary: 'Truck & Tractor Pull, Mud Bog, Fun Run, parade, vendors, and fireworks anchor the summer.',
        details: 'Needs parade marshals, beer tent crews, bingo callers, kids’ zone leads, and Sunday cleanup captains.'
      }
    ],
    gallery: [
      {
        src: 'https://www.cadottcommunity.com/uploads/6/4/1/1/64113323/19554523-681567728699006-1629257994525173032-n_orig.jpg',
        alt: 'Nabor Days parade rolling down Main Street in Cadott',
        caption: 'Main Street parade — the highlight of Nabor Days weekend.'
      },
      {
        src: 'https://www.cadottcommunity.com/uploads/6/4/1/1/64113323/20190727-095834_orig.jpg',
        alt: 'Kids racing in the free 800M dash at Nabor Days',
        caption: 'Fun Run/Walk sets the tone for the whole weekend.'
      },
      {
        src: 'https://www.cadottcommunity.com/uploads/6/4/1/1/64113323/church-band_orig.jpg',
        alt: 'Church band marching in the Nabor Days parade',
        caption: 'Local bands bring the music down Main Street.'
      },
      {
        src: 'https://www.cadottcommunity.com/uploads/6/4/1/1/64113323/fire-trucks_orig.jpg',
        alt: 'Fire trucks rolling through the Nabor Days parade',
        caption: 'Cadott fire trucks lead the parade every year.'
      },
      {
        src: 'https://www.cadottcommunity.com/uploads/6/4/1/1/64113323/2022-queens_orig.jpg',
        alt: '2022 Nabor Days Queens waving from parade float',
        caption: 'Nabor Days royalty rides through downtown Cadott.'
      },
      {
        src: 'https://www.cadottcommunity.com/uploads/6/4/1/1/64113323/67136947-1086973161491792-6632270262052061184-n_orig.jpg',
        alt: 'Queen coronation and Grand Marshal presentation on stage',
        caption: 'Queen coronation night crowns Cadott pride.'
      },
      {
        src: 'https://www.cadottcommunity.com/uploads/6/4/1/1/64113323/20294234-681567775365668-2539194200402116317-n_orig.jpg',
        alt: 'Kids gathered under trees at Cadott Riverview Park during Nabor Days',
        caption: 'Youth games and neighbor reunions beneath the park canopy.'
      },
      {
        src: 'https://www.cadottcommunity.com/uploads/6/4/1/1/64113323/published/60652439-941365229541547-9109737127872036864-n.jpg?1608661349',
        alt: 'Tractor pulling sled at dusk during Nabor Days',
        caption: 'Friday night horsepower lights up the Truck & Tractor Pull.'
      },
      {
        src: 'https://www.cadottcommunity.com/uploads/6/4/1/1/64113323/published/img-20170729-155254084_1.jpg?1684941294',
        alt: 'Trucks roaring through the Cadott mud bog',
        caption: 'Midwest Motorsports Association races through 150 feet of Cadott mud.'
      },
      {
        src: 'https://www.cadottcommunity.com/uploads/6/4/1/1/64113323/park_orig.jpg',
        alt: 'Families at Riverview Park along the Yellow River',
        caption: 'Riverview Park — the heart of Cadott community life.'
      },
      {
        src: 'https://www.cadottcommunity.com/uploads/6/4/1/1/64113323/published/music-1.jpg?1608663177',
        alt: 'Band performing in Riverview Park during Music in the Park',
        caption: 'Music in the Park transforms Riverview Park into a midweek concert hall.'
      },
      {
        src: 'https://www.cadottcommunity.com/uploads/6/4/1/1/64113323/editor/muysic-2.jpg?1608663059',
        alt: 'Crowd enjoying live music under the lights at Riverview Park',
        caption: 'Neighbors pack lawn chairs for free shows under the lights.'
      }
    ],
    faq: {
      heroButtons: [
        {
          label: 'Submit a question',
          href: 'mailto:cadott-community@googlegroups.com',
          style: 'primary'
        },
        {
          label: 'Join the next meeting',
          href: 'contact.html#meetings',
          style: 'secondary'
        }
      ],
      community: [
        {
          question: 'How do I become a member?',
          answer: 'Email cadott-community@googlegroups.com or walk into a meeting. We will cover dues and match you with a job.'
        },
        {
          question: 'When and where do you meet?',
          answer: 'Next meeting TBD. Rotating Cadott locations. We will update the calendar and email list once the board finalizes the details.'
        },
        {
          question: 'Who is on the board?',
          answer: 'Anna Goodman (President), Lucy Meinen (Vice President), Jess Buckli (Secretary), Nikki Ruhe (Treasurer), Ashley Anderson (Media), Cindy Mayer & Nikki Vaughn (Directors).'
        },
        {
          question: 'Can businesses join?',
          answer: 'Yes. Sponsor, host, or simply attend—the more visibility, the better.'
        }
      ],
      festival: [
        {
          question: 'When is Nabor Days?',
          answer: 'Last full weekend in July. Friday night through Sunday afternoon.'
        },
        {
          question: 'Where is everything located?',
          answer: 'Riverview Park hosts vendors, pulls, and music; Mud Bog lines Hwy 27; Sunday parade rolls Main Street.'
        },
        {
          question: 'How do I register for the parade?',
          answer: 'Email cadottchirodc@gmail.com or call 715-579-3011 with entry details.'
        },
        {
          question: 'What else happens?',
          answer: 'Expect bean bag + dart tournaments, bingo, duck races, beer-tent bands, fireworks, pancake breakfast, family zone, kiddie tractor pull.'
        }
      ]
    },
    galleryPage: {
      heroButtons: [
        {
          label: 'See this year’s lineup',
          href: 'events.html',
          style: 'primary'
        },
        {
          label: 'Share your photos',
          href: 'mailto:cadott-community@googlegroups.com',
          style: 'secondary'
        }
      ],
      naborDays: [
        {
          src: 'https://www.cadottcommunity.com/uploads/6/4/1/1/64113323/20294234-681567775365668-2539194200402116317-n_orig.jpg',
          alt: 'Kids gathered under trees at Cadott Riverview Park during Nabor Days',
          caption: 'Youth games and neighbor reunions beneath the Riverview Park canopy.'
        },
        {
          src: 'https://www.cadottcommunity.com/uploads/6/4/1/1/64113323/19554523-681567728699006-1629257994525173032-n_orig.jpg',
          alt: 'Local vendors displaying goods at the Nabor Days craft and vendor show',
          caption: 'Vendors stretching through the park offer handmade goods, décor, and treats.'
        },
        {
          src: 'https://www.cadottcommunity.com/uploads/6/4/1/1/64113323/editor/run-walk.jpg?1527728364',
          alt: 'Runners at the Cadott Nabor Days 5K starting line',
          caption: 'Saturday’s Fun Run/Walk brings families to the Cadott High School track at sunrise.'
        },
        {
          src: 'https://www.cadottcommunity.com/uploads/6/4/1/1/64113323/20190727-095834_orig.jpg',
          alt: 'Kids racing in the free 800M dash',
          caption: 'Little legs, big cheers—kids sprint the free 800M dash after the 5K.'
        },
        {
          src: 'https://www.cadottcommunity.com/uploads/6/4/1/1/64113323/published/60652439-941365229541547-9109737127872036864-n.jpg?1608661349',
          alt: 'Tractor pulling sled at dusk',
          caption: 'Friday night horsepower lights up the Truck & Tractor Pull grandstand.'
        },
        {
          src: 'https://www.cadottcommunity.com/uploads/6/4/1/1/64113323/published/img-20170729-155254084_1.jpg?1684941294',
          alt: 'Trucks roaring through the Cadott mud bog',
          caption: 'Midwest Motorsports Association races through 150 feet of Cadott mud.'
        },
        {
          src: 'https://www.cadottcommunity.com/uploads/6/4/1/1/64113323/67136947-1086973161491792-6632270262052061184-n_orig.jpg',
          alt: 'Nabor Days queen candidates on stage',
          caption: 'Queen coronation night crowns the neighbor who best represents Cadott pride.'
        }
      ],
      winter: [
        {
          src: 'https://www.cadottcommunity.com/uploads/6/4/1/1/64113323/editor/santa.jpg?250',
          alt: 'Santa greeting a child during a Cadott holiday event',
          caption: 'Santa meet-and-greets kick off the Christmas Village festivities.'
        },
        {
          src: 'https://www.cadottcommunity.com/uploads/6/4/1/1/64113323/published/horses.jpg?1608662445',
          alt: 'Horse drawn sleigh crossing Riverview Park',
          caption: 'Horse-drawn sleigh rides glide through the Christmas Village each December.'
        },
        {
          src: 'https://www.cadottcommunity.com/uploads/6/4/1/1/64113323/published/29250085-776904239165354-5852604148729511936-n.jpeg?1522162773',
          alt: 'Neighbors gathered for bingo',
          caption: 'Cozy bingo and meat raffles warm up the winter calendar.'
        },
        {
          src: 'https://www.cadottcommunity.com/uploads/6/4/1/1/64113323/published/crafts.jpg?1608662399',
          alt: 'Holiday craft booth',
          caption: 'Handmade craft markets showcase local makers during the colder months.'
        }
      ],
      music: [
        {
          src: 'https://www.cadottcommunity.com/uploads/6/4/1/1/64113323/published/music-1.jpg?1608663177',
          alt: 'Band performing in Riverview Park',
          caption: 'Music in the Park transforms Riverview Park into a midweek concert hall.'
        },
        {
          src: 'https://www.cadottcommunity.com/uploads/6/4/1/1/64113323/editor/muysic-2.jpg?1608663059',
          alt: 'Crowd enjoying live music under string lights',
          caption: 'Neighbors pack lawn chairs for free shows under the lights.'
        },
        {
          src: 'https://www.cadottcommunity.com/uploads/6/4/1/1/64113323/track_orig.jpg',
          alt: 'Neighbors exercising on the Cadott track',
          caption: 'Community fitness nights keep neighbors active on Monday evenings.'
        },
        {
          src: 'https://www.cadottcommunity.com/uploads/6/4/1/1/64113323/yoga_orig.jpg',
          alt: 'Group yoga session in Riverview Park',
          caption: 'Sunset yoga sessions take over the west lawn every Wednesday.'
        },
        {
          src: 'https://www.cadottcommunity.com/uploads/6/4/1/1/64113323/park_orig.jpg',
          alt: 'Families relaxing along the river during Music in the Park',
          caption: 'Families settle along the riverbank for concessions and community.'
        },
        {
          src: 'https://www.cadottcommunity.com/uploads/6/4/1/1/64113323/published/20476236-681948565327589-17922207873877258-n-1_1.jpeg?1522162797',
          alt: 'Neighbors sharing a laugh during game night',
          caption: 'Game nights and volunteer shifts double as easy networking hubs.'
        }
      ]
    }
  };

  const deepClone = (value) => JSON.parse(JSON.stringify(value));

  const mergeDeep = (target, source) => {
    if (typeof source !== 'object' || source === null) {
      return target;
    }

    Object.keys(source).forEach((key) => {
      const sourceValue = source[key];
      const targetValue = target[key];

      if (Array.isArray(sourceValue)) {
        target[key] = sourceValue.slice();
      } else if (typeof sourceValue === 'object' && sourceValue !== null) {
        target[key] = mergeDeep(targetValue ? { ...targetValue } : {}, sourceValue);
      } else {
        target[key] = sourceValue;
      }
    });

    return target;
  };

  const loadConfig = () => {
    try {
      const stored = localStorage.getItem(STORAGE_KEY);
      if (!stored) {
        return deepClone(defaultConfig);
      }
      const parsed = JSON.parse(stored);
      const merged = mergeDeep(deepClone(defaultConfig), parsed);
      return merged;
    } catch (error) {
      console.warn('CCAConfig: failed to load config, using defaults.', error);
      return deepClone(defaultConfig);
    }
  };

  const saveConfig = (config) => {
    try {
      localStorage.setItem(STORAGE_KEY, JSON.stringify(config));
    } catch (error) {
      console.error('CCAConfig: unable to save config.', error);
    }
  };

  const resetConfig = () => {
    try {
      localStorage.removeItem(STORAGE_KEY);
    } catch (error) {
      console.error('CCAConfig: unable to reset config.', error);
    }
  };

  window.CCAConfig = {
    defaultConfig: deepClone(defaultConfig),
    loadConfig,
    saveConfig,
    resetConfig,
    STORAGE_KEY
  };
})();
