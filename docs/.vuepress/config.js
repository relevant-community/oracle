module.exports = {
  theme: 'cosmos',
  title: 'Relevant Oracle',
  locales: {
    '/': {
      lang: 'en-US',
    },
    kr: {
      lang: 'kr',
    },
    cn: {
      lang: 'cn',
    },
    ru: {
      lang: 'ru',
    },
  },
  base: '/oracle/',
  head: [
    [
      'link',
      {
        rel: 'apple-touch-icon',
        sizes: '180x180',
        href: '/apple-touch-icon.png',
      },
    ],
    [
      'link',
      {
        rel: 'icon',
        type: 'image/png',
        sizes: '32x32',
        href: '/favicon-32x32.png',
      },
    ],
    [
      'link',
      {
        rel: 'icon',
        type: 'image/png',
        sizes: '16x16',
        href: '/favicon-16x16.png',
      },
    ],
    ['link', { rel: 'manifest', href: '/site.webmanifest' }],
    ['meta', { name: 'msapplication-TileColor', content: '#2e3148' }],
    ['meta', { name: 'theme-color', content: '#ffffff' }],
    ['link', { rel: 'icon', type: 'image/svg+xml', href: '/favicon-svg.svg' }],
    [
      'link',
      {
        rel: 'apple-touch-icon-precomposed',
        href: '/apple-touch-icon-precomposed.png',
      },
    ],
  ],
  themeConfig: {
    repo: 'relevant-community/oracle',
    docsRepo: 'relevant-community/oracle',
    docsDir: 'docs',
    label: 'sdk',
    custom: true,
    topbar: {
      banner: false,
    },
    sidebar: {
      auto: false,
      nav: [
        {
          title: 'Documentation',
          children: [
            {
              title: 'Atom/USD Tutorial',
              directory: true,
              path: '/tutorial',
            },
            {
              title: 'Oracle Module Docs',
              directory: true,
              path: '/modules/oracle',
            },
          ],
        },
      ],
    },

    footer: {
      logo: '/logo.svg',
      textLink: {
        text: 'cosmos.network',
        url: 'https://cosmos.network',
      },
      services: [
        {
          service: 'twitter',
          url: 'https://twitter.com/relevantfeed',
        },
        {
          service: 'github',
          url: 'https://github.com/relevant-community/oracle',
        },
      ],
      links: [
        {
          title: 'Contributing',
          children: [
            {
              title: 'Source code on GitHub',
              url: 'https://github.com/relevant-community/oracle',
            },
          ],
        },
        {
          title: 'Related Docs',
          children: [
            {
              title: 'Cosmos SDK',
              url: 'https://cosmos.network/docs',
            },
          ],
        },
      ],
    },
  },
}
