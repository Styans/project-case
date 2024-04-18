$('.cert_slider').slick({
    arrows: false,
    dots: true,
    appendDots: $('.cert_dots'),
    waitForAnimate: false,
    // responsive:
    //   [
    //     {
    //       breakpoint: 700,
    //       settings: {

    //       },
    //     },
    //   ]
  })
  $('.testimonials__prev').on('click', function (e) {
    e.preventDefault()
    $('.testimonials__slider').slick('slickPrev')
  })
  $('.testimonials__next').on('click', function (e) {
    e.preventDefault()
    $('.cert_slider').slick('slickNext')
  })